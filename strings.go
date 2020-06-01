package utils

import (
	"regexp"
	"strings"
)

// ElideString func
func ElideString(s string, maxLen int) string {
	s = strings.Replace(s, "\n", "", -1)

	if maxLen <= 0 {
		return s
	}
	if len(s) <= maxLen {
		return s
	}
	if maxLen < 4 {
		return s[0:maxLen]
	}
	s = s[0:maxLen-3] + "..."

	return s
}

// GetLastNameString func
func GetLastNameString(str string, sep string) (name string) {
	idx := strings.LastIndex(str, sep)
	if idx < 0 {
		return ""
	}
	name = str[idx+1:]
	return name
}

// StringListContain func
func StringListContain(sl []string, str string) bool {
	rc := false
	for _, s := range sl {
		if strings.Contains(s, str) {
			rc = true
			break
		}
	}
	return rc
}

// IsAlphaNumber func, string form is alpha or number
func IsAlphaNumber(str string) bool {
	return regexp.MustCompile(`^[a-z0-9A-Z]+$`).MatchString(str)
}
