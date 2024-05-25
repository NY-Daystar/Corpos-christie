package gui

import (
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/user"
	"go.uber.org/zap"
)

// MVC for appliaction
type GUI struct {
	Model      *GUIModel
	View       *GUIView
	Controller *GUIController
	Logger     *zap.Logger
}

//	config: Configuration with taxes data of each year
//	user: The user with his data
//	logger: zap logger to log inputs
//	display: [optionnal] param to know If we have to display GUI (used for unit tests)
//
// Launch GUI application.
func Start(config *config.Config, user *user.User, logger *zap.Logger, display ...bool) {
	logger.Info("Launch application")

	var model = NewModel(config, user, logger)
	var view = NewView(model, logger)
	var controller = NewController(model, view, logger)

	var gui *GUI = &GUI{
		Model:      model,
		View:       view,
		Controller: controller,
	}

	// Launch GUI if bool ok
	if len(display) == 0 {
		gui.View.Start(controller)
	}
}
