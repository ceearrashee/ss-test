package main

import (
	"errors"

	"solid-software.test-task/pkg/app"
	"solid-software.test-task/pkg/framework/config"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		if !errors.Is(err, app.ErrServerClosed) {
			panic(err)
		}
	}
}
