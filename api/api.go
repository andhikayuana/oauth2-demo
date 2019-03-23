package api

import (
	"fmt"

	"github.com/andhikayuana/oauth2-demo/dependency"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Str("package", "api").Logger()

func Start(app *dependency.App) error {

	config := app.GetConfig()

	gin.SetMode(config.Mode)
	r := gin.New()
	r.Use(injectApp(app))

	r.GET("/", actionIndex)
	r.GET("/protected", actionProtected)
	r.GET("/credentials", actionCredentials)
	r.GET("/token", actionToken)

	logger.Info().Msgf("running oauth2-demo at %s:%d", config.Host, config.Port)
	return r.Run(fmt.Sprintf("%s:%d", config.Host, config.Port))
}

func Shutdown() {
	logger.Info().Msg("oauth2-demo shutdown...")
}

func injectApp(app *dependency.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app", app)
		c.Next()
	}
}
