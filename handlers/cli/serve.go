package cli

import (
	"black_gate/config"
	"black_gate/drivers/sql"
	"black_gate/routes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Serve(){
	e := echo.New()
	db, err := sql.Connect()
	if err != nil{
	logrus.Fatal("Database Connection Error")
	}
	routes.RegisterRoutes(e, db)
	logrus.Fatal(e.Start(fmt.Sprintf(":%s", config.App.Port)))
}
