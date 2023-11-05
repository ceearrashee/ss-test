package middleware

import (
	"time"

	"github.com/kataras/iris/v12"
)

// LoggerHandler returns a middleware handler that logs the duration of each request.
// It logs the request method, request URI and the duration it took to process.
func LoggerHandler() iris.Handler {
	return func(ctx iris.Context) {
		startTime := time.Now().UTC()

		ctx.Next()

		duration := time.Since(startTime)

		go func() {
			ctx.Application().Logger().Infof(
				"%s %s duration: %s",
				ctx.Method(),
				ctx.Request().RequestURI,
				duration.String(),
			)
		}()
	}
}
