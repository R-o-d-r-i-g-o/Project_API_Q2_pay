package main

import (
	"form/repository"
	all "form/routes"

	"github.com/gin-gonic/gin"
)

func init() {

	repository.ConnectToDataBase()
}

func main() {

	router := gin.Default()

	all.AvaibleRoutes(router)

	router.Run("localhost:8000")
}
