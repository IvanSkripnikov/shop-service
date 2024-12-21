package main

import (
	"loyalty_system/httphandler"
	"loyalty_system/logger"
)

func main() {
	logger.Debug("Service starting")
	httphandler.InitHTTPServer()
	logger.Info("Service started")
}
