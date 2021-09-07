package posts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostFromMap(t *testing.T) {
	m := map[string]string{
		"name":    "Foo",
		"message": "Bar",
		"terms":   "true",
	}

	p := FromMap(m)

	assert.Equal(t, m["name"], p.Name)
	assert.Equal(t, m["message"], p.Message)
	assert.Equal(t, m["terms"], p.Terms)
}

func TestPostToMap(t *testing.T) {
	p := Post{
		Name:    "foo",
		Message: "bar",
		Terms:   "true",
	}

	m := p.ToMap()

	assert.Equal(t, m["name"], p.Name)
	assert.Equal(t, m["message"], p.Message)
	assert.Equal(t, m["terms"], p.Terms)
}
