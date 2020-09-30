package migration

import (
	"time"
)

type (
	Version int64

	Migration interface {
		GetVersion() Version
		GetDescription() string
		Up(b *Builder)
		Down(b *Builder)
	}
)

func NewVersion() Version{
	return Version(time.Now().Unix())
}
