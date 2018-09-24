package smtcp

type Params struct {
	values map[string]string
}

func NewParams() *Params {
	return &Params{values: make(map[string]string)}
}

func (p *Params) Get(key string) string {
	v, ok := p.values[key]
	if !ok {
		return ""
	}
	return v
}

func (p *Params) Set(key, value string) {
	p.values[key] = value
}
