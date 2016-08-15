package faces

import "github.com/kovetskiy/lorg"

type Face interface {
	Init() error
	GetVersion() (string, error)
	SetLogger(lorg.Logger)
}
