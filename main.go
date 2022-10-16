package main

import (
	"go-test-api/database"
	"go-test-api/models"
	"go-test-api/routes"
)

func main() {
	database.Init()
	router := routes.Init()
	// test()a
	_ = router.Run(":8080")

}

func test() {
	models.CreateProject(&models.Project{
		Name:        "Proyecto prueba",
		Description: "descripcion",
		StartDate:   "2022-08-08",
		EndDate:     "2022-11-25",
		Budget:      234.34,
	})
}
