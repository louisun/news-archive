package crawl

import (
	"news-archive/pkg/entity"
	"news-archive/pkg/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetTheVergeNew() []entity.CrawlResult {
	url := "https://www.theverge.com"
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

	document.Find(".c-compact-river__entry .c-entry-box--compact--article .c-entry-box--compact__title").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		title := selection.Find("a").Text()

		if boolUrl {
			results = append(results, entity.CrawlResult{
				Title: strings.TrimSpace(title),
				URL:   url,
			})
		}
	})

	return results
}
