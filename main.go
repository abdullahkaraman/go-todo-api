package main

import (
	"github.com/abdullahkaraman/go-todo-api/app"
	"github.com/abdullahkaraman/go-todo-api/config"
)

func main() {
	config := config.GetConf()
	app := &app.App{}
	app.Start(config)
}
