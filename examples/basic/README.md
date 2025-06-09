# Basic Go-L10n Example

This example demonstrates the basic usage of the go-l10n library.

## Features Demonstrated

- Basic translation with `T()` function
- Formatted translation with `F()` function
- Error creation with localized messages using `E()` function
- Language switching between Japanese and English
- Automatic language detection from environment variables
- Fallback behavior for untranslated phrases

## Running the Example

### Default (Auto-detect language)
```bash
go run main.go
```

### Force Japanese
```bash
LANG=ja_JP.UTF-8 go run main.go
```

### Force English
```bash
LANG=en_US.UTF-8 go run main.go
```

### Using other environment variables
```bash
# Using LC_ALL
LC_ALL=ja_JP.UTF-8 go run main.go

# Using LC_MESSAGES
LC_MESSAGES=ja go run main.go

# Using LANGUAGE
LANGUAGE=ja go run main.go
```

## Output Examples

### English Output
```
=== Go-L10n Basic Example ===

Current language: en

Basic translations:
  T("Hello, World!"): Hello, World!
  T("Welcome to Go-L10n"): Welcome to Go-L10n

Formatted translations:
  F("Current language is %s", lang): Current language is en
  F("Processing %d items", 42): Processing 42 items

...
```

### Japanese Output (with LANG=ja_JP.UTF-8)
```
=== Go-L10n Basic Example ===

Current language: ja

Basic translations:
  T("Hello, World!"): こんにちは、世界！
  T("Welcome to Go-L10n"): Go-L10nへようこそ

Formatted translations:
  F("Current language is %s", lang): 現在の言語はjaです
  F("Processing %d items", 42): 42個のアイテムを処理しています

...
```