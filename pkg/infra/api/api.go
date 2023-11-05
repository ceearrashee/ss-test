package api

import (
	"github.com/kataras/iris/v12"
)

// HandleError handles errors and sets HTTP status code.
func HandleError(ctx iris.Context, errorStatus int, err error) {
	if err != nil {
		problem := iris.NewProblem().
			Type(ctx.Request().RequestURI)

		if errorStatus < 500 {
			logAndHandleError(ctx, problem, err, errorStatus)
		} else {
			logAndHandleError(ctx, problem, err, iris.StatusInternalServerError)
		}
	}
}

func logAndHandleError(ctx iris.Context, problem iris.Problem, err error, statusCode int) {
	problem.DetailErr(err).Validate()

	if e := ctx.StopWithProblem(statusCode, problem); e != nil {
		ctx.Application().Logger().Errorf("error while stopping with problem: %v", e)
	}
}
