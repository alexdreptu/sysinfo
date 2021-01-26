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

func (s Size) ToBytes() float64 {
	return float64(s / Byte)
}

func (s Size) ToKibibytes() float64 {
	return float64(s / Kibibyte)
}

func (s Size) ToMebibytes() float64 {
	return float64(s / Mebibyte)
}

func (s Size) ToGibibytes() float64 {
	return float64(s / Gibibyte)
}

func (f Frequency) ToHertz() float64 {
	return float64(f / Hertz)
}

func (f Frequency) ToKilohertz() float64 {
	return float64(f / Kilohertz)
}

func (f Frequency) ToMegahertz() float64 {
	return float64(f / Megahertz)
}

func (f Frequency) ToGigahertz() float64 {
	return float64(f / Gigahertz)
}
