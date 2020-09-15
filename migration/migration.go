package migration

type Migration interface {
	GetVersion() int64
	GetDescription() string
	Up(b *Builder)
	Down(b *Builder)
}
