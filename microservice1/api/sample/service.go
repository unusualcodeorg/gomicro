package sample

import (
	"github.com/unusualcodeorg/gomicro/microservice1/api/sample/dto"
	"github.com/unusualcodeorg/gomicro/microservice1/api/sample/message"
	"github.com/unusualcodeorg/gomicro/microservice1/api/sample/model"
	"github.com/unusualcodeorg/goserve/arch/micro"
	"github.com/unusualcodeorg/goserve/arch/mongo"
	"github.com/unusualcodeorg/goserve/arch/network"
	"github.com/unusualcodeorg/goserve/arch/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	FindSample(id primitive.ObjectID) (*model.Sample, error)
	GetSampleMessage(data *message.SampleMessage) (*message.SampleMessage, error)
}

type service struct {
	network.BaseService
	sampleQueryBuilder   mongo.QueryBuilder[model.Sample]
	infoSampleCache      redis.Cache[dto.InfoSample]
	sampleRequestBuilder micro.RequestBuilder[message.SampleMessage]
}

func NewService(db mongo.Database, store redis.Store, natsClient micro.NatsClient) Service {
	return &service{
		BaseService:          network.NewBaseService(),
		sampleQueryBuilder:   mongo.NewQueryBuilder[model.Sample](db, model.CollectionName),
		infoSampleCache:      redis.NewCache[dto.InfoSample](store),
		sampleRequestBuilder: micro.NewRequestBuilder[message.SampleMessage](natsClient, "microservice2.sample.ping"),
	}
}

func (s *service) GetSampleMessage(data *message.SampleMessage) (*message.SampleMessage, error) {
	return s.sampleRequestBuilder.Request(data).Nats()
}

func (s *service) FindSample(id primitive.ObjectID) (*model.Sample, error) {
	filter := bson.M{"_id": id}

	msg, err := s.sampleQueryBuilder.SingleQuery().FindOne(filter, nil)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
