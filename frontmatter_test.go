package frontmatter

import (
	"testing"

	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"
)

var ex1 = `
---

title: This is a title!

name: Derek Worthen
age: young
contact:
email: email@domain.com
address: some location
pets:
- cat
- dog
- bat
match: !!js/regexp /pattern/gim
run: !!js/function function() { }



---

- item
- item
- item`

func TestExtractFrontMatter(t *testing.T) {
	expected := `title: This is a title!

name: Derek Worthen
age: young
contact:
email: email@domain.com
address: some location
pets:
- cat
- dog
- bat
match: !!js/regexp /pattern/gim
run: !!js/function function() { }`
	assert.Equal(t, expected, Extract(ex1))
}

func TestParseFrontMatter(t *testing.T) {
	parsed, err := Parse(ex1)
	assert.NoError(t, err)
	t.Log(pp.Sprintln(parsed))
}

func TestTrimFrontMatter(t *testing.T) {
	expected := `- item
- item
- item`
	assert.Equal(t, expected, Trim(ex1))
}
