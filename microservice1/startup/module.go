package startup

import (
	"context"

	"github.com/unusualcodeorg/gomicro/microservice1/api/sample"
	"github.com/unusualcodeorg/gomicro/microservice1/config"
	"github.com/unusualcodeorg/goserve/arch/micro"
	coreMW "github.com/unusualcodeorg/goserve/arch/middleware"
	"github.com/unusualcodeorg/goserve/arch/mongo"
	"github.com/unusualcodeorg/goserve/arch/network"
	"github.com/unusualcodeorg/goserve/arch/redis"
)

type Module micro.Module[module]

type module struct {
	Context    context.Context
	Env        *config.Env
	DB         mongo.Database
	Store      redis.Store
	NatsClient micro.NatsClient
}

func (m *module) GetInstance() *module {
	return m
}

func (m *module) Controllers() []micro.Controller {
	return []micro.Controller{
		sample.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), sample.NewService(m.DB, m.Store, m.NatsClient)),
	}
}

func (m *module) RootMiddlewares() []network.RootMiddleware {
	return []network.RootMiddleware{
		coreMW.NewErrorCatcher(),
		coreMW.NewNotFound(),
	}
}

func (m *module) AuthenticationProvider() network.AuthenticationProvider {
	// TODO
	return nil
}

func (m *module) AuthorizationProvider() network.AuthorizationProvider {
	// TODO
	return nil
}

func NewModule(context context.Context, env *config.Env, db mongo.Database, store redis.Store, natsClient micro.NatsClient) Module {
	return &module{
		Context:    context,
		Env:        env,
		DB:         db,
		Store:      store,
		NatsClient: natsClient,
	}
}
