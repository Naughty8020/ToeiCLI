package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Naughty8020/toeicli/config"
)

func generateProblem(apiKey, prompt string) (string, string) {
	reqBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{"parts": []map[string]string{{"text": prompt}}},
		},
	}

	bodyBytes, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(
		"POST",
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent",
		bytes.NewBuffer(bodyBytes),
	)
	req.Header.Set("x-goog-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return "", ""
	}
	defer resp.Body.Close()

	respBytes, _ := ioutil.ReadAll(resp.Body)

	var gemResp config.GeminiResponse
	if err := json.Unmarshal(respBytes, &gemResp); err != nil {
		fmt.Println("Unmarshal error:", err)
		fmt.Println("Raw:", string(respBytes))
		return "", ""
	}

	if len(gemResp.Candidates) == 0 || len(gemResp.Candidates[0].Content.Parts) == 0 {
		return "", ""
	}

	// 出力テキストを取り出す
	text := gemResp.Candidates[0].Content.Parts[0].Text

	// 正解を抽出
	correctAnswer := ""
	for _, line := range strings.Split(text, "\n") {
		if strings.HasPrefix(line, "Answer:") {
			correctAnswer = strings.TrimSpace(strings.TrimPrefix(line, "Answer:"))
		}
	}

	return text, correctAnswer
}
