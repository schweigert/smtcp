package smtcp

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTcpActivePeer(t *testing.T) {
	listener, errL := NewTcpActiveListener("3031", nil)
	assert.NoError(t, errL)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		peerL := listener.Accept()
		errL = listener.Close()
		assert.NoError(t, errL)

		errL = peerL.Close()
		assert.NoError(t, errL)
		wg.Done()
	}()

	peer, err := NewTcpActivePeer("localhost:3031", nil)
	assert.NoError(t, err)
	assert.NotNil(t, peer)
	wg.Wait()
}

func TestActivePeerWorkLoop(t *testing.T) {
	oneCalls := 0
	twoCalls := 0

	wg := &sync.WaitGroup{}

	wg.Add(2)

	lambdaSet := NewLambdaSet().Set("one", func(r *Request) {
		oneCalls++
		if oneCalls <= 10 {
			NewRequest("two", NewParams(), r.Peer).Send()
		} else {
			NewRequest("two", NewParams(), r.Peer).Send()
			wg.Done()
		}
	}).Set("two", func(r *Request) {
		twoCalls++
		if twoCalls <= 10 {
			NewRequest("one", NewParams(), r.Peer).Send()
		} else {
			wg.Done()
		}
	})
	activePeerOne, activePeerTwo := NewActivePipe(lambdaSet)

	activePeerOne.Work()
	activePeerTwo.Work()

	err := activePeerOne.Send(NewRequest("one", NewParams(), activePeerOne.Peer))
	assert.NoError(t, err)

	wg.Wait()

	errOne := activePeerOne.Close()
	errTwo := activePeerTwo.Close()

	assert.NoError(t, errOne)
	assert.NoError(t, errTwo)

	activePeerOne.Wait()
	activePeerTwo.Wait()
}
