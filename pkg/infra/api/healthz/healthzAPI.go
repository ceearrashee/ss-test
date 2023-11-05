package healthz

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"

	"solid-software.test-task/pkg/framework/webservice/route"
	"solid-software.test-task/pkg/infra/api"
)

type (
	healthz struct{}
)

// NewHealthzAPI creates a new instance of HealthzAPI.
func NewHealthzAPI() route.Route {
	return &healthz{}
}

// IsProtected indicates if the TokenAPI is a protected route.
func (*healthz) IsProtected() bool {
	return false
}

// InitRoutes inits the healthz API routes.
func (*healthz) InitRoutes(party router.Party) {
	party.Get("/healthz", handlerHealthz)
}

func handlerHealthz(irisContext iris.Context) {
	_, err := irisContext.WriteString("Ok")
	if err != nil {
		api.HandleError(irisContext, iris.StatusInternalServerError, err)
	}
}
