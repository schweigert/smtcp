package smtcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParams(t *testing.T) {
	p := NewParams()

	assert.NotNil(t, p)

	data := p.Get("chave")
	assert.Equal(t, "", data)

	p.Set("chave", "raw data")
	data = p.Get("chave")

	assert.Equal(t, "raw data", data)
}
