package startup

import (
	"context"

	"github.com/unusualcodeorg/gomicro/auth-service/api/auth"
	authMW "github.com/unusualcodeorg/gomicro/auth-service/api/auth/middleware"
	"github.com/unusualcodeorg/gomicro/auth-service/api/user"
	"github.com/unusualcodeorg/gomicro/auth-service/config"
	"github.com/unusualcodeorg/goserve/arch/micro"
	coreMW "github.com/unusualcodeorg/goserve/arch/middleware"
	"github.com/unusualcodeorg/goserve/arch/mongo"
	"github.com/unusualcodeorg/goserve/arch/network"
	"github.com/unusualcodeorg/goserve/arch/redis"
)

type Module micro.Module[module]

type module struct {
	Context     context.Context
	Env         *config.Env
	DB          mongo.Database
	Store       redis.Store
	NatsClient  micro.NatsClient
	UserService user.Service
	AuthService auth.Service
}

func (m *module) GetInstance() *module {
	return m
}

func (m *module) Controllers() []micro.Controller {
	return []micro.Controller{
		auth.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), m.AuthService),
		user.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), m.UserService),
	}
}

func (m *module) RootMiddlewares() []network.RootMiddleware {
	return []network.RootMiddleware{
		coreMW.NewErrorCatcher(), // NOTE: this should be the first handler to be mounted
		authMW.NewKeyProtection(m.AuthService),
		coreMW.NewNotFound(),
	}
}

func (m *module) AuthenticationProvider() network.AuthenticationProvider {
	return authMW.NewAuthenticationProvider(m.AuthService, m.UserService)
}

func (m *module) AuthorizationProvider() network.AuthorizationProvider {
	return authMW.NewAuthorizationProvider()
}

func NewModule(context context.Context, env *config.Env, db mongo.Database, store redis.Store, natsClient micro.NatsClient) Module {
	userService := user.NewService(db)
	authService := auth.NewService(db, env, userService)
	return &module{
		Context:     context,
		Env:         env,
		DB:          db,
		Store:       store,
		NatsClient:  natsClient,
		UserService: userService,
		AuthService: authService,
	}
}
