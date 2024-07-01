package sample

import (
	"github.com/unusualcodeorg/gomicro/microservice2/api/sample/dto"
	"github.com/unusualcodeorg/gomicro/microservice2/api/sample/model"
	"github.com/unusualcodeorg/goserve/arch/mongo"
	"github.com/unusualcodeorg/goserve/arch/network"
	"github.com/unusualcodeorg/goserve/arch/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	FindSample(id primitive.ObjectID) (*model.Sample, error)
}

type service struct {
	network.BaseService
	sampleQueryBuilder mongo.QueryBuilder[model.Sample]
	infoSampleCache    redis.Cache[dto.InfoSample]
}

func NewService(db mongo.Database, store redis.Store) Service {
	return &service{
		BaseService:        network.NewBaseService(),
		sampleQueryBuilder: mongo.NewQueryBuilder[model.Sample](db, model.CollectionName),
		infoSampleCache:    redis.NewCache[dto.InfoSample](store),
	}
}

func (s *service) FindSample(id primitive.ObjectID) (*model.Sample, error) {
	filter := bson.M{"_id": id}

	msg, err := s.sampleQueryBuilder.SingleQuery().FindOne(filter, nil)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
