package rpc

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (rpc *RpcController) NewRouter() *gin.Engine {
	switch rpc.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	if logrus.GetLevel() > logrus.DebugLevel {
		logger := gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: ginLogFormatter,
			Output:    logrus.StandardLogger().Out,
			SkipPaths: []string{"/"},
		})
		router.Use(logger)
	}

	router.Use(gin.RecoveryWithWriter(logrus.StandardLogger().Out))
	router.Use(cors.Default())
	router.Use(gin.Logger())

	rpc.addRouter(router)
	router.Use(BreakerWrapper)

	return router
}

func (rpc *RpcController) addRouter(router *gin.Engine) *gin.Engine {
	router.GET("/health", rpc.Health)

	return router
}
