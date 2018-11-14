package api_service

import (
	"github.com/profzone/imblock/core"
	"github.com/profzone/imblock/core/message_bus"
	"github.com/profzone/imblock/core/castle"
	"github.com/johnnyeven/libtools/courier/transport_http"
	"github.com/profzone/imblock/api_service/restful/routes"
	"github.com/sirupsen/logrus"
	"github.com/profzone/imblock/global"
	"github.com/johnnyeven/libtools/courier"
)

var apiHttp *ApiHttpServiceBootstrap

type ApiHttpServiceBootstrap struct {
	server transport_http.ServeHTTP
}

func NewApiHttpServiceBootstrap() castle.Service {
	if apiHttp != nil {
		return apiHttp
	}

	apiHttp = &ApiHttpServiceBootstrap{
		server: transport_http.ServeHTTP{
			Name:     "service-api-imblock",
			Port:     global.Config.ServerPort,
			WithCORS: true,
		},
	}

	return apiHttp
}

func (s *ApiHttpServiceBootstrap) Messages() []message_bus.MessageHandler {
	return nil
}

func (s *ApiHttpServiceBootstrap) Protocols() []core.ProtocolHandler {
	return nil
}

func (s *ApiHttpServiceBootstrap) Routes() []*courier.Router {
	return nil
}

func (s *ApiHttpServiceBootstrap) Start() error {
	logrus.Info("[ApiHttpServiceBootstrap] started.")
	castle.GeneralCastle.Iterator(func(service castle.Service) error {
		routers := service.Routes()
		for _, router := range routers {
			routes.VersionRouter.Register(router)
		}
		return nil
	}, true)
	return s.server.Serve(routes.RootRouter)
}

func (s *ApiHttpServiceBootstrap) Stop() error {
	logrus.Info("[ApiHttpServiceBootstrap] stopped.")
	return s.server.Stop()
}
