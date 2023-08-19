package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	server := echo.New()

	json := User{
		Id:       1,
		Username: "john_doe",
		Email:    "john@example.com",
	}

	// Simple get route that responds with a json struct
	server.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, json)
	})

	// Simple post route that responds with a string
	server.POST("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "New hello on the block!")
	})

	server.Logger.Fatal(server.Start(":3000"))
}
