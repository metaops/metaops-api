package app

import (
	"github.com/metaops/metaops-api/model"
)

func (app *App) CreateNode(appID string) (*model.Node, error) {
	node := model.Node{AppID: appID}

	if err := app.db.Create(&node).Error; err != nil {
		app.Logger.Printf("Unable to create node for app *%s*", appID)
		return nil, err
	}

	return &node, nil
}
