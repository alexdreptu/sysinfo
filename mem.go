package sysinfo

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sys/unix"
)

type MemInfo struct {
	Uptime       int64
	Procs        uint16
	TotalMem     uint64
	TotalHighMem uint64
	FreeMem      uint64
	FreeHighMem  uint64
	AvailMem     uint64
	UsedMem      uint64
	TotalUsedMem uint64
	BufferMem    uint64
	CachedMem    uint64
	SharedMem    uint64
	TotalSwap    uint64
	FreeSwap     uint64
	UsedSwap     uint64
	TotalUsed    uint64
}

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
)

// GetUptime returns uptime in human readable form
func (s *MemInfo) GetUptime() time.Duration {
	return time.Duration(s.Uptime) * time.Second
}

func (s *MemInfo) TotalMemInKB() uint64 {
	return ConvertBToKB(s.TotalMem)
}

func (s *MemInfo) TotalMemInMB() float64 {
	return ConvertBToMB(s.TotalMem)
}

func (s *MemInfo) TotalMemInGB() float64 {
	return ConvertBToGB(s.TotalMem)
}

func (s *MemInfo) TotalHighMemInKB() uint64 {
	return ConvertBToKB(s.TotalHighMem)
}

func (s *MemInfo) TotalHighMemInMB() float64 {
	return ConvertBToMB(s.TotalHighMem)
}

func (s *MemInfo) TotalHighMemInGB() float64 {
	return ConvertBToGB(s.TotalHighMem)
}

// GetFreeRAMInKB returns the memory not being used by the system
func (s *MemInfo) FreeMemInKB() uint64 {
	return ConvertBToKB(s.FreeMem)
}

func (s *MemInfo) FreeMemInMB() float64 {
	return ConvertBToMB(s.FreeMem)
}

func (s *MemInfo) FreeMemInGB() float64 {
	return ConvertBToGB(s.FreeMem)
}

func (s *MemInfo) FreeHighMemInKB() uint64 {
	return ConvertBToKB(s.FreeHighMem)
}

func (s *MemInfo) FreeHighMemInMB() float64 {
	return ConvertBToMB(s.FreeHighMem)
}

func (s *MemInfo) FreeHighMemInGB() float64 {
	return ConvertBToGB(s.FreeHighMem)
}

// GetAvailRAMInKB returns available memory
// that can be immediately used by processes
func (s *MemInfo) AvailMemInKB() uint64 {
	return ConvertBToKB(s.AvailMem)
}

func (s *MemInfo) AvailMemInMB() float64 {
	return ConvertBToMB(s.AvailMem)
}

func (s *MemInfo) AvailMemInGB() float64 {
	return ConvertBToGB(s.AvailMem)
}

func (s *MemInfo) UsedMemInKB() uint64 {
	return ConvertBToKB(s.UsedMem)
}

func (s *MemInfo) UsedMemInMB() float64 {
	return ConvertBToMB(s.UsedMem)
}

func (s *MemInfo) UsedMemInGB() float64 {
	return ConvertBToGB(s.UsedMem)
}

func (s *MemInfo) TotalUsedMemInKB() uint64 {
	return ConvertBToKB(s.TotalUsedMem)
}

func (s *MemInfo) TotalUsedMemInMB() float64 {
	return ConvertBToMB(s.TotalUsedMem)
}

func (s *MemInfo) TotalUsedMemInGB() float64 {
	return ConvertBToGB(s.TotalUsedMem)
}

func (s *MemInfo) TotalUsedInKB() uint64 {
	return ConvertBToKB(s.TotalUsed)
}

func (s *MemInfo) TotalUsedInMB() float64 {
	return ConvertBToMB(s.TotalUsed)
}

func (s *MemInfo) TotalUsedInGB() float64 {
	return ConvertBToGB(s.TotalUsed)
}

func (s *MemInfo) BufferMemInKB() uint64 {
	return ConvertBToKB(s.BufferMem)
}

func (s *MemInfo) BufferMemInMB() float64 {
	return ConvertBToMB(s.BufferMem)
}

func (s *MemInfo) BufferMemInGB() float64 {
	return ConvertBToGB(s.BufferMem)
}

func (s *MemInfo) SharedMemInKB() uint64 {
	return ConvertBToKB(s.SharedMem)
}

func (s *MemInfo) SharedMemInMB() float64 {
	return ConvertBToMB(s.SharedMem)
}

func (s *MemInfo) SharedMemInGB() float64 {
	return ConvertBToGB(s.SharedMem)
}

func (s *MemInfo) TotalSwapInKB() uint64 {
	return ConvertBToKB(s.TotalSwap)
}

func (s *MemInfo) TotalSwapInMB() float64 {
	return ConvertBToMB(s.TotalSwap)
}

func (s *MemInfo) TotalSwapInGB() float64 {
	return ConvertBToGB(s.TotalSwap)
}

func (s *MemInfo) FreeSwapInKB() uint64 {
	return ConvertBToKB(s.FreeSwap)
}

func (s *MemInfo) FreeSwapInMB() float64 {
	return ConvertBToMB(s.FreeSwap)
}

func (s *MemInfo) FreeSwapInGB() float64 {
	return ConvertBToGB(s.FreeSwap)
}

func (s *MemInfo) UsedSwapInKB() uint64 {
	return ConvertBToKB(s.UsedSwap)
}

func (s *MemInfo) UsedSwapInMB() float64 {
	return ConvertBToMB(s.UsedSwap)
}

func (s *MemInfo) UsedSwapInGB() float64 {
	return ConvertBToGB(s.UsedSwap)
}

func (s *MemInfo) Get() error {
	si := unix.Sysinfo_t{}
	if err := unix.Sysinfo(&si); err != nil {
		return err
	}

	s.Uptime = si.Uptime
	s.Procs = si.Procs
	s.TotalMem = si.Totalram
	s.TotalHighMem = si.Totalhigh
	s.FreeMem = si.Freeram
	s.FreeHighMem = si.Freehigh
	s.BufferMem = si.Bufferram
	s.SharedMem = si.Sharedram
	s.TotalSwap = si.Totalswap
	s.FreeSwap = si.Freeswap

	var err error
	s.AvailMem, s.CachedMem, err = readMeminfo()
	if err != nil {
		return err
	}

	s.UsedMem = s.TotalMem - s.FreeMem - s.BufferMem - s.CachedMem
	s.TotalUsedMem = s.TotalMem - s.FreeMem
	s.TotalUsed = s.TotalMem - s.FreeMem + s.TotalSwap - s.FreeSwap
	s.UsedSwap = s.TotalSwap - s.FreeSwap

	return nil
}

// read /proc/meminfo and return available and cached memory in bytes
func readMeminfo() (uint64, uint64, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	var content string
	var availMem uint64
	var cachedMem uint64
	var availStatus bool
	var cachedStatus bool
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content = scanner.Text()
		if strings.Contains(content, "MemAvailable") {
			m, err := strconv.Atoi(strings.Fields(content)[1])
			if err != nil {
				return 0, 0, err
			}
			availMem = ConvertKBToB(uint64(m))
			availStatus = true
		}

		if strings.Contains(content, "Cached") {
			c, err := strconv.Atoi(strings.Fields(content)[1])
			if err != nil {
				return 0, 0, err
			}
			cachedMem = ConvertKBToB(uint64(c))
			cachedStatus = true
		}

		if availStatus && cachedStatus {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}

	return availMem, cachedMem, nil
}

func NewMem() *MemInfo { return &MemInfo{} }
