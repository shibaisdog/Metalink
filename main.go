package main

import (
	"metalink-apiserver/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8083")
}
