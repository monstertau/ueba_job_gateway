package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"job-gateway/api"
	"job-gateway/config"
	"job-gateway/service"
	"log"
)

func main() {
	route := gin.Default()
	apiGroup := route.Group("/api/v1")
	services := &service.Services{
		Profiling: service.NewProfilingJobService(),
	}
	jobHandler := api.NewJobHandler(services)
	jobHandler.MakeHandler(apiGroup)

	err := route.Run(fmt.Sprintf("%s:%d", config.AppConfig.Service.Host, config.AppConfig.Service.Port))
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}
