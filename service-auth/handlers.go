package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type jwtCustomClaims struct {
	IDUser string `json:"id_user"`
	jwt.StandardClaims
}

func login(c echo.Context) error {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		IDUser      string `json:"id_user"`
		AccessToken string `json:"access_token"`
	}

	payload := new(request)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Invalid body",
		})
	}

	if payload.Email != "user1@volho.com" || payload.Password != "user1" {
		return echo.ErrUnauthorized
	}

	id := "4c7f551d-3a51-4184-82f7-30f1115feff7"
	claims := &jwtCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	res := response{
		AccessToken: t,
		IDUser:      id,
	}

	return c.JSON(http.StatusOK, res)
}
