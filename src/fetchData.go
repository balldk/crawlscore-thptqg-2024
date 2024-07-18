package src

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchRaw(sbd string) string {
	url := fmt.Sprintf("https://vietnamnet.vn/giao-duc/diem-thi/tra-cuu-diem-thi-tot-nghiep-thpt/2024/%s.html", sbd)
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(sbd, err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(sbd, err)
		return ""
	}

	return string(body)
}

func FetchScore(sbd string) *Student {
	var response Response
	htmlBody := FetchRaw(sbd)
	json.Unmarshal([]byte(htmlBody), &response)

	var std *Student

	std = ParseStudent(&htmlBody, &sbd)
	return std
}
