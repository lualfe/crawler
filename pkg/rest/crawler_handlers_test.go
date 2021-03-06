package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo"
	"github.com/lualfe/crawler/pkg/crawling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var expectedCrawlerHandlerJSONResponse = &crawling.ProfessionalPlan{
	Charge:         "R$ 5,00 por boleto pago.",
	TransferTax:    "R$ 7,00",
	MonthlyPayment: "R$ 15,00\n*pagando R$45,00/trimestre",
}

type CrawlerHandlersSuite struct {
	suite.Suite

	e *echo.Echo
	c crawling.Service
}

func (s *CrawlerHandlersSuite) SetupSuite() {
	s.e = echo.New()
	s.c = crawling.NewService()
}

func TestCrawlerHandlersInit(t *testing.T) {
	suite.Run(t, new(CrawlerHandlersSuite))
}

func (s *CrawlerHandlersSuite) TestCrawl() {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/professional_plan", nil)
	ctx := s.e.NewContext(req, rec)
	f := crawl(s.c)
	require.Error(s.T(), f(ctx), "params were not passed an an error is expected")

	p := url.Values{}
	p.Set("url", "https://www.smartmei.com.br/")
	req = httptest.NewRequest(http.MethodGet, "/api/v1/professional_plan?"+p.Encode(), nil)
	ctx = s.e.NewContext(req, rec)
	f = crawl(s.c)
	if assert.NoError(s.T(), f(ctx), "unexpected error coming from crawler") {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
		gotCrawlerHandlerResponse := &crawling.ProfessionalPlan{}
		err := json.NewDecoder(rec.Body).Decode(&gotCrawlerHandlerResponse)
		assert.Nil(s.T(), err)
		assert.Equal(s.T(), expectedCrawlerHandlerJSONResponse, gotCrawlerHandlerResponse)
	}
}
