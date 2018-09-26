package smtcp

import (
	"errors"
	"net"
)

type Peer struct {
	conn net.Conn
}

func NewTcpPeer(host string) (*Peer, error) {
	conn, err := net.Dial("tcp", host)
	return NewPeer(conn), err
}

func NewPeer(conn net.Conn) *Peer {
	return &Peer{conn: conn}
}

func NewPipe() (*Peer, *Peer) {
	p1, p2 := net.Pipe()
	return NewPeer(p1), NewPeer(p2)
}

func (p *Peer) Close() error {
	return p.conn.Close()
}

func (p *Peer) Send(r *Request) error {
	return p.writeBytes(r.Envelope())
}

func (p *Peer) Receive() (*Request, error) {
	return NewRequestFromPeer(p)
}

func (p *Peer) writeBytes(raw []byte) error {
	n, err := p.conn.Write([]byte(raw))
	if n != len(raw) {
		return errors.New("can not write entire package")
	}

	return err
}

func (p *Peer) readString() (string, error) {
	size, err := p.readBytes(4)
	if err != nil {
		return "", err
	}

	iSize := BytesToUint(size)

	if iSize == 0 {
		return "", nil
	}

	bytes, err := p.readBytes(iSize)
	return string(bytes), err
}

func (p *Peer) readBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	s, err := p.conn.Read(b)
	if n != uint32(s) {
		err = errors.New("can not read entire package")
	}

	return b, err
}
