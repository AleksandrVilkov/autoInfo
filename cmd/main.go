package main

import (
	"awesomeProject/internal"
	"fmt"
)

func main() {
	config := internal.GetConfig()
	fmt.Println(config)
}
