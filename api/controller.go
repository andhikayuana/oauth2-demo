package api

import (
	"github.com/andhikayuana/oauth2-demo/dependency"
	"github.com/gin-gonic/gin"
)

func actionIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"name":    "oauth2-demo",
		"version": "1.0.0",
	})
}

func actionProtected(c *gin.Context) {
	c.String(200, "protected here")
}

func actionCredentials(c *gin.Context) {
	// c.String(200, "generate client_id and client_secret here")

	// clientID :=

	app := c.MustGet("app").(*dependency.App)

	c.JSON(200, gin.H{
		"mode":  app.GetConfig().Mode,
		"debug": app.GetConfig().Debug,
	})
}

func actionToken(c *gin.Context) {
	c.String(200, "generate token here")
}
