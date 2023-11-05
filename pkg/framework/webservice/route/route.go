package route

import (
	"github.com/kataras/iris/v12/core/router"
)

type (
	// Route is an interface that provides specification for routing-related operations.
	Route interface {
		// IsProtected checks if the route is need to be protected by authentication.
		IsProtected() bool
		// InitRoutes initializes endpoint or group of endpoints.
		InitRoutes(p router.Party)
	}
)
