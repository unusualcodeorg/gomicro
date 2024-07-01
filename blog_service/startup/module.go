package startup

import (
	"context"

	"github.com/unusualcodeorg/gomicro/blog-service/api/auth"
	authMW "github.com/unusualcodeorg/gomicro/blog-service/api/auth/middleware"
	"github.com/unusualcodeorg/gomicro/blog-service/api/author"
	"github.com/unusualcodeorg/gomicro/blog-service/api/blog"
	"github.com/unusualcodeorg/gomicro/blog-service/api/blogs"
	"github.com/unusualcodeorg/gomicro/blog-service/api/editor"
	"github.com/unusualcodeorg/gomicro/blog-service/config"
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
	AuthService auth.Service
	BlogService blog.Service
}

func (m *module) GetInstance() *module {
	return m
}

func (m *module) Controllers() []micro.Controller {
	return []micro.Controller{
		blog.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), m.BlogService),
		blogs.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), blogs.NewService(m.DB, m.Store)),
		author.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), author.NewService(m.DB, m.BlogService)),
		editor.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), editor.NewService(m.DB, m.AuthService)),
	}
}

func (m *module) RootMiddlewares() []network.RootMiddleware {
	return []network.RootMiddleware{
		coreMW.NewErrorCatcher(), // NOTE: this should be the first handler to be mounted
		coreMW.NewNotFound(),
	}
}

func (m *module) AuthenticationProvider() network.AuthenticationProvider {
	return authMW.NewAuthenticationProvider(m.AuthService)
}

func (m *module) AuthorizationProvider() network.AuthorizationProvider {
	return authMW.NewAuthorizationProvider(m.AuthService)
}

func NewModule(context context.Context, env *config.Env, db mongo.Database, store redis.Store, natsClient micro.NatsClient) Module {
	authService := auth.NewService(natsClient)
	blogService := blog.NewService(db, store, authService)

	return &module{
		Context:     context,
		Env:         env,
		DB:          db,
		Store:       store,
		NatsClient:  natsClient,
		AuthService: authService,
		BlogService: blogService,
	}
}
