package main

import (
	"github.com/dmskdlghs213/go-mongoDB/app/model"
	"github.com/dmskdlghs213/go-mongoDB/app/mongodb"
	"github.com/labstack/echo"
)

func main() {
	e := Router()
	e.Logger.Fatal(e.Start(":8888")) // listen and serve on :8888
}

func Router() *echo.Echo {
	e := echo.New()
	di := Di()
	e.POST("/users", di.Create)
	e.POST("/user-groups", di.Creates)
	e.GET("/users", di.Find)
	e.GET("/user-groups", di.Finds)
	e.PATCH("/users/:user_id", di.Update)
	e.PATCH("/users", di.Updates)
	e.DELETE("/users/:user_id", di.Delete)

	return e
}

func Di() *model.UserHandler {
	m := mongodb.NewMongoQuery(mongodb.Mongo)
	u := model.NewUserHandler(m)
	return u
}
