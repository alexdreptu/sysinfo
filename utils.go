package sysinfo

import (
	"io/ioutil"
	"strings"
)

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
)

const (
	Hz  float64 = 1
	Khz         = 1000 * Hz
	Mhz         = 1000 * Khz
	Ghz         = 1000 * Mhz
)

func convertInt8ArrayToString(s [65]int8) string {
	b := make([]byte, len(s))
	for i, v := range s {
		b[i] = byte(v)
	}
	return string(b)
}

func convertBToKB(size uint64) uint64 {
	return size / uint64(KB)
}

func convertBToMB(size uint64) float64 {
	return float64(size) / MB
}

func convertBToGB(size uint64) float64 {
	return float64(size) / GB
}

func convertKBToB(size uint64) uint64 {
	return size * uint64(KB)
}

func convertKBToMB(size uint64) float64 {
	return float64(size) / KB
}

func convertMBToB(size float64) uint64 {
	return uint64(size * MB)
}

func convertGBToB(size float64) uint64 {
	return uint64(size * GB)
}

func convertKhzToMhz(freq uint64) uint64 {
	return freq / uint64(Khz)
}

func convertKhzToGhz(freq uint64) float64 {
	return float64(freq) / Mhz
}

func readSingleValueFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}
