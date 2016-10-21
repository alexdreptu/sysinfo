package sysinfo

import (
	"io/ioutil"
	"strings"
)

const (
	_          = iota
	kb float64 = 1 << (10 * iota)
	mb
	gb
)

const (
	hz  float64 = 1
	khz         = 1000 * hz
	mhz         = 1000 * khz
	ghz         = 1000 * mhz
)

func convertInt8ArrayToString(s [65]int8) string {
	b := make([]byte, len(s))
	for i, v := range s {
		b[i] = byte(v)
	}
	return string(b)
}

func convertBToKB(size uint64) uint64 {
	return size / uint64(kb)
}

func convertBToMB(size uint64) float64 {
	return float64(size) / mb
}

func convertBToGB(size uint64) float64 {
	return float64(size) / gb
}

func convertKBToB(size uint64) uint64 {
	return size * uint64(kb)
}

func convertKBToMB(size uint64) float64 {
	return float64(size) / kb
}

func convertKhzToMhz(freq uint64) uint64 {
	return freq / uint64(khz)
}

func convertKhzToGhz(freq uint64) float64 {
	return float64(freq) / mhz
}

func readSingleValueFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}
