package main

import (
	"miniproject/configs"
	"miniproject/routes"
)

func main() {
	configs.ConnectDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1312"))
}
