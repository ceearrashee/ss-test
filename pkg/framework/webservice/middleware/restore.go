package middleware

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

// RecoveryHandler is a middleware that recovers from panics anywhere in the chain.
func RecoveryHandler() iris.Handler {
	return func(ctx iris.Context) {
		defer handlePanic(ctx)
		ctx.Next()
	}
}

func handlePanic(ctx iris.Context) {
	panicVal := recover()
	if panicVal == nil {
		return
	}

	if ctx.IsStopped() {
		return
	}

	stacktrace := getStacktrace()
	err := fmt.Errorf("panic: %v\n\n%s", panicVal, stacktrace)

	logger := ctx.Application().Logger()
	logger.Error(err)

	correlationID := uuid.NewString()
	problem := createInternalServerErrorProblem(correlationID)

	err = ctx.StopWithProblem(iris.StatusInternalServerError, problem)
	if err != nil {
		logger.Error(err)

		return
	}
}

func createInternalServerErrorProblem(correlationID string) iris.Problem {
	problem := iris.NewProblem()
	problem.Title("Internal Server Error").
		TempKey("correlationID", correlationID).
		Detail("An unexpected error occurred. Please try again later or contact support and tell correlation id.").
		Validate()

	return problem
}

func getStacktrace() string {
	var stacktrace strings.Builder

	for i := 1; ; i++ {
		_, file, line, got := runtime.Caller(i)
		if !got {
			break
		}

		stacktrace.WriteString(fmt.Sprintf("%s:%d\n", file, line)) // nolint: revive // error always nil
	}

	return stacktrace.String()
}
