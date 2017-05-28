package app

import (
	"encoding/json"
	"fmt"
	"github.com/metaops/metaops-api/model"
)

func (app *App) CreateDeployment(appID string, gitURL string) (*model.Deployment, error) {
	deployment := &model.Deployment{
		AppID:  appID,
		Status: "initialized",
	}

	if err := app.db.Create(deployment).Error; err != nil {
		return nil, err
	}

	nodes := []model.Node{}
	if err := app.db.Where("app_id = ?", appID).Find(&nodes).Error; err != nil {
		return nil, err
	}

	for _, node := range nodes {
		nodeDeployment := &model.NodeDeployment{
			DeploymentID: deployment.ID,
			NodeID:       node.ID,
			Status:       "initialized",
		}
		if err := app.db.Create(nodeDeployment).Error; err != nil {
			return nil, err
		}
	}

	_, err := app.redis.Do("PUBLISH", appID, app.toJSON(map[string]interface{}{
		"type": "deploy",
		"data": map[string]interface{}{
			"id":     deployment.ID,
			"gitURL": gitURL,
			"config": map[string]string{},
		},
	}))
	if err != nil {
		app.Logger.Println(err)
	}

	return deployment, nil
}

func (app *App) UpdateDeployment(appID, deploymentID, nodeID, status string) (*model.Deployment, error) {
	app.Logger.Println("~~~~~~~~~~~")
	app.Logger.Println(status)
	app.Logger.Println("~~~~~~~~~~~")
	deployment := &model.Deployment{}
	nodeDeployment := &model.NodeDeployment{}

	if err := app.db.First(deployment, "id = ?", deploymentID).Error; err != nil {
		return nil, err
	}

	if err := app.db.Where("deployment_id = ? AND node_id = ?", deploymentID, nodeID).First(nodeDeployment).Error; err != nil {
		return nil, err
	}
	app.Logger.Println("~~~~~~~~~~~")
	fmt.Printf("%+v\n", nodeDeployment)
	app.Logger.Println("~~~~~~~~~~~")
	nodeDeployment.Status = status
	if err := app.db.Save(nodeDeployment).Error; err != nil {
		return nil, err
	}

	switch status {
	case "build_failed":
		//Notificar a todos del build failed y actualizar el status del deployment
		deployment.Status = "build_failed"
		if err := app.db.Save(deployment).Error; err != nil {
			return nil, err
		}
		_, err := app.redis.Do("PUBLISH", appID, app.toJSON(map[string]interface{}{
			"type": "build_failed",
			"data": map[string]string{
				"id": deploymentID,
			},
		}))

		if err != nil {
			app.Logger.Println(err)
		}

		return deployment, nil
	case "build_succeeded":
		// if count(build_succeeded) == count(nodes) -> notificar a todos de build_succeeded
		var succeededCount int
		app.db.Model(&model.NodeDeployment{}).
			Where("deployment_id = ? AND status = ?", deploymentID, "build_succeeded").
			Count(&succeededCount)

		var nodeCount int
		app.db.Model(&model.Node{}).
			Where("app_id = ?", appID).
			Count(&nodeCount)

		if nodeCount == succeededCount {
			deployment.Status = "build_succeeded"
			if err := app.db.Save(deployment).Error; err != nil {
				return nil, err
			}
			_, err := app.redis.Do("PUBLISH", appID, app.toJSON(map[string]interface{}{
				"type": "build_succeeded",
				"data": map[string]string{
					"id": deploymentID,
				},
			}))

			if err != nil {
				app.Logger.Println(err)
			}

			return deployment, nil
		}

	case "release_failed":
		// PAILA :'(
		app.Logger.Printf("FATAL ERROR release_failed [%s, %s]\n", appID, deploymentID)
		deployment.Status = "release_failed"
		if err := app.db.Save(deployment).Error; err != nil {
			return nil, err
		}

		return deployment, nil
	case "release_succeeded":
		deployment.Status = "release_succeded"
		if err := app.db.Save(deployment).Error; err != nil {
			return nil, err
		}

		return deployment, nil
	default:
		app.Logger.Printf("Invalid status [%s]\n", status)
	}

	return deployment, nil
}

func (app *App) toJSON(data map[string]interface{}) string {
	jsonString, err := json.Marshal(data)
	if err != nil {
		app.Logger.Printf("Error: %v\n", err)
		return ""
	}

	return string(jsonString)
}
