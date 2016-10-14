package sysinfo

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sys/unix"
)

type Sysinfo struct {
	Uptime       int64
	Procs        uint16
	TotalRAM     uint64
	TotalHighRAM uint64
	FreeRAM      uint64
	FreeHighRAM  uint64
	AvailRAM     uint64
	BufferRAM    uint64
	SharedRAM    uint64
	TotalSwap    uint64
	FreeSwap     uint64
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

func (s *Sysinfo) TotalRAMInKB() uint64 {
	return ConvertBToKB(s.TotalRAM)
}

func (s *Sysinfo) TotalRAMInMB() float64 {
	return ConvertBToMB(s.TotalRAM)
}

func (s *Sysinfo) TotalRAMInGB() float64 {
	return ConvertBToGB(s.TotalRAM)
}

func (s *Sysinfo) TotalHighRAMInKB() uint64 {
	return ConvertBToKB(s.TotalHighRAM)
}

func (s *Sysinfo) TotalHighRAMInMB() float64 {
	return ConvertBToMB(s.TotalHighRAM)
}

func (s *Sysinfo) TotalHighRAMInGB() float64 {
	return ConvertBToGB(s.TotalHighRAM)
}

// GetFreeRAMInKB returns the memory not being used by the system
func (s *Sysinfo) FreeRAMInKB() uint64 {
	return ConvertBToKB(s.FreeRAM)
}

func (s *Sysinfo) FreeRAMInMB() float64 {
	return ConvertBToMB(s.FreeRAM)
}

func (s *Sysinfo) FreeRAMInGB() float64 {
	return ConvertBToGB(s.FreeRAM)
}

func (s *Sysinfo) FreeHighRAMInKB() uint64 {
	return ConvertBToKB(s.FreeHighRAM)
}

func (s *Sysinfo) FreeHighRAMInMB() float64 {
	return ConvertBToMB(s.FreeHighRAM)
}

func (s *Sysinfo) FreeHighRAMInGB() float64 {
	return ConvertBToGB(s.FreeHighRAM)
}

// GetAvailRAMInKB returns available memory
// that can be immediately used by processes
func (s *Sysinfo) AvailRAMInKB() uint64 {
	return ConvertBToKB(s.AvailRAM)
}

func (s *Sysinfo) AvailRAMInMB() float64 {
	return ConvertBToMB(s.AvailRAM)
}

func (s *Sysinfo) AvailRAMInGB() float64 {
	return ConvertBToGB(s.AvailRAM)
}

func (s *Sysinfo) BufferRAMInKB() uint64 {
	return ConvertBToKB(s.BufferRAM)
}

func (s *Sysinfo) BufferRAMInMB() float64 {
	return ConvertBToMB(s.BufferRAM)
}

func (s *Sysinfo) BufferRAMInGB() float64 {
	return ConvertBToGB(s.BufferRAM)
}

func (s *Sysinfo) SharedRAMInKB() uint64 {
	return ConvertBToKB(s.SharedRAM)
}

func (s *Sysinfo) SharedRAMInMB() float64 {
	return ConvertBToMB(s.SharedRAM)
}

func (s *Sysinfo) SharedRAMInGB() float64 {
	return ConvertBToGB(s.SharedRAM)
}

func (s *Sysinfo) TotalSwapInKB() uint64 {
	return ConvertBToKB(s.TotalSwap)
}

func (s *Sysinfo) TotalSwapInMB() float64 {
	return ConvertBToMB(s.TotalSwap)
}

func (s *Sysinfo) TotalSwapInGB() float64 {
	return ConvertBToGB(s.TotalSwap)
}

func (s *Sysinfo) FreeSwapInKB() uint64 {
	return ConvertBToKB(s.FreeSwap)
}

func (s *Sysinfo) FreeSwapInMB() float64 {
	return ConvertBToMB(s.FreeSwap)
}

func (s *Sysinfo) FreeSwapInGB() float64 {
	return ConvertBToGB(s.FreeSwap)
}

func (s *Sysinfo) Get() error {
	sys := &unix.Sysinfo_t{}
	if err := unix.Sysinfo(sys); err != nil {
		return err
	}

	s.Uptime = sys.Uptime
	s.Procs = sys.Procs
	s.TotalRAM = sys.Totalram
	s.TotalHighRAM = sys.Totalhigh
	s.FreeRAM = sys.Freeram
	s.FreeHighRAM = sys.Freehigh
	s.BufferRAM = sys.Bufferram
	s.SharedRAM = sys.Sharedram
	s.TotalSwap = sys.Totalswap
	s.FreeSwap = sys.Freeswap

	var err error
	s.AvailRAM, err = readMeminfo()
	if err != nil {
		return err
	}

	return nil
}

// read /proc/meminfo and return available ram in bytes
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

	return ConvertKBToB(uint64(availmem)), nil
}
