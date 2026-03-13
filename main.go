package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
