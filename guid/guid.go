package guid

import "github.com/rs/xid"

type Guid struct{}

func (g *Guid) New() string {
	return xid.New().String()
}
