package l10n

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/language"
)

// LexiconMap maps base phrases to translated phrases.
type LexiconMap map[string]string

// WorldMap maps language codes to their respective lexicons.
type WorldMap map[string]LexiconMap

var (
	// Currently active language.
	Language = getDefaultLanguage()

	// Forced language for testing (empty string means no forcing).
	forcedLanguage string
)

// DetectLanguage detects and sets the current language based on environment variables
// or forced settings. It respects the forced language setting and test mode.
func DetectLanguage() {
	// Use forced language if it's set
	if forcedLanguage != "" {
		Language = forcedLanguage
		return
	}

	// Check if running in test mode (defaults to default language during tests)
	if isTestMode() {
		Language = getDefaultLanguage()
		return
	}

	// Detect language from environment variables
	langs := []string{
		os.Getenv("LANGUAGE"),
		os.Getenv("LC_ALL"),
		os.Getenv("LC_MESSAGES"),
		os.Getenv("LANG"),
	}

	matcher := language.NewMatcher([]language.Tag{
		language.Japanese,
		language.English,
	})

	// Start with default language
	Language = getDefaultLanguage()

	// Switch to Japanese if specified in environment variables
	for _, l := range langs {
		if l != "" {
			tag, _ := language.MatchStrings(matcher, l)
			if tag == language.Japanese {
				Language = "ja"
				break
			}
		}
	}
}

// Register merges translation phrases into the phrase map
// Extension applications can use Register to add their own translation phrases.
func Register(lang string, lex LexiconMap) {
	l, ok := World[lang]
	if !ok {
		World[lang] = lex
		return
	}

	for k, v := range lex {
		l[k] = v
	}
}

// T translates a phrase.
func T(phrase string) string {
	l, ok := World[Language]
	if !ok {
		return phrase
	}

	t, ok := l[phrase]
	if !ok {
		return phrase
	}

	return t
}

// F translates a phrase and formats it with fmt.Sprintf
func F(phrase string, args ...interface{}) string {
	f := T(phrase)
	return fmt.Sprintf(f, args...)
}

// E translates a phrase and returns it as an error with fmt.Errorf
func E(phrase string, args ...interface{}) error {
	f := T(phrase)
	return fmt.Errorf(f, args...)
}

var (
	// Global phrase map
	World = WorldMap{
		"ja": LexiconMap{},
	}
)

// isTestMode checks if the code is running in test mode
func isTestMode() bool {
	// Check if test mode is explicitly set via environment variable
	if os.Getenv("L10N_TEST_MODE") == "1" {
		return true
	}

	// During Go test execution, the executable name contains .test
	// or the arguments contain test-related flags
	for _, arg := range os.Args {
		if strings.Contains(arg, ".test") ||
			strings.Contains(arg, "-test.") ||
			strings.HasSuffix(arg, "_test") {
			return true
		}
	}
	return false
}

// ForceLanguage overrides the language to a specific value (primarily for testing)
// This setting takes precedence over environment variable detection.
func ForceLanguage(lang string) {
	forcedLanguage = lang
	Language = lang
}

// ResetLanguage resets language detection to automatic mode
// This clears any forced language setting and re-runs language detection.
func ResetLanguage() {
	forcedLanguage = ""
	DetectLanguage()
}

// GetCurrentLanguage returns the currently active language code
// (e.g., "en", "ja")
func GetCurrentLanguage() string {
	return Language
}

// getDefaultLanguage returns the default language from L10N_DEFAULT_LANGUAGE
// environment variable, or "en" if not set
func getDefaultLanguage() string {
	defaultLang := os.Getenv("L10N_DEFAULT_LANGUAGE")
	if defaultLang == "" {
		return "en"
	}
	return defaultLang
}

func init() {
	if os.Getenv("L10N_SKIP_DETECTION") == "" {
		DetectLanguage()
	}
}
