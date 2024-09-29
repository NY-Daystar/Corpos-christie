package gui

import (
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/model"
	"github.com/NY-Daystar/corpos-christie/user"
	"go.uber.org/zap"
)

// MVC for appliaction
type GUI struct {
	Model      *model.GUIModel
	View       *GUIView
	Controller *GUIController
	Logger     *zap.Logger
}

//		config: Configuration with taxes data of each year
//		user: The user with his data
//		logger: zap logger to log inputs
//	 	path: [string] path to new version of the project
//		display: [optionnal] param to know If we have to display GUI (used for unit tests)
//
// Launch GUI application.
func Start(config *config.Config, user *user.User, logger *zap.Logger, path string, display ...bool) {
	logger.Info("Launch application")

	var model = model.NewModel(config, user, logger)
	var view = NewView(model, logger)
	var controller = NewController(model, view, logger)

	var gui *GUI = &GUI{
		View: view,
	}

	// Launch GUI if bool ok
	if len(display) == 0 {
		gui.View.Start(controller, path)
	}
}
