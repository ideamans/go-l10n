package web

import (
	"fmt"

	"github.com/ideamans/go-l10n"
)

// Web module registers its own translations
func init() {
	l10n.Register("ja", l10n.LexiconMap{
		// Server messages
		"Starting web server on port %d": "ポート%dでWebサーバーを起動しています",
		"Server started successfully":    "サーバーが正常に起動しました",
		"Server startup failed: %s":      "サーバーの起動に失敗しました: %s",

		// Request handling
		"Handling request: %s":     "リクエストを処理しています: %s",
		"Request completed: %d ms": "リクエスト完了: %d ms",
		"Request failed: %s":       "リクエスト失敗: %s",

		// Response messages
		"Sending response: %d":       "レスポンスを送信: %d",
		"Response sent successfully": "レスポンスを正常に送信しました",

		// Server shutdown
		"Stopping web server": "Webサーバーを停止しています",
		"Server stopped":      "サーバーが停止しました",

		// Common HTTP status messages
		"Not Found":             "見つかりません",
		"Internal Server Error": "内部サーバーエラー",
		"Bad Request":           "不正なリクエスト",
		"Unauthorized":          "認証が必要です",
	})
}

// StartServer simulates starting a web server
func StartServer(port int) {
	fmt.Println(l10n.F("Starting web server on port %d", port))
	// Simulate server startup
	fmt.Println(l10n.T("Server started successfully"))
}

// HandleRequest simulates handling an HTTP request
func HandleRequest(path string) {
	fmt.Println(l10n.F("Handling request: %s", path))
	// Simulate request processing
	fmt.Println(l10n.F("Request completed: %d ms", 125))
	fmt.Println(l10n.F("Sending response: %d", 200))
}

// StopServer simulates stopping the web server
func StopServer() {
	fmt.Println(l10n.T("Stopping web server"))
	// Simulate server shutdown
	fmt.Println(l10n.T("Server stopped"))
}
