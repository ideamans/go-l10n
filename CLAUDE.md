# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go internationalization (i18n) library inspired by Movable Type's localization system. It provides automatic language detection and translation support for Go applications. The library detects the user's language from environment variables and provides translation functions with test-friendly language override capabilities.

## Core Architecture

The library uses a two-level map structure with automatic language detection:
- `WorldMap`: Maps language codes (ja, en) to their respective lexicons
- `LexiconMap`: Maps phrase keys to translated strings
- Language detection from environment variables (LANGUAGE, LC_ALL, LC_MESSAGES, LANG)
- Test mode detection that forces English during test execution

Key components:
- `World` global variable: Central registry storing Japanese translations (English is fallback)
- `Language` global variable: Current active language (auto-detected via `DetectLanguage()`)
- Language override mechanism: `ForceLanguage()` and `ResetLanguage()` for testing
- Automatic initialization via `init()` function

## Common Development Commands

```bash
# Run tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Build the module
go build ./...

# Check for issues
go vet ./...

# Format code
go fmt ./...

# Run linter (if golangci-lint is available)
golangci-lint run

# Test specific functionality
go test -run TestFunctionName ./...
```

## Key Functions

- `Register(lang, lex)`: Add/merge translations for a language (typically called in init() functions)
- `T(phrase)`: Simple translation lookup with fallback to original phrase
- `DetectLanguage()`: Auto-detect language from environment variables
- `ForceLanguage(lang)`: Override language setting (for testing)
- `ResetLanguage()`: Reset to automatic language detection
- `GetCurrentLanguage()`: Get currently active language

## Usage Pattern

Each Go source file should register its translations in the `init()` function:

```go
func init() {
    l10n.Register("ja", l10n.LexiconMap{
        "Error loading file": "ファイルの読み込みエラー",
        "File not found":     "ファイルが見つかりません",
        "Success":            "成功",
    })
}
```

This allows each package/module to contribute its own translations, which are automatically merged into the global translation registry.

## Language Detection Logic

1. If language is overridden via `ForceLanguage()`, use that
2. If in test mode (detected via os.Args), default to English
3. Otherwise detect from environment variables using golang.org/x/text/language matcher
4. Supports Japanese (ja) and English (en) with English as fallback