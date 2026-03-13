# 貢献ガイドライン (Contributing Guide)

TOEICLI への興味を持っていただきありがとうございます！バグ報告、機能改善、ドキュメントの修正など、あらゆる形での貢献を歓迎します。

## 🚀 貢献の方法

### 1. バグを見つけた場合
何か問題が発生した場合は、GitHub の [Issues](https://github.com/Naughty8020/toeicli/issues) から報告してください。その際、以下の情報を含めていただけると助かります。
- 使用している OS と Go のバージョン
- エラーの内容（ターミナルの出力など）
- 再現手順

### 2. 機能の提案・改善
「こんな機能が欲しい」というアイデアがある場合も、まずは Issue で議論を始めましょう。大きな変更を加える前に方針を相談することで、スムーズに開発を進められます。

### 3. プルリクエスト (Pull Request) を送る
以下の手順で進めてください。
1. このリポジトリを Fork する
2. 新しいブランチを作成する (`git checkout -b feature/amazing-feature`)
3. 変更をコミットする (`git commit -m 'Add some amazing feature'`)
4. ブランチを Push する (`git push origin feature/amazing-feature`)
5. Pull Request を作成する

## 🛠 コーディング規約

このプロジェクトでは、Go の標準的な規約に従います。
- **gofmt**: コードを送る前に必ず `go fmt ./...` を実行してください。
- **Linter**: 可能であれば `golangci-lint` を通してください。
- **Exported Types**: 外部パッケージから参照する構造体（`types.go` など）や関数は、ドキュメント用のコメントを添えてください。

## 📜 ライセンス
このプロジェクトに提供された貢献は、すべて本プロジェクトの [MIT License](./LICENSE) の下に置かれるものとします。

---

ご協力ありがとうございます！あなたの参加を楽しみにしています。
