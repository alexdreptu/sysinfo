package cpu

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/alexdreptu/sysinfo/convert"
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
	M         bool // mock
}

type freqScale int

const (
	minFreq freqScale = iota
	maxFreq
)

const (
	cpuInfoPath    = "/proc/cpuinfo"
	cpuFreqPath    = "/sys/devices/system/cpu/cpu0/cpufreq"
	cpuMinFreqPath = "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_min_freq"
	cpuMaxFreqPath = "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_max_freq"
)

func (c *CPU) CacheSizeInMebibytes() float64 {
	return (convert.Unit(c.CacheSize) * convert.Kibibyte).Mebibytes()
}

func (c *CPU) MinFreqInMegahertz() float64 {
	return (convert.Unit(c.MinFreq) * convert.Kilohertz).Megahertz()
}

func (c *CPU) MinFreqInGigahertz() float64 {
	return (convert.Unit(c.MinFreq) * convert.Kilohertz).Gigahertz()
}

func (c *CPU) MaxFreqInMegahertz() float64 {
	return (convert.Unit(c.MaxFreq) * convert.Kilohertz).Megahertz()
}

func (c *CPU) MaxFreqInGigahertz() float64 {
	return (convert.Unit(c.MaxFreq) * convert.Kilohertz).Gigahertz()
}

// Fetch updates the CPU struct woth new values
func (c *CPU) Fetch() error {
	if !c.M {
		if _, err := os.Stat(cpuFreqPath); os.IsNotExist(err) {
			return nil
		}

		if err := c.readCpuinfo(); err != nil {
			return err
		}

		if err := c.readCPUFreq(minFreq); err != nil {
			return err
		}

		if err := c.readCPUFreq(maxFreq); err != nil {
			return err
		}
	}

	return nil
}

func (c *CPU) readCPUFreq(scale freqScale) error {
	if scale == minFreq {
		value, err := readSingleValueFile(cpuMinFreqPath)
		if err != nil {
			return err
		}

		c.MinFreq, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
	} else if scale == maxFreq {
		value, err := readSingleValueFile(cpuMaxFreqPath)
		if err != nil {
			return err
		}

		c.MaxFreq, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CPU) readCpuinfo() error {
	file, err := os.Open(cpuInfoPath)
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
				return err
			}

		case "flags":
			if len(c.Flags) > 0 {
				continue
			}
			c.Flags = strings.Fields(value)
		}
	}

	return scanner.Err()
}

func readSingleValueFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}

func New() (*CPU, error) {
	cpu := &CPU{}
	if err := cpu.Fetch(); err != nil {
		return cpu, err
	}

	return cpu, nil
}
