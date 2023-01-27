package app

import (
	"awesomeProject/internal/dto"
	"awesomeProject/pkg/logger"
	myServer "awesomeProject/pkg/server"
	"gopkg.in/yaml.v3"
	"os"
)

func Run() {
	server := new(myServer.Server)
	newLogger, _ := logger.NewLogger("httpServer")
	in, _ := os.ReadFile("/home/vilkov/GolandProjects/autoInfo/internal/config/serverConfig.yaml")
	var config dto.ServerConfig
	err := yaml.Unmarshal(in, &config)

	if err != nil {
		newLogger.WriteError("Error generate Server config")
	}

	newLogger.WriteInfo("get server configuration successfully")
	if err := server.Run(config.Port); err != nil {
		newLogger.WriteError("Error running http server on port " + config.Port + ": " + err.Error())
	}
}
