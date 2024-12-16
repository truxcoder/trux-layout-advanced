//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"github.com/truxcoder/trux-layout-advanced/internal/handler"
	"github.com/truxcoder/trux-layout-advanced/internal/repository"
	"github.com/truxcoder/trux-layout-advanced/internal/server"
	"github.com/truxcoder/trux-layout-advanced/internal/service"
	"github.com/truxcoder/trux-layout-advanced/pkg/app"
	"github.com/truxcoder/trux-layout-advanced/pkg/jwt"
	"github.com/truxcoder/trux-layout-advanced/pkg/log"
	"github.com/truxcoder/trux-layout-advanced/pkg/server/http"
	"github.com/truxcoder/trux-layout-advanced/pkg/sid"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(
	httpServer *http.Server,
	job *server.Job,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
