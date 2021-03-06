package check

import (
	"strings"
	"testing"

	"github.com/jdkato/regexp"
	"github.com/stretchr/testify/assert"
)

type changeCase struct {
	match      bool
	heading    string
	exceptions []string
	indicators []string
}

func TestSentence(t *testing.T) {
	headings := []changeCase{
		{heading: "Top-level entities", match: true},
		{heading: "Non-member Predicates", match: false},
		{heading: "Non-member predicates", match: true},
		{heading: "Client's key share and finish", match: true},
		{heading: "Client’s key share and finish", match: true},
		{
			heading:    "Find the thief: Introduction",
			match:      true,
			indicators: []string{":"},
		},
		{
			heading: "Find the thief: Introduction",
			match:   false,
		},
		{
			heading:    "Creating a connection to Event Store",
			match:      true,
			exceptions: []string{"Event Store"},
		},
	}

	for _, h := range headings {
		var r *regexp.Regexp
		if len(h.exceptions) > 0 {
			r = regexp.MustCompile(strings.Join(h.exceptions, "|"))
		}
		assert.Equal(t, sentence(
			h.heading, h.exceptions, h.indicators, r), h.match)
	}
}
