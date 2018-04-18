package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Link struct {
	Id    int64  `json:"id,omitempty";db:"id"`
	Url   string `json:"url";db:"url"`
	Short string `json:"short";db:"short"`
}

func CreateLinkHandler(c echo.Context) error {
	ac := c.(AppContext)

	link := new(Link)
	if err := c.Bind(link); err != nil {
		panic(err)
	}

	var short string
	var id int
	for {
		short = genShort()
		ac.Db.Select("id").From("links").Where("short = ?", short).LoadValue(&id)
		if id == 0 {
			break
		}
	}

	link.Short = short
	_, err := ac.Db.InsertInto("links").Columns("url", "short").Record(link).Exec()
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, link)
}

func GetStatsHandler(c echo.Context) error {
	return nil
}

func RedirectHandler(c echo.Context) error {
	ac := c.(AppContext)
	short := ac.Param("short")
	var url string
	ac.Db.Select("url").From("links").Where("short = ?", short).LoadValue(&url)
	return ac.Redirect(http.StatusMovedPermanently, url)
}
