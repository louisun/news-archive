package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"news-archive/pkg/crawl"
	"news-archive/pkg/entity"
	"news-archive/pkg/utils"
	"os"
)

type HotNewsImportDto struct {
	Data      []entity.CrawlResult `json:"email" binding:"required"`
	Website   string               `json:"website" binding:"required"`
	SecretKey string               `json:"secretKey" binding:"required"`
}

func main() {
	if len(os.Getenv("API_URL")) == 0 || len(os.Getenv("SECRET_KEY")) == 0 {
		panic("API_URL and SECRET_KEY must be set")
	}

	apiUrl := os.Getenv("API_URL")
	secretKey := os.Getenv("SECRET_KEY")

	for website, handler := range crawl.WebsiteHandlerMap {
		result := handler()
		if result != nil && len(result) > 0 {
			jsonBytes, err := json.Marshal(result)
			if err != nil {
				utils.Log().WithError(err).Errorf("failed to marshal json, website=%s", website)
				continue
			}

			dto := HotNewsImportDto{
				Data:      result,
				Website:   website,
				SecretKey: secretKey,
			}

			bodyBytes, err := json.Marshal(dto)
			if err != nil {
				utils.Log().WithError(err).Errorf("failed to marshal http body json, website=%s", website)
				continue
			}

			// export data to server
			exportToServer(apiUrl, website, string(bodyBytes))

			// 文件写入
			filePath := fmt.Sprintf("data/%s.json", website)

			err = ioutil.WriteFile(filePath, jsonBytes, 0644)
			if err != nil {
				utils.Log().WithError(err).Error("failed to write file, website=%s", website)
				continue
			}

			utils.Log().Infof("successfully update file, website=%s", website)
		}
	}
}

func exportToServer(apiUrl string, website string, body string) {
	resp, err := utils.DoPost(apiUrl, body)
	if err != nil {
		utils.Log().WithError(err).Errorf("failed to post data, website=%s", website)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		utils.Log().Errorf("failed to post data, website=%s, statusCode=%d", website, resp.StatusCode)
		return
	}


	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.Log().WithError(err).Errorf("failed to read response body, website=%s", website)
		return
	}

	utils.Log().Infof("successfully post data, website=%s, resp=%s", website, string(bytes))
}
