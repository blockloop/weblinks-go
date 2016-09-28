package weblinks

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testHeader = `<http://www.google.com/?page=2&page_size=10>; rel="self",
                  <http://www.google.com/?page=3&page_size=10>; rel="next",
                  <http://www.google.com/?page=1&page_size=10>; rel="first",
                  <http://www.google.com/?page=1&page_size=10>; rel="prev",
                  <http://www.google.com/?page=10&page_size=10>; rel="last"`
)

var (
	testLinks, _ = Create("http://www.google.com/", 2, 10, 100)
)

func Test_parse_should_set_fields(t *testing.T) {
	links, _ := Parse(testHeader)
	assert.NotNil(t, links.Self)
	assert.Equal(t, links.Self, testLinks.Self)
	assert.Equal(t, links.Next, testLinks.Next)
	assert.Equal(t, links.Prev, testLinks.Prev)
	assert.Equal(t, links.First, testLinks.First)
	assert.Equal(t, links.Last, testLinks.Last)
}

func compareLinks(t *testing.T, x, y *WebLinks) {
	xj, _ := json.Marshal(x)
	yj, _ := json.Marshal(y)

	assert.Equal(t, xj, yj)
}
