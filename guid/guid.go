package guid

import "github.com/rs/xid"

type Guid struct{}

func New() *Guid {
	return &Guid{}
}

func (g *Guid) Generate() string {
	return xid.New().String()
}
