package dependency

import (
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

type App struct {
	config               *Config
	authorizationManager *manage.Manager
	clientStore          *store.ClientStore
	authorizationServer  *server.Server
}

func NewApp(
	config *Config,
	authorizationManager *manage.Manager,
	clientStore *store.ClientStore,
	authorizationServer *server.Server) *App {
	return &App{
		config:               config,
		authorizationManager: authorizationManager,
		clientStore:          clientStore,
		authorizationServer:  authorizationServer,
	}
}

func (app *App) GetConfig() *Config {
	return app.config
}

func (app *App) GetAuthorizationManager() *manage.Manager {
	return app.authorizationManager
}

func (app *App) GetClientStore() *store.ClientStore {
	return app.clientStore
}

func (app *App) GetAuthorizationServer() *server.Server {
	return app.authorizationServer
}
