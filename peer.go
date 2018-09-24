package smtcp

import "net"

type Peer struct {
	conn net.Conn
}

func NewPeer(conn net.Conn) *Peer {
	return &Peer{conn: conn}
}

func NewPipe() (*Peer, *Peer) {
	p1, p2 := net.Pipe()
	return NewPeer(p1), NewPeer(p2)
}
