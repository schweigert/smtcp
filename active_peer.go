package smtcp

type ActivePeer struct {
	Peer      *Peer
	LambdaSet *LambdaSet
	Closed    bool
}

func NewActivePeer(peer *Peer, lambdaSet *LambdaSet) *ActivePeer {
	return &ActivePeer{Peer: peer, LambdaSet: lambdaSet, Closed: false}
}

func NewActivePipe(lambdaSet *LambdaSet) (*ActivePeer, *ActivePeer) {
	peerOne, peerTwo := NewPipe()
	return NewActivePeer(peerOne, lambdaSet), NewActivePeer(peerTwo, lambdaSet)
}

func (ap *ActivePeer) Close() {
	ap.Closed = true
	ap.Peer.Close()
}

func (ap *ActivePeer) Work() {
	go func() {
		for {
			if ap.Closed {
				return
			}
			ap.loop()
		}
	}()
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
