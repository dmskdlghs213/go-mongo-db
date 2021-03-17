package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := Router()
	e.Logger.Fatal(e.Start(":8888")) // listen and serve on :8888
}

func Router() *echo.Echo {
	e := echo.New()
	e.POST("/insert", Insert)

	return e
}

func Insert(c echo.Context) error {

	dbName := os.Getenv("MONGO_DBNAME")
	fmt.Println(dbName)

	fmt.Println("insertに成功しました")

	return c.JSON(http.StatusOK, nil)
}
