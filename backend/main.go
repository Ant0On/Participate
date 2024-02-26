package main

import (
	"fmt"

	"backend/controllers"
	"backend/logger"
	"backend/models"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var r *gin.Engine

type arguments struct {
	LogLevel       string
	BindAddress    string
	BindPort       int
	StaticContents string
}

func runServer(args arguments) error {
	logger.InitLogger(args.LogLevel)

	logger.Logger.WithFields(logrus.Fields{
		"args": args,
	}).Info("Given options")

	r = gin.Default()
	r.Use(static.Serve("/", static.LocalFile(args.StaticContents, false)))

	controllers.RegisterRoutes(r)

	if err := r.Run(fmt.Sprintf("%s:%d", args.BindAddress, args.BindPort)); err != nil {
		return err
	}

	return nil
}

func main() {
	args := arguments{
		LogLevel:       "info",
		BindAddress:    "0.0.0.0",
		BindPort:       3000,
		StaticContents: "./frontend/static",
	}

	models.ConnectDatabase()

	if err := runServer(args); err != nil {
		logger.Logger.WithError(err).Fatal("Server exits with error")
	}

	c := cron.New()
	_, err := c.AddFunc("@daily", func() {
		err := models.CheckReservations()
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}
	c.Start()

	select {}
}
