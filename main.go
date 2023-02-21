package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("", HelloServer)
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

func HelloServer(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, &JsonMessage{"200", "Success", "User Processed to Sailpoint"})
}
