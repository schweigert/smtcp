package smtcp

type Request struct {
	Name   string
	Params *Params
	Peer   *Peer
}

func NewRequest(name string, params *Params, peer *Peer) *Request {
	return &Request{
		Name:   name,
		Params: params,
		Peer:   peer,
	}
}

func NewRequestFromPeer(p *Peer) (*Request, error) {
	name, err := p.readString()
	if err != nil {
		return nil, err
	}

	paramsSize, err := p.readBytes(4)
	if err != nil {
		return nil, err
	}

	params := NewParams()

	for i := 0; i < int(BytesToUint(paramsSize)); i++ {
		key, err := p.readString()
		if err != nil {
			return nil, err
		}

		value, err := p.readString()
		if err != nil {
			return nil, err
		}
		params.Set(key, value)
	}

	return &Request{Name: name, Params: params, Peer: p}, nil
}

func (r *Request) Envelope() []byte {
	envelope := ""

	name := r.Name
	nameSize := string(Uint32ToBytes(uint32(len(r.Name))))
	paramsSize := string(Uint32ToBytes(uint32(len(r.Params.values))))

	envelope = envelope + nameSize
	envelope = envelope + name
	envelope = envelope + paramsSize

	for key, el := range r.Params.values {
		keySize := string(Uint32ToBytes(uint32(len(key))))
		elSize := string(Uint32ToBytes(uint32(len(el))))

		envelope = envelope + keySize
		envelope = envelope + key
		envelope = envelope + elSize
		envelope = envelope + el
	}

	return []byte(envelope)
}

func (r *Request) Send() {
	r.Peer.Send(r)
}
