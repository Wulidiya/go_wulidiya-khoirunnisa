package main

import (
	"project_structure1/config"
	"project_structure1/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
