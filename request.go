package smtcp

type Request struct {
	Name   string
	Params *Params
	Peer   *Peer
}

func (p *Peer) Send(*Request) error {
	return nil
}

func (p *Peer) Receive() (*Request, error) {
	return nil, nil
}
