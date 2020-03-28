package main

import (
	"github.com/labstack/echo"
	"github.com/lualfe/crawler/pkg/crawling"
	"github.com/lualfe/crawler/pkg/rest"
)

func main() {
	e := echo.New()
	cs := crawling.NewService()
	rest.CrawlerHandlers(e, cs)
	e.Logger.Fatal(e.Start(":1323"))
}
