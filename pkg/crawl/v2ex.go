package crawl

import (
	"github.com/PuerkitoBio/goquery"
	"news-archive/pkg/entity"
	"news-archive/pkg/utils"
)

func GetV2ex() []entity.CrawlResult {
	url := "https://www.v2ex.com/?tab=hot"
	// part 1 : send requests
	resp, err := utils.DoGet(url, nil)
	if err != nil {
		utils.Log().WithError(err).WithField("url", url).Errorf("DoGet failed")
		return nil
	}

	defer resp.Body.Close()

	// part 2 : parse items
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		utils.Log().WithError(err).WithField("url", url).Errorf("NewDocumentFromReader failed")
		return nil
	}

	var results []entity.CrawlResult
	document.Find(".item_title").Each(func(i int, selection *goquery.Selection) {
		url, exists := selection.Find("a").Attr("href")
		text := selection.Find("a").Text()
		if exists {
			results = append(results, entity.CrawlResult{Title: text, URL: "https://www.v2ex.com" + url})
		}
	})

	return results
}
