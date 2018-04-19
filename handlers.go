package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateLinkHandler(c echo.Context) error {
	ac := c.(AppContext)

	// TODO: validate
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
	ac := c.(AppContext)
	short := ac.Param("short")
	var stats []Stat

	ac.Db.Select("referer", "user_agent", "ip", "stats.created_at").
		From("links").
		LeftJoin("stats", "links.id = stats.link_id").
		Where("short = ?", short).
		Load(&stats)

	return ac.JSON(http.StatusOK, stats)
}

func RedirectHandler(c echo.Context) error {
	ac := c.(AppContext)
	short := ac.Param("short")
	link := &Link{}
	err := ac.Db.Select("id", "url", "short").From("links").Where("short = ?", short).LoadStruct(link)
	if err != nil {
		return ac.JSON(http.StatusNotFound, ErrorResponse{"Link not found."})
	}

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
