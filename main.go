package main

import (
	"fmt"

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

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile(args.StaticContents, false)))

	if err := r.Run(fmt.Sprintf("%s:%d", args.BindAddress, args.BindPort)); err != nil {
		return err
	}

	return nil
}

func main() {
	args := arguments{
		LogLevel:       "info",
		BindAddress:    "0.0.0.0",
		BindPort:       9080,
		StaticContents: "./frontend/static",
	}

	if err := runServer(args); err != nil {
		logger.WithError(err).Fatal("Server exits with error")
	}
	//models.ConnectDataBase()
	//
	//r := gin.Default()
	//
	//public := r.Group("/api")
	//
	//public.POST("/register", controllers.Register)
	//public.POST("/login", controllers.Login)
	//
	//protected := r.Group("/api/admin")
	//protected.Use(middlewares.JwtAuthMiddleware())
	//protected.GET("/customer", controllers.CurrentCustomer)

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

	//if err := r.Run(":8080"); err != nil {
	//	log.Fatal(err)
	//} // listen and serve on 0.0.0.0:8080
}
