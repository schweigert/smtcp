package smtcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPeer(t *testing.T) {
	p1, p2 := NewPipe()

	assert.NotNil(t, p1)
	assert.NotNil(t, p2)
}
