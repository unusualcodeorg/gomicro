package auth

import (
	"github.com/unusualcodeorg/gomicro/blog-service/api/auth/message"
	"github.com/unusualcodeorg/goserve/arch/micro"
	"github.com/unusualcodeorg/goserve/arch/network"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Authenticate(token string) (*message.User, error)
	Authorize(user *message.User, roles ...string) error
	FindUserPublicProfile(userId primitive.ObjectID) (*message.User, error)
}

type service struct {
	network.BaseService
}

func NewService(natsClient micro.NatsClient) Service {
	return &service{
		BaseService: network.NewBaseService(),
	}
}

func (s *service) Authenticate(token string) (*message.User, error) {
	return nil, nil
}

func (s *service) Authorize(user *message.User, roles ...string) error {
	return nil
}

func (s *service) FindUserPublicProfile(userId primitive.ObjectID) (*message.User, error) {
	return nil, nil
}
