package cpu_test

import (
	"testing"

	. "github.com/alexdreptu/sysinfo/cpu"
	"github.com/stretchr/testify/assert"
)

// test values
const (
	cacheSizeInKibibytes uint64  = 6144
	cacheSizeInMebibytes float64 = 6

	minFreqInKilohertz uint64  = 800000
	minFreqInMegahertz float64 = 800
	minFreqInGigahertz float64 = 0.8

	maxFreqInKilohertz uint64  = 3900000
	maxFreqInMegahertz float64 = 3900
	maxFreqInGigahertz float64 = 3.9
)

const delta = 0.01

func TestCPUInfo(t *testing.T) {
	cpu := &CPU{}
	cpu.M = true
	cpu.CacheSize = cacheSizeInKibibytes
	cpu.MinFreq = minFreqInKilohertz
	cpu.MaxFreq = maxFreqInKilohertz

	assert.NoError(t, cpu.Fetch()) // unnecessary for now

	t.Run("cache size", func(t *testing.T) {
		assert.InDelta(t, cacheSizeInMebibytes, cpu.CacheSizeInMebibytes(), delta)
	})

	t.Run("min frequency", func(t *testing.T) {
		assert.InDelta(t, minFreqInMegahertz, cpu.MinFreqInMegahertz(), delta)
		assert.InDelta(t, minFreqInGigahertz, cpu.MinFreqInGigahertz(), delta)
	})

	t.Run("max frequency", func(t *testing.T) {
		assert.InDelta(t, maxFreqInMegahertz, cpu.MaxFreqInMegahertz(), delta)
		assert.InDelta(t, maxFreqInGigahertz, cpu.MaxFreqInGigahertz(), delta)
	})

}
