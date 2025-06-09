# go-l10n

[日本語版 README はこちら / Japanese README](README.ja.md)

A Go internationalization (i18n) library inspired by Movable Type's localization system. It provides automatic language detection and translation support for Go applications with test-friendly language override capabilities.

## Features

- **Automatic Language Detection**: Automatically detects language from environment variables
- **Simple Translation**: Easy translation with the `T()` function
- **Distributed Registration**: Each package can provide its own translations
- **Test-Friendly**: Language forcing functionality for testing
- **Environment Variable Support**: Control default language and detection behavior

## Installation

```bash
go get github.com/ideamans/go-l10n
```

## Basic Usage

### 1. Register Translations

Register translations in the `init()` function of each Go source file:

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

### 2. Use Translations

```go
package main

import (
    "fmt"
    "github.com/ideamans/go-l10n"
)

func main() {
    // Language is automatically detected based on environment variables
    fmt.Println(l10n.T("Hello, World!"))
    // In Japanese environment: "こんにちは、世界！"
    // In English environment: "Hello, World!"
}
```

## Language Detection Logic

Language is determined in the following priority order:

1. **Forced Setting**: Language set by `ForceLanguage()`
2. **Test Mode**: Uses default language during test execution
3. **Environment Variables**: Detected from the following environment variables:
   - `LANGUAGE`
   - `LC_ALL`
   - `LC_MESSAGES`
   - `LANG`

Currently supported languages:
- `ja`: Japanese
- `en`: English (fallback)

## Environment Variables

### L10N_DEFAULT_LANGUAGE

Sets the default language. If not set, `"en"` is used.

```bash
export L10N_DEFAULT_LANGUAGE=fr
```

### L10N_SKIP_DETECTION

If set to a non-empty value, skips automatic language detection during initialization.

```bash
export L10N_SKIP_DETECTION=1
```

## API Reference

### Type Definitions

```go
type LexiconMap map[string]string  // Maps base phrases to translated phrases
type WorldMap map[string]LexiconMap // Maps language codes to lexicons
```

### Functions

#### Register(lang string, lex LexiconMap)

Registers translations for the specified language. Merges with existing translations.

```go
l10n.Register("ja", l10n.LexiconMap{
    "Save": "保存",
    "Load": "読込",
})
```

#### T(phrase string) string

Translates a phrase. Returns the original phrase if no translation is found.

```go
message := l10n.T("File saved successfully")
```

#### F(phrase string, args ...interface{}) string

Translates a phrase and formats it with fmt.Sprintf.

```go
message := l10n.F("Found %d files", count)
```

#### E(phrase string, args ...interface{}) error

Translates a phrase and returns it as an error with fmt.Errorf.

```go
err := l10n.E("Failed to open %s", filename)
```

#### ForceLanguage(lang string)

Forces the language to a specific value (primarily for testing).

```go
l10n.ForceLanguage("ja")
```

#### ResetLanguage()

Resets language detection to automatic mode.

```go
l10n.ResetLanguage()
```

#### GetCurrentLanguage() string

Returns the currently active language code.

```go
currentLang := l10n.GetCurrentLanguage()
```

#### DetectLanguage()

Detects and sets the language from environment variables.

```go
l10n.DetectLanguage()
```

## Examples

The `examples/` directory contains practical examples demonstrating various usage patterns:

- **[Basic Example](examples/basic/)**: Demonstrates fundamental usage including translations, language switching, and environment variables
- **[Modular Example](examples/modular/)**: Shows how to organize translations in a multi-package application

To run the examples:

```bash
# Run basic example
cd examples/basic
go run main.go

# Run with Japanese locale
LANG=ja_JP.UTF-8 go run main.go

# Run modular example
cd examples/modular
go run .
```

## Usage Examples

### Translation Registration Across Multiple Packages

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

### Language Control in Tests

```go
func TestJapaneseMessages(t *testing.T) {
    // Force Japanese for testing
    l10n.ForceLanguage("ja")
    defer l10n.ResetLanguage() // Reset after test

    message := l10n.T("Hello, World!")
    if message != "こんにちは、世界！" {
        t.Errorf("Expected Japanese translation, got: %s", message)
    }
}
```

### Formatted Translations

```go
// Using F() for formatted messages
count := 5
message := l10n.F("Found %d files", count)
// English: "Found 5 files"
// Japanese: "5個のファイルが見つかりました" (if registered)

// Using E() for formatted errors
filename := "config.json"
err := l10n.E("Failed to read %s", filename)
// Returns error with translated message
```

## Testing

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific test
go test -run TestT_BasicTranslation ./...
```

## Inspiration

This project is inspired by Movable Type's localization system, providing a simple yet powerful approach to internationalization in Go applications.

## Contributing

Bug reports and feature requests are welcome via Issues. Pull requests are also appreciated.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

[日本語版 README はこちら / Japanese README](README.ja.md)