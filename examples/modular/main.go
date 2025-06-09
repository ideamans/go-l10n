package main

import (
	"fmt"
	"strings"

	"github.com/ideamans/go-l10n"
	"github.com/ideamans/go-l10n/examples/modular/database"
	"github.com/ideamans/go-l10n/examples/modular/web"
)

// Main package translations
func init() {
	l10n.Register("ja", l10n.LexiconMap{
		"Application Starting":     "アプリケーションを開始します",
		"Loading modules...":       "モジュールを読み込んでいます...",
		"All systems operational":  "すべてのシステムが正常に動作しています",
		"Shutting down gracefully": "正常にシャットダウンしています",
		"Language: %s":             "言語: %s",
	})
}

func main() {
	fmt.Println("=== Modular Go-L10n Example ===")
	fmt.Println()

	// Show language switching with modular translations
	demonstrateInLanguage("en")
	fmt.Println()
	demonstrateInLanguage("ja")

	// Reset language
	l10n.ResetLanguage()
}

func demonstrateInLanguage(lang string) {
	l10n.ForceLanguage(lang)

	fmt.Println(l10n.F("Language: %s", lang))
	fmt.Println(strings.Repeat("-", 40))

	// Main module messages
	fmt.Println("[Main Module]")
	fmt.Println(l10n.T("Application Starting"))
	fmt.Println(l10n.T("Loading modules..."))
	fmt.Println()

	// Database module messages
	fmt.Println("[Database Module]")
	database.Connect()
	database.ExecuteQuery("SELECT * FROM users")
	database.Close()
	fmt.Println()

	// Web module messages
	fmt.Println("[Web Module]")
	web.StartServer(8080)
	web.HandleRequest("/api/users")
	web.StopServer()
	fmt.Println()

	// Main module completion
	fmt.Println("[Main Module]")
	fmt.Println(l10n.T("All systems operational"))
}
