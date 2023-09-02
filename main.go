package main

import (
	"github.com/asidikfauzi/test-recipes-be-golang/repository/inject"
	"github.com/asidikfauzi/test-recipes-be-golang/routes"
)

func main() {

	routes := routes.InitPackage()
	configInstance := inject.DependencyInjection(inject.InjectData{
		Routes: routes,
	})

	if msg, err := configInstance.InitDB(); err != nil {
		panic("Failed to initialize database: " + err.Error())
	} else {
		println(msg)
	}

	routes.InitRoutes()
}
