package weblinks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_should_set_self_with_correct_params(t *testing.T) {
	links, _ := Create("http://www.google.com/", 1, 10, 100)
	assert.NotNil(t, links.Self)

	q := links.Self.Query()
	assert.NotNil(t, q)

	assert.Equal(t, "1", q.Get("page"))
	assert.Equal(t, "10", q.Get("page_size"))
}

func Test_should_set_next_with_params(t *testing.T) {
	links, _ := Create("http://www.google.com/", 1, 10, 100)
	assert.NotNil(t, links.Next)

	q := links.Next.Query()
	assert.NotNil(t, q)

	assert.Equal(t, "2", q.Get("page"))
	assert.Equal(t, "10", q.Get("page_size"))
}

func Test_should_set_prev_with_correct_params(t *testing.T) {
	links, _ := Create("http://www.google.com/", 2, 10, 100)
	assert.NotNil(t, links.Prev)

	q := links.Prev.Query()
	assert.NotNil(t, q)

	assert.Equal(t, "1", q.Get("page"))
	assert.Equal(t, "10", q.Get("page_size"))
}

func Test_should_set_first_with_correct_params(t *testing.T) {
	links, _ := Create("http://www.google.com/", 1, 10, 100)
	assert.NotNil(t, links.First)

	q := links.First.Query()
	assert.NotNil(t, q)

	assert.Equal(t, "1", q.Get("page"))
	assert.Equal(t, "10", q.Get("page_size"))
}

func Test_should_set_last_with_correct_params(t *testing.T) {
	links, _ := Create("http://www.google.com/", 1, 10, 100)
	assert.NotNil(t, links.Last)

	q := links.Last.Query()
	assert.NotNil(t, q)

	assert.Equal(t, "10", q.Get("page"))
	assert.Equal(t, "10", q.Get("page_size"))
}

func Test_should_not_set_prev_if_page_less_than_2(t *testing.T) {
	links, _ := Create("http://www.google.com/", 1, 10, 100)

	assert.Nil(t, links.Prev)
}

func Test_should_not_set_next_if_on_last_page(t *testing.T) {
	links, _ := Create("http://www.google.com/", 10, 10, 100)

	assert.Nil(t, links.Next)
}