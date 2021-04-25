package routes

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	limits "github.com/gin-contrib/size"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/pyprism/Hiren-Finder/config"
	"go.uber.org/zap"
	"os"
	"time"
)

func NewRouter() *gin.Engine {
	var logger *zap.Logger
	if os.Getenv("DEBUG") != "True" {
		gin.SetMode(gin.ReleaseMode)

		c := config.GetConfig()
		if err := sentry.Init(sentry.ClientOptions{
			Dsn: c.GetString("sentry_dsn"),
		}); err != nil {
			fmt.Printf("Sentry initialization failed: %v\n", err)
		}

		logger, _ = zap.NewProduction()
	}
	router := gin.New()
	if os.Getenv("DEBUG") != "True" {
		router.Use(sentrygin.New(sentrygin.Options{}))
		router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
		router.Use(ginzap.RecoveryWithZap(logger, true))
	}
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(limits.RequestSizeLimiter(1000))


	return router
}
