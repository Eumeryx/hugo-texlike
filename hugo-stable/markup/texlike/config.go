package texlike

import (
	tl "github.com/Eumeryx/goldmark-texlike"
)

type Config struct {
	Enable bool
	Option tl.Config
}

var DefaultConfig = Config{
	Enable: true,
	Option: tl.DefaultConfig,
}

func NewTexlike(c Config) *tl.Texlike {
	return tl.NewTexlike(c.Option)
}
