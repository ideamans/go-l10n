# Modular Go-L10n Example

This example demonstrates how to use go-l10n in a modular application where different packages register their own translations.

## Structure

- `main.go` - Main application with its own translations
- `database/` - Database module with database-specific translations
- `web/` - Web server module with web-specific translations

## Key Features

1. **Modular Translation Registration**: Each package registers its own translations in its `init()` function
2. **Automatic Merging**: All translations are automatically merged into the global translation registry
3. **Package Independence**: Each module manages its own translations without knowing about others
4. **Language Consistency**: All modules share the same language setting

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

## Sample Output

### English Output
```
=== Modular Go-L10n Example ===

Language: en
----------------------------------------
[Main Module]
Application Starting
Loading modules...

[Database Module]
Connecting to database...
Connected successfully
Executing query: SELECT * FROM users
Query completed in 42ms
Closing database connection
Connection closed

[Web Module]
Starting web server on port 8080
Server started successfully
Handling request: /api/users
Request completed: 125 ms
Sending response: 200
Stopping web server
Server stopped

[Main Module]
All systems operational
```

### Japanese Output
```
=== Modular Go-L10n Example ===

Language: ja
----------------------------------------
[Main Module]
アプリケーションを開始します
モジュールを読み込んでいます...

[Database Module]
データベースに接続しています...
正常に接続しました
クエリを実行しています: SELECT * FROM users
クエリが42msで完了しました
データベース接続を閉じています
接続を閉じました

[Web Module]
ポート8080でWebサーバーを起動しています
サーバーが正常に起動しました
リクエストを処理しています: /api/users
リクエスト完了: 125 ms
レスポンスを送信: 200
Webサーバーを停止しています
サーバーが停止しました

[Main Module]
すべてのシステムが正常に動作しています
```

## Best Practices Demonstrated

1. **Package-level `init()` functions**: Each package registers its translations automatically when imported
2. **Domain-specific translations**: Each module only includes translations relevant to its functionality
3. **Consistent key naming**: Use descriptive, context-aware translation keys
4. **Format string support**: Use `%s`, `%d` etc. for dynamic content in translations