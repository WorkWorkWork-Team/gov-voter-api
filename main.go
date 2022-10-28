package main

import (
	"fmt"
	"os"
	"time"

	"github.com/WorkWorkWork-Team/common-go/databasemysql"
	"github.com/WorkWorkWork-Team/common-go/jwtservice"
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
	jwtService := jwtservice.NewJWTService(appConfig.JWT_SECRET_KEY, appConfig.JWT_ISSUER, appConfig.JWT_TTL*time.Second)
	mysql, err := databasemysql.NewDbConnection(databasemysql.Config{})
	if err != nil {
		os.Exit(1)
		return
	}

	// New Repository
	applyVoteRepository := repository.NewApplyVoteRepository(mysql)

	// New Services
	validityService := service.NewValidityService(applyVoteRepository)

	// New Handler
	validityHandler := handler.NewValidityHandler(jwtService, validityService)

	fmt.Println(validityHandler)
}
