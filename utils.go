package sysinfo

func ConvertInt8ArrayToString(s [65]int8) string {
	b := make([]byte, len(s))
	for i, v := range s {
		b[i] = byte(v)
	}
	return string(b)
}

func ConvertBToKB(size uint64) uint64 {
	return size / uint64(KB)
}

func ConvertBToMB(size uint64) float64 {
	return float64(size) / MB
}

func ConvertBToGB(size uint64) float64 {
	return float64(size) / GB
}

func ConvertKBToB(size uint64) uint64 {
	return size * uint64(KB)
}

func ConvertMBToB(size uint64) uint64 {
	return size * uint64(MB)
}

func ConvertGBToB(size float64) uint64 {
	return uint64(size * GB)
}
