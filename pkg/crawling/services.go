package crawling

import (
	"context"
	"net/http"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// Service is the Rest accessible interface
type Service interface {
	Crawl(*Crawler) (*ProfessionalPlan, error)
}

type service struct{}

// NewService initializes a new service
func NewService() Service {
	return &service{}
}

// Crawl crawls the crawler url and to get Professional Plan information
func (s *service) Crawl(crawler *Crawler) (*ProfessionalPlan, error) {
	if crawler == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "an empty crawler has been received as an argument")
	}
	if !strings.Contains(crawler.URL, "https://www.smartmei.com.br/") {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "crawler url is not supported")
	}
	pInfo := &ProfessionalPlan{}
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()
	log.Print("Crawler started")
	err := chromedp.Run(ctx,
		chromedp.Navigate(crawler.URL),
		chromedp.WaitVisible(`body > footer`),
		chromedp.Click(`#menu-item-88 > a`, chromedp.NodeVisible),
		chromedp.Click(`
		#planos-e-tarifas > 
			div > 
				div > 
					div > 
						ul > 
							li:nth-child(2) > 
								a`, chromedp.NodeVisible),
		chromedp.WaitVisible(`#tarifas-2`),
		chromedp.Text(`
		#tarifas-2 > 
			div:nth-child(2) > 
				div:nth-child(3)`, &pInfo.Charge),
		chromedp.Text(`
		#tarifas-2 > 
			div:nth-child(3) > 
				div:nth-child(3)`, &pInfo.TransferTax),
		chromedp.Text(`
		#tarifas-2 > 
			div:nth-child(5) > 
				div:nth-child(3)`, &pInfo.MonthlyPayment),
	)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve information from url")
	}

	return pInfo, err
}
