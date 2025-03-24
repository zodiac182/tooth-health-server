package main

import (
	"github.com/zodiac182/tooth-health/server/core"
	"github.com/zodiac182/tooth-health/server/db"
	"github.com/zodiac182/tooth-health/server/service"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	db.InitDB()

	service.InitService()

	// 运行服务
	core.RunServer()
}
