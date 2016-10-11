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
	return s.Totalram / uint64(KB)
}

func (s *Sysinfo) GetTotalRAMInMB() uint64 {
	return s.Totalram / uint64(MB)
}

func (s *Sysinfo) GetTotalRAMInGB() float64 {
	return float64(s.Totalram) / GB
}

func (s *Sysinfo) GetTotalHighRAMInKB() uint64 {
	return s.Totalhigh / uint64(KB)
}

func (s *Sysinfo) GetTotalHighRAMInMB() uint64 {
	return s.Totalhigh / uint64(MB)
}

func (s *Sysinfo) GetTotalHighRAMInGB() float64 {
	return float64(s.Totalhigh) / KB
}

// GetFreeRAMInKB returns the memory not being used by the system
func (s *Sysinfo) GetFreeRAMInKB() uint64 {
	return s.Freeram / uint64(KB)
}

func (s *Sysinfo) GetFreeRAMInMB() uint64 {
	return s.Freeram / uint64(MB)
}

func (s *Sysinfo) GetFreeRAMInGB() float64 {
	return float64(s.Freeram) / GB
}

func (s *Sysinfo) GetFreeHighRAMInKB() uint64 {
	return s.Freehigh / uint64(KB)
}

func (s *Sysinfo) GetFreeHighRAMInMB() uint64 {
	return s.Freehigh / uint64(MB)
}

func (s *Sysinfo) GetFreeHighRAMInGB() float64 {
	return float64(s.Freehigh) / GB
}

// GetAvailRAMInKB returns available memory
// that can be immediately used by processes
func (s *Sysinfo) GetAvailRAMInKB() uint64 {
	return s.Availram / uint64(KB)
}

func (s *Sysinfo) GetAvailRAMInMB() uint64 {
	return s.Availram / uint64(MB)
}

func (s *Sysinfo) GetAvailRAMInGB() float64 {
	return float64(s.Availram) / GB
}

func (s *Sysinfo) GetBufferRAMInKB() uint64 {
	return s.Bufferram / uint64(KB)
}

func (s *Sysinfo) GetBufferRAMInMB() uint64 {
	return s.Bufferram / uint64(MB)
}

func (s *Sysinfo) GetBufferRAMInGB() float64 {
	return float64(s.Bufferram) / GB
}

func (s *Sysinfo) GetSharedRAMInKB() uint64 {
	return s.Sharedram / uint64(KB)
}

func (s *Sysinfo) GetSharedRAMInMB() uint64 {
	return s.Sharedram / uint64(MB)
}

func (s *Sysinfo) GetSharedRAMInGB() float64 {
	return float64(s.Sharedram) / GB
}

func (s *Sysinfo) GetTotalSwapInKB() uint64 {
	return s.Totalswap / uint64(KB)
}

func (s *Sysinfo) GetTotalSwapInMB() uint64 {
	return s.Totalswap / uint64(MB)
}

func (s *Sysinfo) GetTotalSwapInGB() float64 {
	return float64(s.Totalswap) / GB
}

func (s *Sysinfo) GetFreeSwapInKB() uint64 {
	return s.Freeswap / uint64(KB)
}

func (s *Sysinfo) GetFreeSwapInMB() uint64 {
	return s.Freeswap / uint64(MB)
}

func (s *Sysinfo) GetFreeSwapInGB() float64 {
	return float64(s.Freeswap) / GB
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
