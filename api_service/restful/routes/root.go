package routes

import (
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/swagger"
	"github.com/profzone/imblock/api_service/restful/routes/peers"
)

var RootRouter = courier.NewRouter(Root{})
var VersionRouter = courier.NewRouter(Version{})

func init() {
	RootRouter.Register(swagger.SwaggerRouter)
	RootRouter.Register(VersionRouter)

	VersionRouter.Register(peers.Router)
}

type Root struct {
	courier.EmptyOperator
}

func (root Root) Path() string {
	return "/api"
}

type Version struct {
	courier.EmptyOperator
}

func (Version) Path() string {
	return "/v0"
}
