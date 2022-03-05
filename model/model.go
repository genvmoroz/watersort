package model

import "strconv"

type (
	Tripod struct {
		Flasks [][]Flask
	}

	Flask struct {
		index   int8
		colours [4]Colour
	}

	Colour int
)

const (
	Non Colour = iota
	Gray
	Red
	LightGreen
	Green
	DarkGreen
	LightBlue
	Blue
	Purple
	Pink
	Orange
	Brown
	Yellow
)

func NewFlask() Flask {
	return Flask{colours: [4]Colour{Non, Non, Non, Non}, index: -1}
}

func (f *Flask) Put(colour Colour) bool {
	if f.Full() {
		return false
	}

	f.colours[f.index+1] = colour
	f.index++
	return true
}

func (f *Flask) Get() Colour {
	if f.Empty() {
		return Non
	}

	colour := f.colours[f.index]
	f.colours[f.index] = Non
	f.index--
	return colour
}

func (f Flask) Empty() bool {
	return f.index == -1
}

func (f Flask) Full() bool {
	if f.index == int8(len(f.colours)-1) {
		return true
	}

	return false
}

func (f Flask) Completed() bool {
	if f.Full() {
		col := f.colours[0]
		for i := 1; i < len(f.colours); i++ {
			if f.colours[i] != col {
				return false
			}
		}

		return true
	}

	return false
}

func (c Colour) String() string {
	switch c {
	case Non:
		return "Non"
	case Gray:
		return "Gray"
	case Red:
		return "Red"
	case LightGreen:
		return "LightGreen"
	case Green:
		return "Green"
	case DarkGreen:
		return "DarkGreen"
	case LightBlue:
		return "LightBlue"
	case Blue:
		return "Blue"
	case Purple:
		return "Purple"
	case Pink:
		return "Pink"
	case Orange:
		return "Orange"
	case Brown:
		return "Brown"
	case Yellow:
		return "Yellow"
	default:
		return strconv.Itoa(int(c))
	}
}
