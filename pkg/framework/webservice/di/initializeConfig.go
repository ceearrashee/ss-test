//go:build wireinject
// +build wireinject

package di

import (
	"github.com/anhro/wire"

	"solid-software.test-task/pkg/framework/config"
)

func InitializeConfig() config.Config {
	wire.Build(config.NewConfig)
	return nil
}
