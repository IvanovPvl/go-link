package main

import (
	"net/http"

	"github.com/labstack/echo"
)

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
	link := &Link{}
	err := ac.Db.Select("id", "url", "short").From("links").Where("short = ?", short).LoadStruct(link)
	if err != nil {
		panic(err)
	}

	// TODO: check link

	stat := &Stat{
		LinkId:    link.Id,
		Ip:        ac.RealIP(),
		Referer:   ac.Request().Referer(),
		UserAgent: ac.Request().UserAgent(),
	}
	_, err = ac.Db.InsertInto("stats").Columns("referer", "user_agent", "ip", "link_id").Record(stat).Exec()
	if err != nil {
		panic(err)
	}

	return ac.Redirect(http.StatusMovedPermanently, link.Url)
}
