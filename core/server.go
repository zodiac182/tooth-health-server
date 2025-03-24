package core

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/zodiac182/tooth-health/server/core/logger"
	"github.com/zodiac182/tooth-health/server/router"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func RunServer() {
	r := router.Routers()
	s := initServer(":8080", r)

	logger.Error("failed to start server: %s", s.ListenAndServe())
}
