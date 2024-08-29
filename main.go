package main

import (
	"coolv0.1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.CollectRoute(r)
	r.Run(":9000")
}
