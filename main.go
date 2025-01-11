package main

import (
	"log"
	"net/http"
	"os"

	"gc-phase3-2/config"
	"gc-phase3-2/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
        e := echo.New()

        e.Use(middleware.Logger())
	    e.Use(middleware.Recover())

        err := godotenv.Load()
        if err != nil {
            log.Fatalf("Error loading .env file")
        }

		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "test heroku livecode 3")
		})
		e.POST("/users/login", service.LoginUser)
        e.POST("/users/register", service.RegisterUser)


        config.InitDB()
	    defer config.CloseDB()

	    config.ClearPreparedStatements()
    
        port := os.Getenv("PORT")
        if port == "" {
            port = "8080" 
        }

        e.Logger.Fatal(e.Start(":" + port))

    }