package smtcp

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPeer(t *testing.T) {
	p1, p2 := NewPipe()

	assert.NotNil(t, p1)
	assert.NotNil(t, p2)
}

func TestPeerReadAndWriteBytes(t *testing.T) {
	p1, p2 := NewPipe()
	wg := &sync.WaitGroup{}

	wg.Add(1)

	go func() {
		err1 := p1.writeBytes([]byte("hello"))
		assert.NoError(t, err1)
		wg.Done()
	}()

	str, err2 := p2.readBytes(5)
	wg.Wait()

	assert.NoError(t, err2)
	assert.Equal(t, []byte("hello"), str)
}

func TestPeerSend(t *testing.T) {
	params := NewParams()
	p1, p2 := NewPipe()

	defer p1.Close()
	defer p2.Close()

	request := NewRequest("request_name", params, p1)

	go func() {
		p1.Send(request)
	}()

	expected := "\f\x00\x00\x00request_name\x00\x00\x00\x00"
	actual, err := p2.readBytes(20)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))
}

func TestPeerReceive(t *testing.T) {
	params := NewParams()
	params.Set("foo1", "bar1")
	params.Set("foo2", "bar2")
	p1, p2 := NewPipe()

	defer p1.Close()
	defer p2.Close()

	request := NewRequest("request_name", params, p1)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		err1 := p1.Send(request)
		assert.NoError(t, err1)
		wg.Done()
	}()

	r, err := p2.Receive()
	assert.NoError(t, err)
	assert.NotNil(t, r)

	assert.Equal(t, "bar1", r.Params.Get("foo1"))
	assert.Equal(t, "bar2", r.Params.Get("foo2"))
	wg.Wait()
}
