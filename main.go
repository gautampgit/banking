package main

import (
	"github.com/gautampgit/banking/logger"

	"github.com/gautampgit/banking/app"
)

func main() {
	logger.Info("starting the Application")
	app.Start()
}
