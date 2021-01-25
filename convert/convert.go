package convert

type (
	Size      float64
	Frequency float64
)

const (
	Byte     Size = 1
	Kibibyte      = 1024 * Byte
	Mebibyte      = 1024 * Kibibyte
	Gibibyte      = 1024 * Mebibyte
)

const (
	Hertz     Frequency = 1
	Kilohertz           = 1000 * Hertz
	Megahertz           = 1000 * Kilohertz
	Gigahertz           = 1000 * Megahertz
)

func (s Size) Bytes() float64 {
	return float64(s / Byte)
}

func (s Size) Kibibytes() float64 {
	return float64(s / Kibibyte)
}

func (s Size) Mebibytes() float64 {
	return float64(s / Mebibyte)
}

func (s Size) Gibibytes() float64 {
	return float64(s / Gibibyte)
}

func (f Frequency) Hertz() float64 {
	return float64(f / Hertz)
}

func (f Frequency) Kilohertz() float64 {
	return float64(f / Kilohertz)
}

func (f Frequency) Megahertz() float64 {
	return float64(f / Megahertz)
}

func (f Frequency) Gigahertz() float64 {
	return float64(f / Gigahertz)
}
