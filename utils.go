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
	KHz         = 1000 * Hz
	MHz         = 1000 * KHz
	GHz         = 1000 * MHz
)

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

func convertKhzToMhz(freq uint64) uint64 {
	return freq / uint64(KHz)
}

func convertKhzToGhz(freq uint64) float64 {
	return float64(freq) / MHz
}

func readSingleValueFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}
