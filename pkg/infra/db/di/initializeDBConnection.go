//go:build wireinject
// +build wireinject

package di

import (
	"github.com/anhro/wire"

	"solid-software.test-task/pkg/infra/db/initializer"
	"solid-software.test-task/pkg/infra/db/interfaces"
)

func InitializeNewDBConnection() interfaces.Connection {
	wire.Build(initializer.InitDBConnection)
	return nil
}
