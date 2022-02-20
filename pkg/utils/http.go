package utils

import (
	"net/http"
	"news-archive/pkg/common"
	"strings"
)

func DoGet(url string, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Log().WithError(err).Errorf("DoGet http.NewRequest failed")
		return nil, err
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}

		if _, exist := header["User-Agent"]; !exist {
			req.Header.Set("User-Agent", common.DefaultUserAgent)
		}
	} else {
		req.Header.Add("User-Agent", common.DefaultUserAgent)
	}

	return common.HttpClient.Do(req)
}

func DoPost(url string, body string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		Log().WithError(err).Errorf("DoPost http.NewRequest failed")
		return nil, err
	}

	req.Header.Add("User-Agent", common.DefaultUserAgent)

	return common.HttpClient.Do(req)
}