package smtcp

import "sync"

type ActivePeer struct {
	Peer      *Peer
	LambdaSet *LambdaSet
	Closed    bool
	WaitGroup *sync.WaitGroup
}

func NewTcpActivePeer(host string, lambdaSet *LambdaSet) (*ActivePeer, error) {
	peer, err := NewTcpPeer(host)
	return NewActivePeer(peer, lambdaSet), err
}

func NewActivePeer(peer *Peer, lambdaSet *LambdaSet) *ActivePeer {
	return &ActivePeer{Peer: peer, LambdaSet: lambdaSet, Closed: false, WaitGroup: &sync.WaitGroup{}}
}

func NewActivePipe(lambdaSet *LambdaSet) (*ActivePeer, *ActivePeer) {
	peerOne, peerTwo := NewPipe()
	return NewActivePeer(peerOne, lambdaSet), NewActivePeer(peerTwo, lambdaSet)
}

func (ap *ActivePeer) Close() error {
	ap.Closed = true
	return ap.Peer.Close()
}

func (ap *ActivePeer) Work() {
	ap.WaitGroup.Add(1)
	go func() {
		for {
			if ap.Closed {
				break
			}
			ap.loop()
		}
		ap.WaitGroup.Done()
	}()
}

func (ap *ActivePeer) Wait() {
	ap.WaitGroup.Wait()
}

func (ap *ActivePeer) Send(r *Request) error {
	return ap.Peer.Send(r)
}

func (ap *ActivePeer) loop() {
	if ap.Closed {
		return
	}

	r, err := ap.Peer.Receive()
	if err == nil {
		ap.LambdaSet.Execute(r)
	}
}
