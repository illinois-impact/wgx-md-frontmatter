package frontmatter

import (
	"errors"
	"regexp"
	"runtime"
	"strings"
	"unicode"

	"gopkg.in/yaml.v2"
)

var (
	regex            *regexp.Regexp
	ErrNoFrontMatter = errors.New("Frontmatter not found")
)

func hasFrontMatter(md string) bool {
	md = strings.TrimLeftFunc(md, unicode.IsSpace)
	lines := strings.Split(md, "\n")
	if len(lines) > 0 && (strings.HasPrefix(lines[0], "= yaml =") ||
		strings.HasPrefix(lines[0], "---")) {
		return regex.MatchString(md)
	}
	return false
}

func trimLeft(md string) string {
	return strings.TrimLeftFunc(md, unicode.IsSpace)
}
func submatches(md string) []string {
	return regex.FindStringSubmatch(md)
}

func extract(md string) string {
	matches := submatches(md)
	if len(matches) <= 4 {
		return ""
	}
	return strings.TrimSpace(matches[3])
}

func Extract(md string) string {
	md = trimLeft(md)
	if hasFrontMatter(md) {
		return extract(md)
	} else {
		return ""
	}
}

func Parse(md string) (map[string]interface{}, error) {
	front := Extract(md)
	if front == "" {
		return map[string]interface{}{}, ErrNoFrontMatter
	}
	res := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(front), &res)
	return res, err
}

func Trim(md string) string {
	md = trimLeft(md)
	if !hasFrontMatter(md) {
		return md
	}
	matches := submatches(md)
	if len(matches) <= 4 {
		return md
	}
	return trimLeft(md[len(matches[0]):])
}

func init() {
	var lineend = ""
	if runtime.GOOS == "windows" {
		lineend = `\\r?`
	}
	// see syntax : https://github.com/google/re2/wiki/Syntax
	// also the tool at : https://regex-golang.appspot.com/
    // regexp is based on https://github.com/jxson/front-matter
	pat := `(?m)^(` +
		`(= yaml =|---)` +
		`$([\s\S]*?)` +
		`(?:((---)|(\.\.\.)))` +
		`$` +
		lineend +
		`(?:\\n)?)` +
		`(.*)`
	regex = regexp.MustCompile(pat)
}
