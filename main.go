package main

import (
	"github.com/WorkWorkWork-Team/gov-voter-api/config"
)

var appConfig config.Config

func init() {
	appConfig = config.Load()
}

func main() {

}
