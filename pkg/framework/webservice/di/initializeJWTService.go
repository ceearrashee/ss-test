//go:build wireinject
// +build wireinject

package di

import (
	"github.com/anhro/wire"

	"solid-software.test-task/pkg/framework/config"
	"solid-software.test-task/pkg/framework/webservice/jwt"
)

func InitializeJWTService() jwt.Service {
	wire.Build(jwt.NewJWT, config.NewConfig)
	return nil
}
