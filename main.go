package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"

	"github.com/WorkWorkWork-Team/common-go/databasemysql"
	"github.com/WorkWorkWork-Team/common-go/httpserver"
	"github.com/WorkWorkWork-Team/gov-voter-api/config"
	"github.com/WorkWorkWork-Team/gov-voter-api/handler"
	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
)

var appConfig config.Config

func init() {
	appConfig = config.Load()
}

func main() {
	// New .
	jwtService := service.NewJWTService(appConfig.JWT_SECRET_KEY, appConfig.JWT_ISSUER, time.Duration(appConfig.JWT_TTL)*time.Second)
	mysql, err := databasemysql.NewDbConnection(databasemysql.Config{
		Hostname:     fmt.Sprint(appConfig.MYSQL_HOSTNAME, ":", appConfig.MYSQL_PORT),
		Username:     appConfig.MYSQL_USERNAME,
		Password:     appConfig.MYSQL_PASSWORD,
		DatabaseName: appConfig.MYSQL_DATABASE,
	})
	if err != nil {
		os.Exit(1)
		return
	}

	// New Repository
	applyVoteRepository := repository.NewApplyVoteRepository(mysql)

	// New Services
	validityService := service.NewValidityService(applyVoteRepository)
	applyvoteService := service.NewApplyvoteService(applyVoteRepository)

	// New Handler
	validityHandler := handler.NewValidityHandler(jwtService, validityService)
	applyvoteHandler := handler.NewApplyVoteHandler(applyvoteService)

	// Init Gin.
	server := httpserver.NewHttpServer()
	server.GET("/validity", handler.AuthorizeJWT(jwtService, appConfig), validityHandler.Validity)
	server.POST("/applyvote", handler.AuthorizeJWT(jwtService, appConfig), applyvoteHandler.ApplyVote)

	if appConfig.Env != "prod" {
		devHandler := handler.NewDevHandler(jwtService)
		devGroup := server.Group("/dev")
		devGroup.GET("/token/:id/", devHandler.NewTestToken)
	}
	server.Run(fmt.Sprint(":", appConfig.LISTENING_PORT))
}
