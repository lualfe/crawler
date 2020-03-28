package rest

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lualfe/crawler/pkg/crawling"
)

// CrawlerHandlers represents crawler routes
func CrawlerHandlers(e *echo.Echo, cs crawling.Service) {
	e.GET("/api/v1/crawl", crawl(cs))
}

func crawl(cs crawling.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		crawler := &crawling.Crawler{}
		c.Bind(crawler)
		if crawler.URL == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "url cannot be an empty parameter")
		}
		info, err := cs.Crawl(crawler)
		if err != nil {
			return err
		}
		c.JSON(http.StatusOK, info)
		return nil
	}
}
