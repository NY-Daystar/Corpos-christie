// gui/controller/main.go
package controller

import (
	"strconv"

	"github.com/NY-Daystar/corpos-christie/gui/model"
	"github.com/NY-Daystar/corpos-christie/gui/view"
	"github.com/NY-Daystar/corpos-christie/tax"
	"github.com/NY-Daystar/corpos-christie/utils"
	"go.uber.org/zap"
)

// GuiController of MVC to handle event and data
type GUIController struct {
	Model  *model.GUIModel
	View   *view.GUIView
	Logger *zap.Logger
}

// NewController instantiate new controller with data model and event to update view
func NewController(model *model.GUIModel, view *view.GUIView, logger *zap.Logger) *GUIController {
	var controller *GUIController = &GUIController{
		Model:  model,
		View:   view,
		Logger: logger,
	}

	controller.setEvents()
	controller.Logger.Info("Events loaded")

	controller.Logger.Info("Launch controller")
	return controller
}

// setEvents Set the events/trigger of gui widgets
func (controller *GUIController) setEvents() {
	controller.View.EntryIncome.OnChanged = func(input string) {
		controller.calculate()
	}
	controller.View.RadioStatus.OnChanged = func(input string) {
		controller.calculate()
	}
	controller.View.SelectChildren.OnChanged = func(input string) {
		controller.calculate()
	}
}

// calculate Get values of gui to calculate tax
func (controller *GUIController) calculate() {
	controller.Model.User.Income = controller.getIncome()
	controller.Model.User.IsInCouple = controller.getStatus()
	controller.Model.User.Children = controller.getChildren()

	result := tax.CalculateTax(controller.Model.User, controller.Model.Config)
	controller.Logger.Sugar().Debugf("Result taxes %#v", result)

	var tax string = utils.ConvertInt64ToString(int64(result.Tax))
	var remainder string = utils.ConvertInt64ToString(int64(result.Remainder))
	var shares string = utils.ConvertInt64ToString(int64(result.Shares))

	// Set data in tax layout
	controller.Model.Tax.Set(tax)
	controller.Model.Remainder.Set(remainder)
	controller.Model.Shares.Set(shares)

	// Set Tax details
	currency, _ := controller.Model.Currency.Get()
	for index := 0; index < controller.Model.LabelsTrancheTaxes.Length(); index++ {
		var taxTranche string = utils.ConvertIntToString(int(result.TaxTranches[index].Tax))
		controller.Model.LabelsTrancheTaxes.SetValue(index, taxTranche+" "+currency)
	}
}

// getIncome Get value of widget entry
func (controller *GUIController) getIncome() int {
	intVal, err := strconv.Atoi(controller.View.EntryIncome.Text)
	if err != nil {
		return 0
	}
	return intVal
}

// getStatus Get value of widget radioGroup
func (controller *GUIController) getStatus() bool {
	return controller.View.RadioStatus.Selected == "Couple"
}

// getChildren get value of widget select
func (controller *GUIController) getChildren() int {
	children, err := strconv.Atoi(controller.View.SelectChildren.Entry.Text)
	if err != nil {
		return 0
	}
	return children
}
