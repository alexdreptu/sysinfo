package convert

type (
	Unit float64
	Size Unit
	Freq Unit
)

const (
	Byte     Unit = 1
	Kibibyte      = 1024 * Byte
	Mebibyte      = 1024 * Kibibyte
	Gibibyte      = 1024 * Mebibyte
)

const (
	Hertz     Unit = 1
	Kilohertz      = 1000 * Hertz
	Megahertz      = 1000 * Kilohertz
	Gigahertz      = 1000 * Megahertz
)

func (u Unit) Bytes() float64 {
	return float64(u / Byte)
}

func (u Unit) Kibibytes() float64 {
	return float64(u / Kibibyte)
}

func (u Unit) Mebibytes() float64 {
	return float64(u / Mebibyte)
}

func (u Unit) Gibibytes() float64 {
	return float64(u / Gibibyte)
}

func (u Unit) Megahertz() float64 {
	return float64(u / Megahertz)
}

func (u Unit) Gigahertz() float64 {
	return float64(u / Gigahertz)
}
