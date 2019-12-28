package Providers

import (
	"github.com/goKLC/goKLC"
	"github.com/goKLC/goLog"
)

func init() {
	var app = goKLC.GetApp()

	logger, config := goLog.New()

	fileHandler := goLog.NewFileHandler()
	fileHandler.Path = "logs"
	fileHandler.FileName = "goKLC.log"

	terminalHandler := goLog.NewTerminalHandler()

	config.AddHandler(fileHandler)
	config.AddHandler(terminalHandler)

	app.SetLogger(logger)
}
