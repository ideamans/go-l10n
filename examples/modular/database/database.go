package database

import (
	"fmt"

	"github.com/ideamans/go-l10n"
)

// Database module registers its own translations
func init() {
	l10n.Register("ja", l10n.LexiconMap{
		// Connection messages
		"Connecting to database...": "データベースに接続しています...",
		"Connected successfully":    "正常に接続しました",
		"Connection failed: %s":     "接続に失敗しました: %s",

		// Query messages
		"Executing query: %s":     "クエリを実行しています: %s",
		"Query completed in %dms": "クエリが%dmsで完了しました",
		"Query error: %s":         "クエリエラー: %s",

		// Transaction messages
		"Beginning transaction":    "トランザクションを開始します",
		"Committing transaction":   "トランザクションをコミットしています",
		"Rolling back transaction": "トランザクションをロールバックしています",

		// Connection pool messages
		"Closing database connection": "データベース接続を閉じています",
		"Connection closed":           "接続を閉じました",
	})
}

// Connect simulates database connection
func Connect() {
	fmt.Println(l10n.T("Connecting to database..."))
	// Simulate connection
	fmt.Println(l10n.T("Connected successfully"))
}

// ExecuteQuery simulates query execution
func ExecuteQuery(query string) {
	fmt.Println(l10n.F("Executing query: %s", query))
	// Simulate query execution
	fmt.Println(l10n.F("Query completed in %dms", 42))
}

// Close simulates closing the database connection
func Close() {
	fmt.Println(l10n.T("Closing database connection"))
	// Simulate closing
	fmt.Println(l10n.T("Connection closed"))
}
