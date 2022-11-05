package main

import (
	"fmt"
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

	populationRepository := repository.NewPopulationRepository(mysql)
	applyVoteRepository := repository.NewApplyVoteRepository(mysql)

	// New Services

	jwtService := service.NewJWTService(appConfig.JWT_SECRET_KEY, appConfig.JWT_ISSUER, time.Duration(appConfig.JWT_TTL)*time.Second)
	voteService := service.NewVoteService(applyVoteRepository)
	authenticationService := service.NewAuthenticationService(jwtService, populationRepository)
	populationService := service.NewPopulationService(populationRepository)

	// New Handler
	userHandler := handler.NewUserHandler(populationService, jwtService, voteService)
	authenticationHandler := handler.NewAuthenticateHandler(authenticationService)

	server := httpserver.NewHttpServer()
	server.GET("/user/info", handler.AuthorizeJWT(jwtService, appConfig), userHandler.GetPopulationInfo)
	server.GET("/validity", handler.AuthorizeJWT(jwtService, appConfig), userHandler.Validity)
	server.POST("/applyvote", handler.AuthorizeJWT(jwtService, appConfig), userHandler.ApplyVote)
	server.POST("/auth/login/", authenticationHandler.AuthAndGenerateToken)

	if appConfig.Env != "prod" {
		devHandler := handler.NewDevHandler(jwtService)
		devGroup := server.Group("/dev")
		devGroup.GET("/token/:id/", devHandler.NewTestToken)
	}
	server.Run(fmt.Sprint(":", appConfig.LISTENING_PORT))
}
