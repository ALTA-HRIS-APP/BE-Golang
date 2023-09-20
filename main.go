package main

import (
	"be_golang/klp3/app/config"
	"be_golang/klp3/app/database"
	"be_golang/klp3/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// api.ApiGetUser()
	cfg := config.InitConfig()
	mysql := database.InitMysql(cfg)
	database.InitialMigration(mysql)
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	
	router.InitRouter(e, mysql)
	e.Logger.Fatal(e.Start(":80"))
}
