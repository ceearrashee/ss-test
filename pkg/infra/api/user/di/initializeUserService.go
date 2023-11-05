//go:build wireinject
// +build wireinject

package di

import (
	"github.com/anhro/wire"

	"solid-software.test-task/pkg/domain/user"
	"solid-software.test-task/pkg/infra/db"
)

func InitializeUserService() user.Service {
	wire.Build(
		user.NewUserService, db.GetRawDBConnection,
	)
	return nil
}
