package entity

type CrawlResult struct {
	Heat       string `json:"heat"`       // 热度信息
	Image      string `json:"image"`      // 图片地址
	Title      string `json:"title"`      // 标题
	SubTitle   string `json:"subTitle"`   // 副标题
	URL        string `json:"url"`        // 链接
	CommentURL string `json:"commentUrl"` // 评论地址
}
