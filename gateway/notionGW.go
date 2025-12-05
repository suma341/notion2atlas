package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"notion2atlas/domain"
	"os"
)

func GetNotionData(resourceType domain.NotionResourceType, id string) (map[string]any, error) {
	secret := os.Getenv("NOTION_TOKEN_HORIZON")

	var url string
	var method string
	var body map[string]any = nil
	url, method, body = resourceType.GetRequestQuery(id)

	var buf *bytes.Buffer
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			fmt.Println("error in gateway/GetNotionData/json.Marshal")
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	} else {
		buf = bytes.NewBuffer([]byte("{}"))
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		fmt.Println("error in gateway/GetNotionData/http.NewRequest")
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+secret)
	req.Header.Set("Notion-Version", "2022-06-28")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in gateway/GetNotionData/client.Do")
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]any
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("error in gateway/GetNotionData/json.NewDecoder(resp.Body).Decode(&data)")
		return nil, err
	}

	return data, nil
}
