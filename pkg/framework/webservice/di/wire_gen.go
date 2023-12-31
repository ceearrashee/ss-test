// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/anhro/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"solid-software.test-task/pkg/framework/config"
	"solid-software.test-task/pkg/framework/webservice/jwt"
)

// Injectors from initializeConfig.go:

func InitializeConfig() config.Config {
	configConfig := config.NewConfig()
	return configConfig
}

// Injectors from initializeJWTService.go:

func InitializeJWTService() jwt.Service {
	configConfig := config.NewConfig()
	jwtService := jwt.NewJWT(configConfig)
	return jwtService
}
