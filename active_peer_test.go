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
