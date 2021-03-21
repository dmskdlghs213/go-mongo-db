package model

import (
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	mongoAccessor MongoAccessor
}

func NewUserHandler(am MongoAccessor) *UserHandler {
	return &UserHandler{
		mongoAccessor: am,
	}
}

func (u *UserHandler) Create(c echo.Context) error {

	var req ReqCreate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "failed", "message": (err.Error())})
	}

	if err := u.mongoAccessor.Create(c.Request().Context(), &req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "failed", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, "ok")
}

func (u *UserHandler) Creates(c echo.Context) error {

	var req []ReqCreate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "failed", "message": (err.Error())})
	}

	if err := u.mongoAccessor.Creates(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "failed", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, "ok")
}

func (u *UserHandler) Find(c echo.Context) error {

	name := c.QueryParam("name")

	user, err := u.mongoAccessor.Find(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "failed", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) Finds(c echo.Context) error {

	var req ReqFinds
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "failed", "message": (err.Error())})
	}

	user, err := u.mongoAccessor.Finds(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "failed", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) Update(c echo.Context) error {

	userID := c.Param("user_id")

	var req ReqUpdate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "failed", "message": (err.Error())})
	}
	mongoUserID, _ := primitive.ObjectIDFromHex(userID)
	req.ID = mongoUserID

	if err := u.mongoAccessor.Update(c.Request().Context(), &req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "failed", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, "ok")
}

func (u *UserHandler) Updates(c echo.Context) error {

	var req ReqUpdates
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "failed", "message": (err.Error())})
	}

	mongoIDs := make([]primitive.ObjectID, 0, len(req.ReqID))
	for i := range req.ReqID {
		mongoUserID, _ := primitive.ObjectIDFromHex(req.ReqID[i])
		mongoIDs = append(mongoIDs, mongoUserID)
	}
	req.ID = mongoIDs

	if err := u.mongoAccessor.Updates(c.Request().Context(), &req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "failed", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, "ok")
}

func (u *UserHandler) Delete(c echo.Context) error {

	userID := c.Param("user_id")
	mongoUserID, _ := primitive.ObjectIDFromHex(userID)
	if err := u.mongoAccessor.Delete(c.Request().Context(), mongoUserID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "failed", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, "ok")
}
