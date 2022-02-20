package crawl

import (
	"news-archive/pkg/entity"
	"news-archive/pkg/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetEngadgetNew() []entity.CrawlResult {
	url := "https://www.engadget.com"
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

	document.Find("div[id='module-latest'] article[data-component='PostCard']").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("div[data-component='PostInfo'] a").Attr("href")
		title := selection.Find("div[data-component='PostInfo'] a").Text()

		if boolUrl {
			results = append(results, entity.CrawlResult{
				Title: strings.TrimSpace(title),
				URL:   "https://www.engadget.com" + url,
			})
		}
	})

	return results
}
