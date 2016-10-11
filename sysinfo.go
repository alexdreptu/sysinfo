package sysinfo

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sys/unix"
)

type Sysinfo struct {
	unix.Sysinfo_t
	Availram uint64
}

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
)

// GetUptime returns uptime in human readable form
func (s *Sysinfo) GetUptime() time.Duration {
	return time.Duration(s.Uptime) * time.Second
}

func (s *Sysinfo) GetTotalRAMInKB() uint64 {
	return ConvertBToKB(s.Totalram)
}

func (s *Sysinfo) GetTotalRAMInMB() float64 {
	return ConvertBToMB(s.Totalram)
}

func (s *Sysinfo) GetTotalRAMInGB() float64 {
	return ConvertBToGB(s.Totalram)
}

func (s *Sysinfo) GetTotalHighRAMInKB() uint64 {
	return ConvertBToKB(s.Totalhigh)
}

func (s *Sysinfo) GetTotalHighRAMInMB() float64 {
	return ConvertBToMB(s.Totalhigh)
}

func (s *Sysinfo) GetTotalHighRAMInGB() float64 {
	return ConvertBToGB(s.Totalhigh)
}

// GetFreeRAMInKB returns the memory not being used by the system
func (s *Sysinfo) GetFreeRAMInKB() uint64 {
	return ConvertBToKB(s.Freeram)
}

func (s *Sysinfo) GetFreeRAMInMB() float64 {
	return ConvertBToMB(s.Freeram)
}

func (s *Sysinfo) GetFreeRAMInGB() float64 {
	return ConvertBToGB(s.Freeram)
}

func (s *Sysinfo) GetFreeHighRAMInKB() uint64 {
	return ConvertBToKB(s.Freehigh)
}

func (s *Sysinfo) GetFreeHighRAMInMB() float64 {
	return ConvertBToMB(s.Freehigh)
}

func (s *Sysinfo) GetFreeHighRAMInGB() float64 {
	return ConvertBToGB(s.Freehigh)
}

// GetAvailRAMInKB returns available memory
// that can be immediately used by processes
func (s *Sysinfo) GetAvailRAMInKB() uint64 {
	return ConvertBToKB(s.Availram)
}

func (s *Sysinfo) GetAvailRAMInMB() float64 {
	return ConvertBToMB(s.Availram)
}

func (s *Sysinfo) GetAvailRAMInGB() float64 {
	return ConvertBToGB(s.Availram)
}

func (s *Sysinfo) GetBufferRAMInKB() uint64 {
	return ConvertBToKB(s.Bufferram)
}

func (s *Sysinfo) GetBufferRAMInMB() float64 {
	return ConvertBToMB(s.Bufferram)
}

func (s *Sysinfo) GetBufferRAMInGB() float64 {
	return ConvertBToGB(s.Bufferram)
}

func (s *Sysinfo) GetSharedRAMInKB() uint64 {
	return ConvertBToKB(s.Sharedram)
}

func (s *Sysinfo) GetSharedRAMInMB() float64 {
	return ConvertBToMB(s.Sharedram)
}

func (s *Sysinfo) GetSharedRAMInGB() float64 {
	return ConvertBToGB(s.Sharedram)
}

func (s *Sysinfo) GetTotalSwapInKB() uint64 {
	return ConvertBToKB(s.Totalswap)
}

func (s *Sysinfo) GetTotalSwapInMB() float64 {
	return ConvertBToMB(s.Totalswap)
}

func (s *Sysinfo) GetTotalSwapInGB() float64 {
	return ConvertBToGB(s.Totalswap)
}

func (s *Sysinfo) GetFreeSwapInKB() uint64 {
	return ConvertBToKB(s.Freeswap)
}

func (s *Sysinfo) GetFreeSwapInMB() float64 {
	return ConvertBToMB(s.Freeswap)
}

func (s *Sysinfo) GetFreeSwapInGB() float64 {
	return ConvertBToGB(s.Freeswap)
}

func (s *Sysinfo) Get() error {
	var err error
	info := &unix.Sysinfo_t{}
	if err = unix.Sysinfo(info); err != nil {
		return err
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	if err = enc.Encode(info); err != nil {
		return err
	}

	if err = dec.Decode(&s); err != nil {
		return err
	}

	s.Availram, err = readMeminfo()
	if err != nil {
		return err
	}

	return nil
}

// read /proc/meminfo and return available ram
func readMeminfo() (uint64, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var content string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = scanner.Text()
		if strings.Contains(content, "MemAvailable") {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	availmem, err := strconv.Atoi(
		strings.Trim(strings.Split(content, " ")[3], " "),
	)

	if err != nil {
		return 0, err
	}

	return uint64(availmem), nil
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

func ConvertGBToB(size float64) float64 {
	return size * GB
}
