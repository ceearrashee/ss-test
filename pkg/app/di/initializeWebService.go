//go:build wireinject
// +build wireinject

package di

import (
	"github.com/anhro/wire"

	"solid-software.test-task/pkg/framework/config"
	"solid-software.test-task/pkg/framework/webservice"
	"solid-software.test-task/pkg/framework/webservice/interfaces"
)

func InitializeNewWebService() interfaces.WebService {
	wire.Build(webservice.New, config.NewConfig)
	return nil
}
