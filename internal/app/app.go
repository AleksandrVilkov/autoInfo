package app

import (
	"awesomeProject/internal/entity"
	"awesomeProject/pkg/logger"
	myServer "awesomeProject/pkg/server"
	"gopkg.in/yaml.v3"
	"os"
)

func Run() {
	server := new(myServer.Server)
	newLogger, _ := logger.NewLogger("httpServer", "/home/vilkov/GolandProjects/awesomeProject/internal/logs/")
	in, _ := os.ReadFile("/home/vilkov/GolandProjects/awesomeProject/internal/config/serverConfig.yaml")
	var config entity.ServerConfig
	err := yaml.Unmarshal(in, &config)

	if err != nil {
		newLogger.WriteError("Error generate Server config")
	}

	newLogger.WriteInfo("get server configuration successfully")
	newLogger.WriteInfo("get server configuration successfully2")
	if err := server.Run(config.Port); err != nil {
		newLogger.WriteError("Error running http server on port " + config.Port + ": " + err.Error())
	}
}
