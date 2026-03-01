package constants

type Playstyle uint8

const (
	PlaystyleNotSpecified Playstyle = 0
	PlaystyleMouse        Playstyle = 1 << 0
	PlaystyleTablet       Playstyle = 1 << 1
	PlaystyleKeyboard     Playstyle = 1 << 2
	PlaystyleTouch        Playstyle = 1 << 3
)

func (p Playstyle) Has(flag Playstyle) bool {
	return p&flag != 0
}
