package interfaces

import (
	"solid-software.test-task/pkg/framework/webservice/route"
)

type (
	// WebService provides methods for configuring web service.
	// It allows registering endpoints and control their execution.
	WebService interface {
		// RegisterEndpoints registers routes as endpoint or endpoint group to the web service.
		RegisterEndpoints(routes ...route.Route)
		// Run starts the web service.
		// It will continue the execution until the application is stopped manually,
		// encounters an error or gets terminated.
		// Returns an error if any error occurred during the service execution.
		Run() error
	}
)
