package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("iamAPI/v1/aralia/setUser", setUserAralia)
	router.POST("iamAPI/v1/hcis/setUser", setUserHcis)
	router.Run("0.0.0.0:8081")
}

func setUserAralia(c *gin.Context) {
	SetUser(c, "aralia")
}

func setUserHcis(c *gin.Context) {
	SetUser(c, "hcis")
}
