package common

type Enum interface {
	String()
	Ordinal()
}

type OvercastLevel int

const (
	Clear OvercastLevel = iota
	PartlySunny
	PartlyCloudy
	Cloudy
)

func (o OvercastLevel) String() {

}

func (o OvercastLevel) Ordinal() {

}

type TrailStatus int

const (
	Closed TrailStatus = iota
	OpenGroomed
	OpenUngroomed
)

func (t TrailStatus) String() {

}

func (t TrailStatus) Ordinal() {

}
