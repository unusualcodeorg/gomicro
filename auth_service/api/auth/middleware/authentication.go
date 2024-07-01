package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unusualcodeorg/gomicro/auth-service/api/auth"
	"github.com/unusualcodeorg/gomicro/auth-service/api/user"
	"github.com/unusualcodeorg/gomicro/auth-service/common"
	"github.com/unusualcodeorg/goserve/arch/network"
)

type authenticationProvider struct {
	network.ResponseSender
	common.ContextPayload
	authService auth.Service
	userService user.Service
}

func NewAuthenticationProvider(authService auth.Service, userService user.Service) network.AuthenticationProvider {
	return &authenticationProvider{
		ResponseSender: network.NewResponseSender(),
		ContextPayload: common.NewContextPayload(),
		authService:    authService,
		userService:    userService,
	}
}

func (m *authenticationProvider) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(network.AuthorizationHeader)

		user, keystore, err := m.authService.Authenticate(authHeader)
		if err != nil {
			m.Send(ctx).MixedError(err)
			return
		}

		m.SetUser(ctx, user)
		m.SetKeystore(ctx, keystore)

		ctx.Next()
	}
}
