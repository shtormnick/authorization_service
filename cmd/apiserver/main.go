package main

import (
	"github.com/shitikovkirill/auth-service/internal/app/apiserver"
	"github.com/shitikovkirill/auth-service/internal/app/config"
	"github.com/shitikovkirill/auth-service/internal/app/dependencies/logger"
)

func main() {
	logs := logger.Get()
	config := config.NewConfig()
	if err := apiserver.Start(config); err != nil {
		logs.Error(err)
	}
}
