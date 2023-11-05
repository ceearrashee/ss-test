package webservice

import (
	"context"
	"fmt"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"

	"solid-software.test-task/pkg/framework/config"
	"solid-software.test-task/pkg/framework/webservice/di"
	"solid-software.test-task/pkg/framework/webservice/interfaces"
	"solid-software.test-task/pkg/framework/webservice/middleware"
	"solid-software.test-task/pkg/framework/webservice/route"
)

type (
	webService struct {
		application *iris.Application
		config      config.Config
	}
)

var (
	// ErrServerClosed is returned when the web service is closed.
	ErrServerClosed = iris.ErrServerClosed
)

// New returns a new instance of the web service.
// It initializes the web application if not already done so.
func New(cfg config.Config) interfaces.WebService {
	return &webService{
		application: initializeWebApp(),
		config:      cfg,
	}
}

func initializeWebApp() *iris.Application {
	app := iris.New()
	app.SetRegisterRule(iris.RouteError)
	setUpMiddleware(app)
	disposeWebAppOnInterrupt(app)

	return app
}

func setUpMiddleware(app *iris.Application) {
	app.UseGlobal(
		middleware.RecoveryHandler(), middleware.LoggerHandler(),
		// TODO: implement cors if needed in your infrastructure
		// cors.New().Handler()
	)
}

func disposeWebAppOnInterrupt(app *iris.Application) {
	iris.RegisterOnInterrupt(
		func() {
			if err := app.Shutdown(context.TODO()); err != nil {
				log.Printf("web application shutdown error: %v", err.Error())
			}

			log.Println("web application graceful shutdown")
		},
	)
}

// RegisterEndpoints registers the endpoints for the web service.
func (w *webService) RegisterEndpoints(routes ...route.Route) {
	regularRoute, protectedRoute := createRoutes(w)

	for _, r := range routes {
		if r.IsProtected() {
			r.InitRoutes(protectedRoute)
		} else {
			r.InitRoutes(regularRoute)
		}
	}
}

func createRoutes(w *webService) (router.Party, router.Party) {
	regularRoute := w.application.Party("/api/v1").PartyConfigure("/")
	protectedRoute := regularRoute.Party("/")

	initializeJwt(regularRoute, protectedRoute)

	return regularRoute, protectedRoute
}

func initializeJwt(regularRoute, protectedRoute router.Party) {
	jwtService := di.InitializeJWTService()

	regularRoute.ConfigureContainer(
		func(container *router.APIContainer) {
			container.RegisterDependency(jwtService)
		},
	)

	protectedRoute.ConfigureContainer(
		func(container *router.APIContainer) {
			container.Use(jwtService.GetHandler())
		},
	)
}

// Run starts the web service on port 8080.
// It returns an error if there is issue in listening on the port.
func (w *webService) Run() error {
	return w.tryListen(
		fmt.Sprintf("%s:%d", w.config.GetString("webService.host"), w.config.GetUint("webService.port")),
	)
}

func (w *webService) tryListen(addr string) error {
	if err := w.application.Listen(addr); err != nil {
		return fmt.Errorf("listen error: %w", err)
	}

	return nil
}
