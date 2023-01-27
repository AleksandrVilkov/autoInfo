package main

import (
	"awesomeProject/internal/app"
	"os"
)

func main() {
	if os.Setenv("PATH_LOG", "/home/vilkov/GolandProjects/autoInfo/internal/logs/") == nil {
		app.Run()
	}

}
