package mode

type Mode byte

const (
	Waiting      = iota
	AddOrigin    = iota
	AddTranslate = iota
	ReadAll      = iota
)
