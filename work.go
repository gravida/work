package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gravida/work/models"
	"github.com/gravida/work/pkg/settings"
	"github.com/gravida/work/routers"
	_ "github.com/mattn/go-sqlite3"
)

// GOOS=linux GOARCH=amd64 go build -o coolgo_linux *.go
// GOOS=windows GOARCH=amd64 go build -o coolgo_win *.go
func main() {

	settings.Setup()
	models.Setup()

	g := routers.InitRouter()
	g.Run(settings.AppCfg.Addr)
}
