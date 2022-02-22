package crawl

import "news-archive/pkg/entity"

type CrawlFunc func() []entity.CrawlResult

var WebsiteHandlerMap = map[string]CrawlFunc{
	"v2ex":        GetV2ex,
	"hacker_news": GetHackerNews,
	"theverge_new": GetTheVergeNew,
	"engadget_new": GetEngadgetNew,
}

var HotRankCache = make(map[string][]entity.CrawlResult, len(WebsiteHandlerMap))
