package token

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"

	"solid-software.test-task/pkg/framework/webservice/jwt"
	"solid-software.test-task/pkg/framework/webservice/route"
	"solid-software.test-task/pkg/infra/api"
)

type (
	tokenAPI struct{}
)

// NewTokenAPI creates a new instance of TokenAPI.
func NewTokenAPI() route.Route {
	return &tokenAPI{}
}

// IsProtected indicates if the TokenAPI is a protected route.
func (*tokenAPI) IsProtected() bool {
	return false
}

// InitRoutes inits the token API routes.
func (*tokenAPI) InitRoutes(party router.Party) {
	party.Party("/token").ConfigureContainer(
		func(container *router.APIContainer) {
			container.Get("/generate", generateToken)
		},
	)
}

func generateToken(irisContext iris.Context, jwtService jwt.Service) {
	sampleClaim := jwt.SampleClaim{
		Username: gofakeit.Username(),
	}

	token, err := jwtService.GetToken(sampleClaim)
	if err != nil {
		handleTokenOperationError(irisContext, err, "failed to sign token")
		return
	}

	_, err = irisContext.Write(token)
	if err != nil {
		handleTokenOperationError(irisContext, err, "failed to write token")
	}
}

func handleTokenOperationError(irisContext iris.Context, err error, message string) {
	irisContext.Application().Logger().Errorf("%s: %v", message, err)
	api.HandleError(irisContext, iris.StatusInternalServerError, err)
}
