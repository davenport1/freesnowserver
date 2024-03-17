package common

type IEnum interface {
	String() string
	Ordinal() int
	IsValid() bool
}

type IBitwiseEnum interface {
	String() []string
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
type AvalancheDangerAbove int
type AvalancheDangerAt int
type AvalancheDangerBelow int
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
	OnHold
	LiftScheduled
	LiftOpen
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

func (o OvercastLevel) IsValid() bool {
	return o >= Clear && o <= Cloudy
}

func (t TrailStatus) String() string {
	return [...]string{"TrailClosed", "OpenGroomed", "OpenUngroomed"}[t]
}

func (t TrailStatus) Ordinal() int {
	return int(t)
}

func (t TrailStatus) IsValid() bool {
	return t <= TrailClosed && t >= OpenGroomed
}

func (l LiftStatus) String() string {
	return [...]string{"LiftClosed", "OnHold", "LiftScheduled", "LiftOpen"}[l]
}

func (l LiftStatus) Ordinal() int {
	return int(l)
}

func (l LiftStatus) IsValid() bool {
	return l <= LiftClosed && l >= LiftOpen
}

func (d TrailDifficulty) String() string {
	return [...]string{"Green", "Blue", "Black", "DoubleBlack"}[d]
}

func (d TrailDifficulty) Ordinal() int {
	return int(d)
}

func (d TrailDifficulty) IsValid() bool {
	return d >= Green && d <= DoubleBlack
}

func (aspect AvalancheAspect) String() []string {
	var result []string
	for i := 0; i < 8; i++ { // Assuming there are 8 possible aspects
		if aspect&(1<<uint(i)) != 0 {
			result = append(result, [...]string{"North", "NorthWest", "West", "SouthWest", "South", "SouthEast", "East", "NorthEast"}[i])
		}
	}
	return result
}

func (aspect AvalancheAspect) Ordinal() int {
	return int(aspect)
}

func (aspect AvalancheAspect) IsValid() bool {
	return aspect >= North && aspect <= NorthEast
}

func (elevation AvalancheElevation) String() []string {
	var result []string
	for i := 0; i < 3; i++ { // Assuming there are 3 possible elevations
		if elevation&(1<<uint(i)) != 0 {
			result = append(result, [...]string{"BelowTreeline", "AtTreeLine", "AboveTreeLine"}[i])
		}
	}
	return result
}

func (elevation AvalancheElevation) Ordinal() int {
	return int(elevation)
}

func (elevation AvalancheElevation) IsValid() bool {
	return elevation >= BelowTreeline && elevation <= AboveTreeline
}

func (d AvalancheDanger) String() string {
	return [...]string{"NoRating", "Low", "Moderate", "Considerable", "High", "Extreme"}[d]
}

func (d AvalancheDanger) Ordinal() int {
	return int(d)
}

func (d AvalancheDanger) IsValid() bool {
	return d >= NoRating && d <= Extreme
}

func (s AvalancheSize) String() string {
	return [...]string{"D1", "D2", "D3", "D4", "D5"}[s]
}

func (s AvalancheSize) Ordinal() int {
	return int(s)
}

func (s AvalancheSize) IsValid() bool {
	return s >= D1 && s <= D5
}

func (p AvalancheProblemType) String() []string {
	var result []string
	for i := 0; i < 8; i++ { // Assuming there are 8 possible problem types
		if p&(1<<uint(i)) != 0 {
			result = append(result, [...]string{"WindSlab", "StormSlab", "PersistentSlab", "LooseDry", "PersistentWeakLayer", "CorniceFall", "Glide", "WetSnow"}[i])
		}
	}
	return result
}

func (p AvalancheProblemType) Ordinal() int {
	return int(p)
}

func (p AvalancheProblemType) IsValid() bool {
	return p >= WindSlab && p <= WetSnow
}

func (l AvalancheLikelihood) String() string {
	return [...]string{"Unlikely", "Possible", "Likely", "VeryLikely", "Certain"}[l-1]
}

func (l AvalancheLikelihood) Ordinal() int {
	return int(l)
}

func (l AvalancheLikelihood) IsValid() bool {
	return l >= Unlikely && l <= Certain
}
