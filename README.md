
# 📝 TOEICLI Part5 CLI 🚀
- TOEIC Part5 の短文穴埋め問題を CLI上 で勉強できるツールです。
- Google Gemini API を使って問題を自動生成します。

## ✨ 特徴
- AI問題生成: Gemini API で TOEIC Part5 形式の問題を無限に自動生成 🧠

- インタラクティブ: CLI 上でサクサク回答、その場で正誤判定 ⌨️

- 圧倒的コスパ: 一日数万問分なら無料枠で利用可能 💸

## 🛠 事前準備
- Go 言語 がインストールされていること 🐹

- Gemini API キー を取得（Google Cloud AI）🔑

- 🔑 環境変数設定(ターミナルで API キーを設定してください)

### macOS / Linux 🐧🍏
.zshrc や .bashrc に追記すると便利です。

```bash
export GEMINI_API_KEY="YOUR_API_KEY"
```

### Windows (PowerShell) 🪟
```PowerShell
setx GEMINI_API_KEY "YOUR_API_KEY"
```

## 🚀 インストール & 使い方
1. インストール
```bash
go install github.com/Naughty8020/toeicli@latest
```

2. 実行
```bash
toeicli
```
### 💡 開発用（コードをビルドする場合）
```bash
git clone https://github.com/Naughty8020/toeicli.git
cd TOEICLI
go run main.go
```

## 💬 実行例
```bash
TOEIC Part5 Problem:
Sentence: "The manager ______ the report yesterday."
A) reviewed
B) review
C) reviews
D) reviewing

Your answer (A/B/C/D): A
✅ Correct! ✨
```

## 📂 ファイル構成
- main.go : CLI 本体 🚀
- go.mod : Go モジュール定義 📦
- README.md : この説明書 📝

### ⚠️ 注意
- Gemini API の無料枠には制限があります。連続リクエストにはご注意ください。
- APIキーは絶対に公開（GitHubへのコミットなど）しないでください。 🔒


