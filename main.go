package main

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func main() {
	kill := make(chan bool)
	e := echo.New()
	e.Debug = true
	e.HideBanner = true

	e.Start("8001")

	log.Info("initializing project in port 8001")

	<-kill
}
