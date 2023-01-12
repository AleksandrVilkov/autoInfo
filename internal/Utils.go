package internal

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func GetConfig() *Config {
	in, _ := os.ReadFile("/home/vilkov/GolandProjects/awesomeProject/resources/config.yaml")
	var config Config
	err := yaml.Unmarshal(in, &config)
	if err == nil {
		fmt.Println(config)
	} else {
		fmt.Println(err)
	}
	return &config
}
