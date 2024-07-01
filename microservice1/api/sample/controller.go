package sample

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unusualcodeorg/gomicro/microservice1/api/sample/dto"
	"github.com/unusualcodeorg/gomicro/microservice1/api/sample/message"
	coredto "github.com/unusualcodeorg/goserve/arch/dto"
	"github.com/unusualcodeorg/goserve/arch/micro"
	"github.com/unusualcodeorg/goserve/arch/network"
	"github.com/unusualcodeorg/goserve/common"
	"github.com/unusualcodeorg/goserve/utils"
)

type controller struct {
	micro.BaseController
	common.ContextPayload
	service Service
}

func NewController(
	authMFunc network.AuthenticationProvider,
	authorizeMFunc network.AuthorizationProvider,
	service Service,
) micro.Controller {
	return &controller{
		BaseController: micro.NewBaseController("/sample", authMFunc, authorizeMFunc),
		ContextPayload: common.NewContextPayload(),
		service:        service,
	}
}

func (c *controller) MountNats(group micro.NatsGroup) {
	group.AddEndpoint("ping", micro.NatsHandlerFunc(c.pingHandler))
}

func (c *controller) MountRoutes(group *gin.RouterGroup) {
	group.GET("/ping", c.getEchoHandler)
	group.GET("/service/ping", c.getServicePingHandler)
	group.GET("/id/:id", c.getSampleHandler)
}

func (c *controller) pingHandler(req micro.NatsRequest) {
	fmt.Println(string(req.Data()))
	msg := message.NewSampleMessage("from", "microservice1")
	micro.Respond(req, msg, nil)
}

func (c *controller) getEchoHandler(ctx *gin.Context) {
	c.Send(ctx).SuccessMsgResponse("pong!")
}

func (c *controller) getServicePingHandler(ctx *gin.Context) {
	msg := message.NewSampleMessage("from", "microservice1")
	received, err := micro.Request(c.Context(), "microservice2.sample.ping", msg, message.EmptySampleMessage())
	if err != nil {
		c.Send(ctx).MixedError(err)
		return
	}
	c.Send(ctx).SuccessDataResponse("success", received)
}

func (c *controller) getSampleHandler(ctx *gin.Context) {
	mongoId, err := network.ReqParams(ctx, coredto.EmptyMongoId())
	if err != nil {
		c.Send(ctx).BadRequestError(err.Error(), err)
		return
	}

	sample, err := c.service.FindSample(mongoId.ID)
	if err != nil {
		c.Send(ctx).NotFoundError("sample not found", err)
		return
	}

	data, err := utils.MapTo[dto.InfoSample](sample)
	if err != nil {
		c.Send(ctx).InternalServerError("something went wrong", err)
		return
	}

	c.Send(ctx).SuccessDataResponse("success", data)
}
