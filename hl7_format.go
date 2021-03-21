package hl7reader

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	spRegexp = regexp.MustCompile(`\\.sp(\d*)\\`)
	skRegexp = regexp.MustCompile(`\\.sk(\d*)\\`)
	tiRegexp = regexp.MustCompile(`\\.ti(\d*)\\`)
	inRegexp = regexp.MustCompile(`\\.in(\d*)\\`)
)

func FormatString(str string) string {
	// These are simple 1:1 replacement patterns, so we don't need to utilize
	// regular expressions here.
	out := strings.NewReplacer(
		`\H\`, "",
		`\N\`, "",
		`\.fi\`, "",
		`\.nf\`, "",
		`\F\`, "|",
		`\S\`, "^",
		`\T\`, "&",
		`\R\`, "~",
		`\E\`, `\`,
		`\.br\`, "\n",
		`\.ce\`, "\n",
	).Replace(string(str))

	out = formatSp(out)
	out = formatSkipSpaces(skRegexp, out)
	out = formatSkipSpaces(tiRegexp, out)
	out = formatSkipSpaces(inRegexp, out)

	return out
}

// Replace \.sp(\d*)\ patterns. This is defined as a newline and then some
// number of spaces.
func formatSp(str string) string {
	return spRegexp.ReplaceAllStringFunc(str, func(str string) string {
		match := str[4 : len(str)-1]

		if match == "" {
			return "\n"
		}

		return "\n" + parseRepetition(match, " ")
	})
}

// Replace \.sk(\d*)\ patterns. This is defined as just skipping a number of
// spaces to the right.
func formatSkipSpaces(re *regexp.Regexp, str string) string {
	return re.ReplaceAllStringFunc(str, func(str string) string {
		match := str[4 : len(str)-1]

		if match == "" {
			return ""
		}

		return parseRepetition(match, " ")
	})
}

func parseRepetition(numStr, repeatStr string) string {
	count, err := strconv.Atoi(numStr)
	if err != nil {
		return ""
	}

	return strings.Repeat(repeatStr, count)
}
