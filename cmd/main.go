package main

import (
	"fmt"

	"github.com/jessemillman/vitae/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	homeHandler := handler.HomeHandler{}
	app.GET("/", homeHandler.HandleHomeShow)

	app.Start(":3000") // switch to env

	fmt.Println("it is working")
}
