package sysinfo

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

type MemInfo struct {
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

func (s *MemInfo) TotalMemInKB() uint64 {
	return convertBToKB(s.TotalMem)
}

func (s *MemInfo) TotalMemInMB() float64 {
	return convertBToMB(s.TotalMem)
}

func (s *MemInfo) TotalMemInGB() float64 {
	return convertBToGB(s.TotalMem)
}

func (s *MemInfo) TotalHighMemInKB() uint64 {
	return convertBToKB(s.TotalHighMem)
}

func (s *MemInfo) TotalHighMemInMB() float64 {
	return convertBToMB(s.TotalHighMem)
}

func (s *MemInfo) TotalHighMemInGB() float64 {
	return convertBToGB(s.TotalHighMem)
}

// GetFreeRAMInKB returns the memory not being used by the system
func (s *MemInfo) FreeMemInKB() uint64 {
	return convertBToKB(s.FreeMem)
}

func (s *MemInfo) FreeMemInMB() float64 {
	return convertBToMB(s.FreeMem)
}

func (s *MemInfo) FreeMemInGB() float64 {
	return convertBToGB(s.FreeMem)
}

func (s *MemInfo) FreeHighMemInKB() uint64 {
	return convertBToKB(s.FreeHighMem)
}

func (s *MemInfo) FreeHighMemInMB() float64 {
	return convertBToMB(s.FreeHighMem)
}

func (s *MemInfo) FreeHighMemInGB() float64 {
	return convertBToGB(s.FreeHighMem)
}

// GetAvailRAMInKB returns available memory
// that can be immediately used by processes
func (s *MemInfo) AvailMemInKB() uint64 {
	return convertBToKB(s.AvailMem)
}

func (s *MemInfo) AvailMemInMB() float64 {
	return convertBToMB(s.AvailMem)
}

func (s *MemInfo) AvailMemInGB() float64 {
	return convertBToGB(s.AvailMem)
}

func (s *MemInfo) UsedMemInKB() uint64 {
	return convertBToKB(s.UsedMem)
}

func (s *MemInfo) UsedMemInMB() float64 {
	return convertBToMB(s.UsedMem)
}

func (s *MemInfo) UsedMemInGB() float64 {
	return convertBToGB(s.UsedMem)
}

func (s *MemInfo) TotalUsedMemInKB() uint64 {
	return convertBToKB(s.TotalUsedMem)
}

func (s *MemInfo) TotalUsedMemInMB() float64 {
	return convertBToMB(s.TotalUsedMem)
}

func (s *MemInfo) TotalUsedMemInGB() float64 {
	return convertBToGB(s.TotalUsedMem)
}

func (s *MemInfo) TotalUsedInKB() uint64 {
	return convertBToKB(s.TotalUsed)
}

func (s *MemInfo) TotalUsedInMB() float64 {
	return convertBToMB(s.TotalUsed)
}

func (s *MemInfo) TotalUsedInGB() float64 {
	return convertBToGB(s.TotalUsed)
}

func (s *MemInfo) BufferMemInKB() uint64 {
	return convertBToKB(s.BufferMem)
}

func (s *MemInfo) BufferMemInMB() float64 {
	return convertBToMB(s.BufferMem)
}

func (s *MemInfo) BufferMemInGB() float64 {
	return convertBToGB(s.BufferMem)
}

func (s *MemInfo) SharedMemInKB() uint64 {
	return convertBToKB(s.SharedMem)
}

func (s *MemInfo) SharedMemInMB() float64 {
	return convertBToMB(s.SharedMem)
}

func (s *MemInfo) SharedMemInGB() float64 {
	return convertBToGB(s.SharedMem)
}

func (s *MemInfo) TotalSwapInKB() uint64 {
	return convertBToKB(s.TotalSwap)
}

func (s *MemInfo) TotalSwapInMB() float64 {
	return convertBToMB(s.TotalSwap)
}

func (s *MemInfo) TotalSwapInGB() float64 {
	return convertBToGB(s.TotalSwap)
}

func (s *MemInfo) FreeSwapInKB() uint64 {
	return convertBToKB(s.FreeSwap)
}

func (s *MemInfo) FreeSwapInMB() float64 {
	return convertBToMB(s.FreeSwap)
}

func (s *MemInfo) FreeSwapInGB() float64 {
	return convertBToGB(s.FreeSwap)
}

func (s *MemInfo) UsedSwapInKB() uint64 {
	return convertBToKB(s.UsedSwap)
}

func (s *MemInfo) UsedSwapInMB() float64 {
	return convertBToMB(s.UsedSwap)
}

func (s *MemInfo) UsedSwapInGB() float64 {
	return convertBToGB(s.UsedSwap)
}

func (s *MemInfo) Get() error {
	si := unix.Sysinfo_t{}
	if err := unix.Sysinfo(&si); err != nil {
		return err
	}

	s.Procs = si.Procs
	s.TotalMem = si.Totalram
	s.TotalHighMem = si.Totalhigh
	s.FreeMem = si.Freeram
	s.FreeHighMem = si.Freehigh
	s.BufferMem = si.Bufferram
	s.SharedMem = si.Sharedram
	s.TotalSwap = si.Totalswap
	s.FreeSwap = si.Freeswap

	meminfo, err := readMeminfo()
	if err != nil {
		return err
	}

	s.AvailMem = meminfo.availMem
	s.CachedMem = meminfo.cachedMem

	s.UsedMem = s.TotalMem - s.FreeMem - s.BufferMem - s.CachedMem
	s.TotalUsedMem = s.TotalMem - s.FreeMem
	s.TotalUsed = s.TotalMem - s.FreeMem + s.TotalSwap - s.FreeSwap
	s.UsedSwap = s.TotalSwap - s.FreeSwap

	return nil
}

// read /proc/meminfo and return available and cached memory in bytes
//func readMeminfo() (uint64, uint64, error) {
func readMeminfo() (struct{ availMem, cachedMem uint64 }, error) {
	meminfo := struct{ availMem, cachedMem uint64 }{}

	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return meminfo, err
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
		value := strings.Fields(fields[1])[0]

		switch key {
		case "MemAvailable":
			n, err := strconv.Atoi(value)
			if err != nil {
				return meminfo, err
			}
			meminfo.availMem = convertKBToB(uint64(n))

		case "Cached":
			n, err := strconv.Atoi(value)
			if err != nil {
				return meminfo, err
			}
			meminfo.cachedMem = convertKBToB(uint64(n))
		}
	}

	if err := scanner.Err(); err != nil {
		return meminfo, err
	}

	return meminfo, nil
}

func NewMem() *MemInfo {
	return &MemInfo{}
}
