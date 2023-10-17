package main

import (
	"fmt"

	"Participate/backend/models"
	"Participate/backend/utils"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

var logLevelMap = map[string]logrus.Level{
	"trace": logrus.TraceLevel,
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

type arguments struct {
	LogLevel       string
	BindAddress    string
	BindPort       int
	StaticContents string
}

var r *gin.Engine

func runServer(args arguments) error {
	level, ok := logLevelMap[args.LogLevel]
	if !ok {
		return fmt.Errorf("invalid log level: %s", args.LogLevel)
	}
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"args": args,
	}).Info("Given options")

	r = gin.Default()

	r.Use(static.Serve("/", static.LocalFile(args.StaticContents, false)))

	r.GET("/api/v1/hello", func(c *gin.Context) {
		c.String(200, `{"message":"hello, hello, hello"}`)
	})

	utils.RegisterRoutes(r)

	if err := r.Run(fmt.Sprintf("%s:%d", args.BindAddress, args.BindPort)); err != nil {
		return err
	}

	return nil
}

func main() {
	models.ConnectDatabase()

	args := arguments{
		LogLevel:       "info",
		BindAddress:    "0.0.0.0",
		BindPort:       9080,
		StaticContents: "./frontend/static",
	}

	if err := runServer(args); err != nil {
		logger.WithError(err).Fatal("Server exits with error")
	}

	// TODO To be deleted, left as an example
	//r.GET("/users/:id", func(c *gin.Context) {
	//	var user Osoba
	//
	//	if err := db.Find(&user).Error; err != nil {
	//		err.Error()
	//		c.JSON(500, gin.H{"error": "Failed to fetch users"})
	//		return
	//	}
	//
	//	c.JSON(200, user)
	//})
}
