package crawling

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var expectedProfessionalPlan = &ProfessionalPlan{
	Charge:         "R$ 5,00 por boleto pago.",
	TransferTax:    "R$ 7,00",
	MonthlyPayment: "R$ 15,00\n*pagando R$45,00/trimestre",
}

type CrawlingSuite struct {
	suite.Suite

	svc *service
}

func (s *CrawlingSuite) SetupSuite() {
	s.svc = &service{}
}

func TestCrawlingInit(t *testing.T) {
	suite.Run(t, new(CrawlingSuite))
}

func (s *CrawlingSuite) TestCrawl() {
	crawler := &Crawler{}
	_, err := s.svc.Crawl(crawler)
	require.Error(s.T(), err, "an empty crawler was passed and no error returned")

	crawler.URL = "https://google.com"
	_, err = s.svc.Crawl(crawler)
	require.Error(s.T(), err, "an unknown url has been passed within the crawler and an error was expected")

	crawler.URL = "https://www.smartmei.com.br/"
	gotProfessionalPlan, err := s.svc.Crawl(crawler)
	if assert.NoError(s.T(), err, "unexpected error crawling the website") {
		assert.Equal(s.T(), expectedProfessionalPlan, gotProfessionalPlan)
	}
}
