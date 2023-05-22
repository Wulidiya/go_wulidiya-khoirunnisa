package main

import (
	"struktur_projek/config"
	"struktur_projek/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
