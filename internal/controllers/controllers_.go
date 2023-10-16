package controllers

type Controller interface {
	Build()
}

func BuildController() {
	var controllers []Controller
	controllers = append(controllers, NewAuthController())
	controllers = append(controllers, NewCategoryController())
	for _, controller := range controllers {
		controller.Build()
	}
}
