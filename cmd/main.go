package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/andhikayuana/oauth2-demo/dependency"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"

	"gopkg.in/oauth2.v3/store"

	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v3/manage"

	"github.com/andhikayuana/oauth2-demo/api"
	"github.com/rs/zerolog"

	"github.com/rs/zerolog/log"
)

// flags
var (
	debug = flag.Bool("debug", false, "Run in debug mode")
	host  = flag.String("host", "localhost", "Host for oauth2-demo")
	port  = flag.Int("port", 3000, "Port for this oauth2-demo")
	mode  = gin.ReleaseMode
)

var logger = log.With().Str("package", "main").Logger()

func main() {

	flag.Parse()
	fmt.Println()

	logger.Info().Msg("starting oauth2-demo...")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		logger.Info().Msg("running in debug mode")
		mode = gin.DebugMode
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	manager := manage.NewDefaultManager()
	clientStore := store.NewClientStore()
	srv := server.NewDefaultServer(manager)

	app := dependency.NewApp(
		&dependency.Config{
			Debug: *debug,
			Host:  *host,
			Port:  *port,
			Mode:  mode,
		},
		manager,
		clientStore,
		srv,
	)

	configoAuth2Server(app)

	var apiErrChannel = make(chan error, 1)
	go func() {
		apiErrChannel <- api.Start(app)
	}()

	var signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChannel:
		logger.Info().Msg("got an interrupt, exitting...")
		api.Shutdown()
	case err := <-apiErrChannel:
		if err != nil {
			logger.Error().Err(err).Msg("error while running oauth2-demo, exitting...")
		}
	}
}

func configoAuth2Server(app *dependency.App) {
	manager := app.GetAuthorizationManager()
	clientStore := app.GetClientStore()
	srv := app.GetAuthorizationServer()

	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "secret|enol",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	srv.SetInternalErrorHandler(func(err error) (res *errors.Response) {
		logger.Info().Msgf("Internal Error : %s", err.Error())
		return
	})
	srv.SetResponseErrorHandler(func(res *errors.Response) {
		logger.Info().Msgf("Response Error, %s", res.Error.Error())
	})
}
