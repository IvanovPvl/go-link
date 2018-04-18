package main

import (
	"os"

	"github.com/gocraft/dbr"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type AppContext struct {
	echo.Context
	Db *dbr.Session
}

func main() {
	godotenv.Load()
	port := os.Getenv("APP_PORT")

	conn, err := dbr.Open("postgres", os.Getenv("DATABASE_URL"), nil)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	if os.Getenv("APP_DEBUG") == "true" {
		e.Debug = true
	}

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := AppContext{c, conn.NewSession(nil)}
			return h(ac)
		}
	})

	g := e.Group("/api")
	g.POST("/links", CreateLinkHandler)
	g.GET("/stats/:short", GetStatsHandler)
	e.GET("/:short", RedirectHandler)

	e.Logger.Fatal(e.Start(":" + port))
}
