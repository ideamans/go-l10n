# go-l10n

[English README](README.md)

Movable Type のローカライゼーションシステムにインスパイアされた Go 用国際化（i18n）ライブラリ。自動言語検出と翻訳サポートを提供し、テスト対応の言語オーバーライド機能を備えています。

## 特徴

- **自動言語検出**: 環境変数から自動的に言語を検出
- **シンプルな翻訳**: `T()` 関数での簡単な翻訳
- **分散型登録**: 各パッケージが独自の翻訳を提供可能
- **テスト対応**: テスト時の言語強制機能
- **環境変数対応**: デフォルト言語や検出スキップを制御可能

## インストール

```bash
go get github.com/ideamans/go-l10n
```

## 基本的な使用方法

### 1. 翻訳の登録

各 Go ソースファイルの `init()` 関数で翻訳を登録します：

```go
package main

import "github.com/ideamans/go-l10n"

func init() {
    l10n.Register("ja", l10n.LexiconMap{
        "Hello, World!":     "こんにちは、世界！",
        "File not found":    "ファイルが見つかりません",
        "Error occurred":    "エラーが発生しました",
        "Operation success": "操作が成功しました",
    })
}
```

### 2. 翻訳の使用

```go
package main

import (
    "fmt"
    "github.com/ideamans/go-l10n"
)

func main() {
    // 環境変数に基づいて自動的に言語が検出されます
    fmt.Println(l10n.T("Hello, World!"))
    // 日本語環境では: "こんにちは、世界！"
    // 英語環境では: "Hello, World!"
}
```

## 言語検出ロジック

言語は以下の優先順位で決定されます：

1. **強制設定**: `ForceLanguage()` で設定された言語
2. **テストモード**: テスト実行中はデフォルト言語を使用
3. **環境変数**: 以下の環境変数から検出
   - `LANGUAGE`
   - `LC_ALL`
   - `LC_MESSAGES`
   - `LANG`

現在サポートされている言語：
- `ja`: 日本語
- `en`: 英語（フォールバック）

## 環境変数

### L10N_DEFAULT_LANGUAGE

デフォルト言語を設定します。設定されていない場合は `"en"` が使用されます。

```bash
export L10N_DEFAULT_LANGUAGE=fr
```

### L10N_SKIP_DETECTION

空でない値が設定されている場合、初期化時の自動言語検出をスキップします。

```bash
export L10N_SKIP_DETECTION=1
```

## API リファレンス

### 型定義

```go
type LexiconMap map[string]string  // 基本フレーズから翻訳フレーズへのマップ
type WorldMap map[string]LexiconMap // 言語コードから辞書へのマップ
```

### 関数

#### Register(lang string, lex LexiconMap)

指定した言語の翻訳を登録します。既存の翻訳にマージされます。

```go
l10n.Register("ja", l10n.LexiconMap{
    "Save": "保存",
    "Load": "読込",
})
```

#### T(phrase string) string

フレーズを翻訳します。翻訳が見つからない場合は元のフレーズを返します。

```go
message := l10n.T("File saved successfully")
```

#### ForceLanguage(lang string)

言語を強制的に設定します（主にテスト用）。

```go
l10n.ForceLanguage("ja")
```

#### ResetLanguage()

言語検出を自動モードにリセットします。

```go
l10n.ResetLanguage()
```

#### GetCurrentLanguage() string

現在アクティブな言語を取得します。

```go
currentLang := l10n.GetCurrentLanguage()
```

#### DetectLanguage()

環境変数から言語を検出し、設定します。

```go
l10n.DetectLanguage()
```

## 使用例

### 複数パッケージでの翻訳登録

```go
// package database
package database

import "github.com/ideamans/go-l10n"

func init() {
    l10n.Register("ja", l10n.LexiconMap{
        "Connection failed":    "接続に失敗しました",
        "Query executed":       "クエリを実行しました",
        "Transaction started":  "トランザクションを開始しました",
    })
}

// package auth
package auth

import "github.com/ideamans/go-l10n"

func init() {
    l10n.Register("ja", l10n.LexiconMap{
        "Login successful":     "ログインに成功しました",
        "Invalid credentials":  "認証情報が無効です",
        "Session expired":      "セッションが期限切れです",
    })
}
```

### テストでの言語制御

```go
func TestJapaneseMessages(t *testing.T) {
    // テスト用に日本語を強制設定
    l10n.ForceLanguage("ja")
    defer l10n.ResetLanguage() // テスト後にリセット

    message := l10n.T("Hello, World!")
    if message != "こんにちは、世界！" {
        t.Errorf("Expected Japanese translation, got: %s", message)
    }
}
```

## テスト

```bash
# すべてのテストを実行
go test ./...

# 詳細な出力でテストを実行
go test -v ./...

# 特定のテストを実行
go test -run TestT_BasicTranslation ./...
```

## ライセンス

このプロジェクトは Movable Type のローカライゼーションシステムにインスパイアされています。

## 貢献

バグ報告や機能リクエストは Issue でお知らせください。プルリクエストも歓迎します。

## ライセンス

このプロジェクトはMITライセンスの下でライセンスされています。詳細は[LICENSE](LICENSE)ファイルをご覧ください。

---

[English README](README.md)