package main

import (
	service "github.com/xellossxiv/golangapi/tree/main/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("iamAPI/v1/aralia/setUser", setUserAralia)
	router.POST("iamAPI/v1/hcis/setUser", setUserHcis)
	router.Run("localhost:8081")
}

func setUserAralia(c *gin.Context) {
	service.SetUser(c, "aralia")
}

func setUserHcis(c *gin.Context) {
	service.SetUser(c, "hcis")
}
