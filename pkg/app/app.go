package app

import (
	"fmt"

	"solid-software.test-task/pkg/app/di"
	"solid-software.test-task/pkg/framework/webservice"
	"solid-software.test-task/pkg/infra/api/healthz"
	"solid-software.test-task/pkg/infra/api/token"
	"solid-software.test-task/pkg/infra/api/user"
)

var (
	// ErrServerClosed is returned when the web service is closed.
	ErrServerClosed = webservice.ErrServerClosed
)

// Run bootstraps and starts the web service.
// It first initializes a new web service instance,
// then registers the necessary endpoints and finally starts the service.
func Run() error {
	service := di.InitializeNewWebService()
	service.RegisterEndpoints(token.NewTokenAPI(), user.NewUserAPI(), healthz.NewHealthzAPI())

	err := service.Run()
	if err != nil {
		return fmt.Errorf("failed to run the service: %w", err)
	}

	return nil
}
