package sysinfo_test

import (
	"testing"

	"github.com/alexdreptu/sysinfo"
)

func TestCPUInfo(t *testing.T) {
	cpu := sysinfo.CPU{}
	if err := cpu.Get(); err != nil {
		t.Fatal(err)
	}

	t.Logf("Name: %s\n", cpu.Name)
	t.Logf("Model: %d\n", cpu.Model)
	t.Logf("Vendor: %s\n", cpu.Vendor)
	t.Logf("Family: %d\n", cpu.Family)
	t.Logf("Stepping: %d\n", cpu.Stepping)
	t.Logf("Cores: %d\n", cpu.Cores)
	t.Logf("Cache Size: %d KB\n", cpu.CacheSize)
	t.Logf("Cache Size: %.1f MB\n", cpu.CacheSizeInMB())
	t.Logf("Min Frequency: %d Khz\n", cpu.MinFreq)
	t.Logf("Max Frequency: %d Khz\n", cpu.MaxFreq)
	t.Logf("Min Frequency: %d Mhz\n", cpu.MinFreqInMhz())
	t.Logf("Max Frequency: %d Mhz\n", cpu.MaxFreqInMhz())
	t.Logf("Min Frequency: %.1f Ghz\n", cpu.MinFreqInGhz())
	t.Logf("Max Frequency: %.1f Ghz\n", cpu.MaxFreqInGhz())
	t.Logf("Flags %s\n", cpu.Flags)
}
