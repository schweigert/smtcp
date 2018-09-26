package smtcp

import "net"

type ActivePeerListener struct {
	Listener  net.Listener
	LambdaSet *LambdaSet
}

func NewTcpActiveListener(port string, lambdaSet *LambdaSet) (*ActivePeerListener, error) {
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	return NewActivePeerListener(listener, lambdaSet), err
}

func NewActivePeerListener(listener net.Listener, lambdaSet *LambdaSet) *ActivePeerListener {
	return &ActivePeerListener{Listener: listener, LambdaSet: lambdaSet}
}

func (l *ActivePeerListener) Accept() *ActivePeer {
	for {
		conn, err := l.Listener.Accept()
		if err != nil {
			continue
		}

		return NewActivePeer(NewPeer(conn), l.LambdaSet)
	}
}

func (l *ActivePeerListener) Close() error {
	return l.Listener.Close()
}
