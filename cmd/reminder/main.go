package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudSarkar/reminder/routes"
)

func main() {
	r := gin.Default()

	routes.RegisterAPIRoute(r)

	r.Run(":8000")
}
