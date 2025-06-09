# Go-L10n Examples

This directory contains examples demonstrating various usage patterns of the go-l10n library.

## Available Examples

### 1. Basic Example (`basic/`)
Demonstrates fundamental usage of go-l10n:
- Basic translation with `T()` function
- Formatted translation with `F()` function
- Error creation with `E()` function
- Language switching
- Environment variable detection
- Fallback behavior

### 2. Modular Example (`modular/`)
Shows how to use go-l10n in a modular application:
- Package-level translation registration
- Multiple modules with independent translations
- Automatic translation merging
- Consistent language handling across modules

## Running the Examples

Each example can be run with Go:

```bash
# Run basic example
cd basic
go run main.go

# Run with Japanese locale
LANG=ja_JP.UTF-8 go run main.go

# Run with English locale
LANG=en_US.UTF-8 go run main.go
```

## Environment Variables

The following environment variables control language detection:
- `LANGUAGE` - Highest priority
- `LC_ALL` - Second priority
- `LC_MESSAGES` - Third priority
- `LANG` - Lowest priority
- `L10N_DEFAULT_LANGUAGE` - Sets the default language (defaults to "en")
- `L10N_SKIP_DETECTION` - Skip automatic language detection if set

## Common Patterns

### 1. Registering Translations
```go
func init() {
    l10n.Register("ja", l10n.LexiconMap{
        "Hello": "こんにちは",
        "Error: %s": "エラー: %s",
    })
}
```

### 2. Basic Translation
```go
msg := l10n.T("Hello")  // Returns "こんにちは" in Japanese mode
```

### 3. Formatted Translation
```go
msg := l10n.F("Error: %s", "file not found")  // Returns "エラー: file not found" in Japanese mode
```

### 4. Error Creation
```go
err := l10n.E("Error: %s", "connection failed")  // Creates localized error
```

### 5. Language Management
```go
// Get current language
lang := l10n.GetCurrentLanguage()

// Force a specific language
l10n.ForceLanguage("ja")

// Reset to automatic detection
l10n.ResetLanguage()
```