package middleware

import (
	"context"
	"errors"

	"github.com/kataras/iris/v12"

	"solid-software.test-task/pkg/framework/ctxutils"
	"solid-software.test-task/pkg/framework/webservice/jwt"
	"solid-software.test-task/pkg/infra/api"
)

// AuthHandler returns Iris middleware handler that authorizes user by JWT token.
// If the token is valid it proceeds to set the context with username.
// If the token is invalid an Unauthorized HTTP error is returned to the client.
// It uses jwt.Service to authenticate and authorize the client.
func AuthHandler(service jwt.Service) iris.Handler {
	return func(irisCtx iris.Context) {
		sampleClaim, err := service.VerifyTokenAndReturnClaim(irisCtx)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenInvalid) {
				api.HandleError(irisCtx, iris.StatusUnauthorized, err)
			} else {
				api.HandleError(irisCtx, iris.StatusInternalServerError, err)
			}

			return
		}

		setContextWithUsername(irisCtx, sampleClaim.Username)
	}
}

func setContextWithUsername(irisCtx iris.Context, username string) {
	ctx := context.WithValue(irisCtx.Request().Context(), ctxutils.UsernameContextKey, username)
	irisCtx.Values().Set(string(ctxutils.AppContextKey), ctx)
}
