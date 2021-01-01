package command

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gocolly/colly"
	"spider.ohmgood.com/models"
	"strconv"
	"strings"
)

func storageCompany(e *colly.HTMLElement) {
	companyName := e.DOM.Find("h3[class='company-name wordCut']").Find("a").Text()
	logo, _ := e.DOM.Find("img").Attr("src")
	stageStr := e.DOM.Find("h4[class='indus-stage wordCut']").Text()
	stageArr := strings.Split(stageStr, "/")
	companyType := stageArr[0]
	financingStage := stageArr[1]
	numberOfEmployees := stageArr[2]
	o := orm.NewOrm()
	company := new(models.Company)
	company.Name = companyName
	company.CompanyType = companyType
	company.FinancingStage = financingStage
	company.NumberOfEmployees = numberOfEmployees
	company.Logo = logo
	//buttom
	commentNum := e.DOM.Find("a[class='bottom-item bottom-1 fl']").Find("p[class='green']").Text()
	jobNum := e.DOM.Find("a[class='bottom-item bottom-2 fl']").Find("p[class='green']").Text()
	jobDealRate := e.DOM.Find("a[class='bottom-item bottom-2 fl']").Find("p[class='green']").Text()
	company.CommentNum, _ = strconv.Atoi(commentNum)
	company.JobNum, _ = strconv.Atoi(jobNum)
	company.JobDealRate, _ = strconv.Atoi(jobDealRate)
	fmt.Println(o.InsertOrUpdate(company))
}
func Main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.MaxDepth(1))

	// On every a element which has href attribute call callback
	c.OnHTML("li[class='company-item']", func(e *colly.HTMLElement) {
		storageCompany(e)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.lagou.com/gongsi/")
}
