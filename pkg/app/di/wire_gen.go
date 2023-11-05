// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/anhro/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"solid-software.test-task/pkg/framework/config"
	"solid-software.test-task/pkg/framework/webservice"
	"solid-software.test-task/pkg/framework/webservice/interfaces"
)

// Injectors from initializeWebService.go:

func InitializeNewWebService() interfaces.WebService {
	configConfig := config.NewConfig()
	webService := webservice.New(configConfig)
	return webService
}
