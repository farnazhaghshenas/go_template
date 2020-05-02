package cli

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go_template/config"
	"go_template/drivers/sql"
	"go_template/routes"
)

func Serve() {
	e := echo.New()
	db, err := sql.Connect()
	if err != nil {
		logrus.Fatal("Database Connection Error")
	}
	routes.RegisterRoutes(e, db)
	logrus.Fatal(e.Start(fmt.Sprintf(":%s", config.App.Port)))
}
