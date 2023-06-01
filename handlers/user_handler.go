package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/panurujz/resume-service/config"
	"github.com/panurujz/resume-service/models"
	"github.com/panurujz/resume-service/requests"
	"github.com/panurujz/resume-service/utils"
)

var db = config.Open()

func CreateUser(c echo.Context) (err error) {

	r := requests.UserReq{}
	if err = c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	u := models.User{
		Name:     r.Name,
		Surname:  r.Surname,
		Nickname: r.Nickname,
	}

	newUser := saveUser(u)

	return c.JSON(http.StatusCreated, newUser)
}

func saveUser(user models.User) models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tx := db.WithContext(ctx)
	tx.Create(&user)

	res, err := utils.PrettyStruct(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("create user success. %s \n", res)

	return user
}
