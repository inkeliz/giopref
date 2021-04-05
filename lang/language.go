package lang

import (
	"golang.org/x/text/language"
)

var (
	// DefaultLanguage defines the default language, used by default when impossible to get the preferred language.
	DefaultLanguage = language.English
)

// Language returns the preferred language of the current user.
func Language() language.Tag {
	l := getLanguage()
	if l == "" {
		return DefaultLanguage
	}

	tag, err := language.Parse(l)
	if err != nil {
		return DefaultLanguage
	}

	return tag
}

// Match returns the best language based on the matcher.
func Match(matcher language.Matcher) (tag language.Tag, index int) {
	return language.MatchStrings(matcher, getLanguage())
}
