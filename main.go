package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"job-gateway/api"
	"job-gateway/config"
	"job-gateway/importer"
	"job-gateway/service"
	"log"
)

func main() {
	log.Println("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
	appConfig, err := config.LoadFile(config.DefaultConfigFilePath)
	if err != nil {
		log.Fatalf("failed to parse configuration file %s: %v", config.DefaultConfigFilePath, err)
	}
	config.AppConfig = appConfig

	route := gin.Default()
	apiGroup := route.Group("/api/v1")

	contentImporter, err := importer.NewImporter(appConfig.ContentPath)
	if err != nil {
		log.Fatalf("error in init importer: %v", err)
	}
	importer.GlobalImporter = contentImporter
	go contentImporter.Sync()

	services := &service.Services{
		Profiling: service.NewProfilingJobService(),
	}
	jobHandler := api.NewJobHandler(services)
	jobHandler.MakeHandler(apiGroup)
	err = route.Run(fmt.Sprintf("%s:%d", config.AppConfig.Service.Host, config.AppConfig.Service.Port))
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}
