package smtcp

import (
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActivePeerListenerAccept(t *testing.T) {
	lambdaSet := NewLambdaSet()
	listener, err := NewTcpActiveListener("3030", lambdaSet)
	assert.NoError(t, err)
	assert.NotNil(t, listener)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		conn, errConn := net.Dial("tcp", "localhost:3030")
		assert.NoError(t, errConn)
		errConn = conn.Close()
		assert.NoError(t, errConn)
		wg.Done()
	}()

	peer := listener.Accept()
	wg.Wait()

	assert.NotNil(t, peer)
	assert.Equal(t, lambdaSet, peer.LambdaSet)

	err = peer.Close()
	assert.NoError(t, err)

	err = listener.Close()
	assert.NoError(t, err)
}
