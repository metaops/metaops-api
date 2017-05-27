package app

import (
	"github.com/metaops/metaops-api/model"
)

func (app *App) CreateApp(name string) (*model.App, error) {
	userApp := model.App{Name: name}

	if err := app.db.Create(&userApp).Error; err != nil {
		app.Logger.Printf("Unable to create app *%s*", name)
		return nil, err
	}

	return &userApp, nil
}
