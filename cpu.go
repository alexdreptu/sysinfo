package sysinfo

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	Name      string
	Model     uint64
	Vendor    string
	Family    uint64
	Stepping  uint64
	CacheSize uint64
	MinFreq   uint64
	MaxFreq   uint64
	Cores     uint64
	Flags     []string
}

const (
	cpuFreqPath    = "/sys/devices/system/cpu/cpu0/cpufreq"
	cpuMinFreqPath = "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_min_freq"
	cpuMaxFreqPath = "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_max_freq"
)

func (c *CPU) CacheSizeInMB() float64 {
	return convertKBToMB(c.CacheSize)
}

func (c *CPU) MinFreqInMhz() uint64 {
	return convertKhzToMhz(c.MinFreq)
}

func (c *CPU) MaxFreqInMhz() uint64 {
	return convertKhzToMhz(c.MaxFreq)
}

func (c *CPU) MinFreqInGhz() float64 {
	return convertKhzToGhz(c.MinFreq)
}

func (c *CPU) MaxFreqInGhz() float64 {
	return convertKhzToGhz(c.MaxFreq)
}

func (c *CPU) Get() error {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		fields := strings.Split(line, ":")
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])

		switch key {
		case "vendor_id":
			if c.Vendor != "" {
				continue
			}
			c.Vendor = value

		case "cpu family":
			if c.Family != 0 {
				continue
			}
			c.Family, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}

		case "model":
			if c.Model != 0 {
				continue
			}
			c.Model, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}

		case "model name":
			if c.Name != "" {
				continue
			}
			c.Name = value

		case "stepping":
			if c.Stepping != 0 {
				continue
			}
			c.Stepping, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}

		case "cache size":
			if c.CacheSize != 0 {
				continue
			}
			c.CacheSize, err = strconv.ParseUint(strings.Fields(value)[0], 10, 64)
			if err != nil {
				return err
			}

		case "cpu cores":
			if c.Cores != 0 {
				continue
			}
			c.Cores, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				return nil
			}

		case "flags":
			if len(c.Flags) > 0 {
				continue
			}
			c.Flags = strings.Fields(value)
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	if _, err := os.Stat(cpuFreqPath); os.IsNotExist(err) {
		return nil
	}

	minFreq, err := readSingleValueFile(cpuMinFreqPath)
	if err != nil {
		return err
	}

	c.MinFreq, err = strconv.ParseUint(minFreq, 10, 64)
	if err != nil {
		return err
	}

	maxFreq, err := readSingleValueFile(cpuMaxFreqPath)
	if err != nil {
		return err
	}

	c.MaxFreq, err = strconv.ParseUint(maxFreq, 10, 64)
	if err != nil {
		return err
	}

	return nil
}
