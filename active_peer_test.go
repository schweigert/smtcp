package smtcp

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	defer activePeerOne.Close()
	defer activePeerTwo.Close()

	activePeerOne.Work()
	activePeerTwo.Work()

	NewRequest("one", NewParams(), activePeerOne.Peer).Send()
	wg.Wait()

	assert.Equal(t, 11, oneCalls)
	assert.Equal(t, 11, twoCalls)
}
