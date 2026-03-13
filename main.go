package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Part struct {
	Text string `json:"text"`
}

type Content struct {
	Parts []Part `json:"parts"`
	Role  string `json:"role"`
}

type Candidate struct {
	Content      Content `json:"content"`
	FinishReason string  `json:"finishReason"`
	Index        int     `json:"index"`
}

type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
}

func main() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Println("GEMINI_API_KEY not set")
		return
	}

	// 問題生成プロンプト
	prompt := `
Create 1 TOEIC Part5 problem (short sentence with one blank).
Provide the sentence and 4 options (A, B, C, D), 
then indicate the correct answer in format:
Sentence: "______ is very important."
A) It
B) They
C) This
D) These
Answer: A
`

	// 問題生成
	problemText, correctAnswer := generateProblem(apiKey, prompt)
	if problemText == "" {
		fmt.Println("Failed to generate problem")
		return
	}

	// CLIに問題だけ表示（答えは消す）
	lines := strings.Split(problemText, "\n")
	fmt.Println("TOEIC Part5 Problem:")
	for _, line := range lines {
		if strings.HasPrefix(line, "Answer:") {
			continue
		}
		fmt.Println(line)
	}

	// ユーザー入力
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nYour answer (A/B/C/D): ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(strings.ToUpper(userInput))

	// 判定
	if userInput == correctAnswer {
		fmt.Println("✅ Correct!")
	} else {
		fmt.Println("❌ Incorrect.")
		fmt.Println("Correct answer:", correctAnswer)
	}
}

// 問題生成＋正解抽出
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

	var gemResp GeminiResponse
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
