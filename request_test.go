package smtcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T) {
	params := NewParams()
	p1, p2 := NewPipe()

	defer p1.Close()
	defer p2.Close()

	request := NewRequest("request_name", params, p1)

	actual := request.Envelope()
	expected := "\f\x00\x00\x00request_name\x00\x00\x00\x00"
	assert.Equal(t, expected, string(actual))

	params.Set("foo1", "bar1")
	params.Set("foo2", "bar2")

	actual = request.Envelope()
	expected_opt1 := "\f\x00\x00\x00request_name\x02\x00\x00\x00\x04\x00\x00\x00foo1\x04\x00\x00\x00bar1\x04\x00\x00\x00foo2\x04\x00\x00\x00bar2"
	expected_opt2 := "\f\x00\x00\x00request_name\x02\x00\x00\x00\x04\x00\x00\x00foo2\x04\x00\x00\x00bar2\x04\x00\x00\x00foo1\x04\x00\x00\x00bar1"
	assert.True(t, expected_opt1 == string(actual) || expected_opt2 == string(actual))
}
