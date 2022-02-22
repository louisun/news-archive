package crawl

import (
	"news-archive/pkg/entity"
	"news-archive/pkg/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetHackerNews() []entity.CrawlResult {
	url := "https://news.ycombinator.com/news"
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
	document.Find(".itemlist tbody tr").Each(func(i int, selection *goquery.Selection) {
		if selection.HasClass("athing") {
			url, exists := selection.Find(".title a").Attr("href")
			title := selection.Find(".titlelink").Text()
			result := entity.CrawlResult{
				Title: title,
				URL:   url,
			}
			if exists {
				results = append(results, result)
			}
			return
		}

		if td := selection.Find(".subtext"); len(results) > 0 && td.Length() > 0 {
			td.Find("a").Each(func(i int, selection *goquery.Selection) {
				if strings.Contains(selection.Text(), "comments") {
					results[len(results)-1].Heat = getHackerNewsHeat(selection.Text())
					commentLink, exists := selection.Attr("href")
					if exists {
						results[len(results)-1].CommentURL = "https://news.ycombinator.com/" + commentLink
					}
				}
			})
		}
	})

	return results

}

func getHackerNewsHeat(heat string) string {
	return heat
}
