package bootstrap

import (
	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	Db  *gorm.DB
}

type ServerConfig struct {
	Host string
	Port string
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Db = NewPostgresDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseConnection(app.Db)
}