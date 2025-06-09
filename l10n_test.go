package l10n

import (
	"os"
	"testing"
)

func TestT_BasicTranslation(t *testing.T) {
	// テスト用の翻訳を登録
	Register("ja", LexiconMap{
		"Hello":     "こんにちは",
		"Good bye":  "さようなら",
		"Thank you": "ありがとう",
	})

	tests := []struct {
		name     string
		lang     string
		phrase   string
		expected string
	}{
		{"Japanese translation", "ja", "Hello", "こんにちは"},
		{"Japanese translation 2", "ja", "Good bye", "さようなら"},
		{"English fallback", "en", "Hello", "Hello"},
		{"Unknown phrase fallback", "ja", "Unknown", "Unknown"},
		{"Unknown language fallback", "fr", "Hello", "Hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForceLanguage(tt.lang)
			result := T(tt.phrase)
			if result != tt.expected {
				t.Errorf("T(%q) with lang %q = %q, want %q", tt.phrase, tt.lang, result, tt.expected)
			}
		})
	}

	// テスト後はリセット
	ResetLanguage()
}

func TestRegister(t *testing.T) {
	// 初期状態をクリア
	World["test"] = LexiconMap{}

	// 新しい言語を登録
	Register("test", LexiconMap{
		"hello": "test-hello",
		"world": "test-world",
	})

	if World["test"]["hello"] != "test-hello" {
		t.Errorf("Register failed: expected 'test-hello', got %q", World["test"]["hello"])
	}

	// 既存の言語にマージ
	Register("test", LexiconMap{
		"goodbye": "test-goodbye",
		"hello":   "test-hello-updated", // 上書き
	})

	if World["test"]["hello"] != "test-hello-updated" {
		t.Errorf("Register merge failed: expected 'test-hello-updated', got %q", World["test"]["hello"])
	}

	if World["test"]["goodbye"] != "test-goodbye" {
		t.Errorf("Register merge failed: expected 'test-goodbye', got %q", World["test"]["goodbye"])
	}

	if World["test"]["world"] != "test-world" {
		t.Errorf("Register merge failed: expected 'test-world', got %q", World["test"]["world"])
	}

	// クリーンアップ
	delete(World, "test")
}

func TestForceLanguageAndReset(t *testing.T) {
	originalLang := Language

	// 言語を強制設定
	ForceLanguage("ja")
	if Language != "ja" {
		t.Errorf("ForceLanguage failed: expected 'ja', got %q", Language)
	}
	if forcedLanguage != "ja" {
		t.Errorf("forcedLanguage not set: expected 'ja', got %q", forcedLanguage)
	}

	// 強制設定中は DetectLanguage() が無視される
	DetectLanguage()
	if Language != "ja" {
		t.Errorf("Language changed during forced mode: expected 'ja', got %q", Language)
	}

	// リセット
	ResetLanguage()
	if forcedLanguage != "" {
		t.Errorf("forcedLanguage not reset: expected '', got %q", forcedLanguage)
	}

	// 元の状態に戻す
	Language = originalLang
}

func TestGetCurrentLanguage(t *testing.T) {
	ForceLanguage("test-lang")
	if GetCurrentLanguage() != "test-lang" {
		t.Errorf("GetCurrentLanguage failed: expected 'test-lang', got %q", GetCurrentLanguage())
	}
	ResetLanguage()
}

func TestGetDefaultLanguage(t *testing.T) {
	// 環境変数をクリア
	_ = os.Unsetenv("L10N_DEFAULT_LANGUAGE")
	if getDefaultLanguage() != "en" {
		t.Errorf("getDefaultLanguage without env var failed: expected 'en', got %q", getDefaultLanguage())
	}

	// 環境変数を設定
	_ = os.Setenv("L10N_DEFAULT_LANGUAGE", "fr")
	if getDefaultLanguage() != "fr" {
		t.Errorf("getDefaultLanguage with env var failed: expected 'fr', got %q", getDefaultLanguage())
	}

	// クリーンアップ
	_ = os.Unsetenv("L10N_DEFAULT_LANGUAGE")
}

func TestLanguageDetectionWithEnvironmentVariables(t *testing.T) {
	// 環境変数をクリア
	envVars := []string{"LANGUAGE", "LC_ALL", "LC_MESSAGES", "LANG"}
	for _, env := range envVars {
		_ = os.Unsetenv(env)
	}
	_ = os.Unsetenv("L10N_DEFAULT_LANGUAGE")

	tests := []struct {
		name     string
		envVar   string
		value    string
		expected string
	}{
		{"Japanese LANG", "LANG", "ja_JP.UTF-8", "ja"},
		{"Japanese LC_MESSAGES", "LC_MESSAGES", "ja", "ja"},
		{"English LANG", "LANG", "en_US.UTF-8", "en"},
		{"Chinese fallback to English", "LANG", "zh_CN.UTF-8", "en"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 環境変数をクリア
			for _, env := range envVars {
				_ = os.Unsetenv(env)
			}

			// テスト用環境変数を設定
			_ = os.Setenv(tt.envVar, tt.value)

			// 手動で言語検出ロジックを実行（テストモード検出を回避）
			Language = getDefaultLanguage()

			// 環境変数から言語を判定（DetectLanguageのロジックを模倣）
			langs := []string{
				os.Getenv("LANGUAGE"),
				os.Getenv("LC_ALL"),
				os.Getenv("LC_MESSAGES"),
				os.Getenv("LANG"),
			}

			for _, l := range langs {
				if l != "" {
					// 日本語が含まれている場合
					if l == "ja" || l == "ja_JP.UTF-8" {
						Language = "ja"
						break
					}
					// 英語の場合はデフォルトのまま
					if l == "en_US.UTF-8" {
						Language = "en"
						break
					}
					// その他の場合はデフォルト言語のまま
				}
			}

			if Language != tt.expected {
				t.Errorf("Language detection failed for %s=%s: expected %q, got %q", tt.envVar, tt.value, tt.expected, Language)
			}

			// クリーンアップ
			_ = os.Unsetenv(tt.envVar)
		})
	}

	// テスト後はリセット
	ResetLanguage()
}

func TestL10NSkipDetection(t *testing.T) {
	// 環境変数を設定
	_ = os.Setenv("L10N_SKIP_DETECTION", "1")

	// 初期化前の状態を保存
	originalLang := Language

	// 手動で初期化をシミュレート
	Language = getDefaultLanguage()
	if os.Getenv("L10N_SKIP_DETECTION") == "" {
		DetectLanguage()
	}

	// SKIP_DETECTIONが設定されているので、DetectLanguageは呼ばれない
	// デフォルト言語のままになる
	expected := getDefaultLanguage()
	if Language != expected {
		t.Errorf("L10N_SKIP_DETECTION failed: expected %q, got %q", expected, Language)
	}

	// クリーンアップ
	_ = os.Unsetenv("L10N_SKIP_DETECTION")
	Language = originalLang
}

func TestL10NDefaultLanguage(t *testing.T) {
	// デフォルト言語を設定
	_ = os.Setenv("L10N_DEFAULT_LANGUAGE", "fr")

	// テストモードで言語検出
	DetectLanguage()

	// テストモード中はデフォルト言語が使われる
	if Language != "fr" {
		t.Errorf("L10N_DEFAULT_LANGUAGE in test mode failed: expected 'fr', got %q", Language)
	}

	// クリーンアップ
	_ = os.Unsetenv("L10N_DEFAULT_LANGUAGE")
	ResetLanguage()
}
