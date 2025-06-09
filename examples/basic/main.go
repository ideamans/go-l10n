package main

import (
	"fmt"
	"os"

	"github.com/ideamans/go-l10n"
)

// Register Japanese translations during initialization
func init() {
	l10n.Register("ja", l10n.LexiconMap{
		// Basic messages
		"Hello, World!":          "こんにちは、世界！",
		"Welcome to Go-L10n":     "Go-L10nへようこそ",
		"Current language is %s": "現在の言語は%sです",

		// Application messages
		"Starting application...": "アプリケーションを起動しています...",
		"Processing %d items":     "%d個のアイテムを処理しています",
		"Operation completed":     "操作が完了しました",

		// Error messages
		"Error: %s":          "エラー: %s",
		"File not found: %s": "ファイルが見つかりません: %s",
		"Invalid input":      "無効な入力です",
		"Connection failed":  "接続に失敗しました",
	})
}

func main() {
	fmt.Println("=== Go-L10n Basic Example ===")
	fmt.Println()

	// Show current language
	currentLang := l10n.GetCurrentLanguage()
	fmt.Printf("Current language: %s\n", currentLang)
	fmt.Println()

	// Basic translation with T()
	fmt.Println("Basic translations:")
	fmt.Printf("  T(\"Hello, World!\"): %s\n", l10n.T("Hello, World!"))
	fmt.Printf("  T(\"Welcome to Go-L10n\"): %s\n", l10n.T("Welcome to Go-L10n"))
	fmt.Println()

	// Formatted translation with F()
	fmt.Println("Formatted translations:")
	fmt.Printf("  F(\"Current language is %%s\", lang): %s\n", l10n.F("Current language is %s", currentLang))
	fmt.Printf("  F(\"Processing %%d items\", 42): %s\n", l10n.F("Processing %d items", 42))
	fmt.Println()

	// Error creation with E()
	fmt.Println("Error translations:")
	err1 := l10n.E("File not found: %s", "config.json")
	fmt.Printf("  E(\"File not found: %%s\", \"config.json\"): %v\n", err1)
	err2 := l10n.E("Invalid input")
	fmt.Printf("  E(\"Invalid input\"): %v\n", err2)
	fmt.Println()

	// Language switching demonstration
	fmt.Println("Language switching demonstration:")
	fmt.Println()

	// Switch to Japanese
	fmt.Println("Switching to Japanese (ja)...")
	l10n.ForceLanguage("ja")
	demonstrateTranslations()

	fmt.Println()

	// Switch to English
	fmt.Println("Switching to English (en)...")
	l10n.ForceLanguage("en")
	demonstrateTranslations()

	fmt.Println()

	// Reset to automatic detection
	fmt.Println("Resetting to automatic language detection...")
	l10n.ResetLanguage()
	fmt.Printf("Language after reset: %s\n", l10n.GetCurrentLanguage())
	fmt.Println()

	// Environment variable demonstration
	fmt.Println("Environment variable demonstration:")
	if lang := os.Getenv("LANG"); lang != "" {
		fmt.Printf("  LANG=%s\n", lang)
	}
	if lcAll := os.Getenv("LC_ALL"); lcAll != "" {
		fmt.Printf("  LC_ALL=%s\n", lcAll)
	}
	if lcMessages := os.Getenv("LC_MESSAGES"); lcMessages != "" {
		fmt.Printf("  LC_MESSAGES=%s\n", lcMessages)
	}
	if language := os.Getenv("LANGUAGE"); language != "" {
		fmt.Printf("  LANGUAGE=%s\n", language)
	}
	fmt.Println()

	// Fallback behavior
	fmt.Println("Fallback behavior:")
	fmt.Printf("  Untranslated phrase: \"%s\"\n", l10n.T("This phrase has no translation"))
	fmt.Printf("  Unknown language (forcing 'fr'): ")
	l10n.ForceLanguage("fr")
	fmt.Printf("\"%s\"\n", l10n.T("Hello, World!"))

	// Reset language before exit
	l10n.ResetLanguage()
}

func demonstrateTranslations() {
	fmt.Printf("  - %s\n", l10n.T("Starting application..."))
	fmt.Printf("  - %s\n", l10n.F("Processing %d items", 10))
	fmt.Printf("  - %s\n", l10n.T("Operation completed"))
	fmt.Printf("  - Error: %v\n", l10n.E("Connection failed"))
}
