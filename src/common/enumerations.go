package common

type IEnum interface {
	String() string
	Ordinal() int
	IsValid() bool
}

type OvercastLevel int
type TrailStatus int
type LiftStatus int
type TrailDifficulty int
type AvalancheAspect int
type AvalancheElevation int
type AvalancheDanger int
type AvalancheSize int
type AvalancheProblemType int
type AvalancheLikelihood int

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

const (
	North AvalancheAspect = 1 << iota
	NorthWest
	West
	SouthWest
	South
	SouthEast
	East
	NorthEast
)

const (
	BelowTreeline AvalancheElevation = 1 << iota
	AtTreeLine
	AboveTreeline
)

const (
	NoRating AvalancheDanger = iota
	Low
	Moderate
	Considerable
	High
	Extreme
)

const (
	D1 AvalancheSize = iota
	D2
	D3
	D4
	D5
)

const (
	WindSlab AvalancheProblemType = 1 << iota
	StormSlab
	PersistentSlab
	LooseDry
	PersistentWeakLayer
	CorniceFall
	Glide
	WetSnow
)

const (
	Unlikely AvalancheLikelihood = iota + 1
	Possible
	Likely
	VeryLikely
	Certain
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
