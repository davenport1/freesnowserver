package common

type IEnum interface {
	String() string
	Ordinal() int
}

type OvercastLevel int
type TrailStatus int
type LiftStatus int
type TrailDifficulty int

const (
	Clear OvercastLevel = iota
	PartlySunny
	PartlyCloudy
	Cloudy
)

const (
	TrailClosed TrailStatus = iota
	OpenGroomed
	OpenUngroomed
)

const (
	LiftClosed LiftStatus = iota
	Scheduled
	Open
)

const (
	Green TrailDifficulty = iota
	Blue
	Black
	DoubleBlack
)

func (o OvercastLevel) String() string {
	return [...]string{"Clear", "PartlySunny", "PartlyCloudy", "Cloudy"}[o]
}

func (o OvercastLevel) Ordinal() int {
	return int(o)
}

func (t TrailStatus) String() string {
	return [...]string{"TrailClosed", "OpenGroomed", "OpenUngroomed"}[t]
}

func (t TrailStatus) Ordinal() int {
	return int(t)
}

func (l LiftStatus) String() string {
	return [...]string{"LiftClosed", "LiftScheduled", "LiftOpen"}[l]
}

func (l LiftStatus) Ordinal() int {
	return int(l)
}

func (d TrailDifficulty) String() string {
	return [...]string{"Green", "Blue", "Black", "DoubleBlack"}[d]
}

func (d TrailDifficulty) Ordinal() int {
	return int(d)
}
