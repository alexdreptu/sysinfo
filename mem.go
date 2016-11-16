package sysinfo

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

type Mem struct {
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

func (s *Mem) TotalMemInKB() uint64 {
	return convertBToKB(s.TotalMem)
}

func (s *Mem) TotalMemInMB() float64 {
	return convertBToMB(s.TotalMem)
}

func (s *Mem) TotalMemInGB() float64 {
	return convertBToGB(s.TotalMem)
}

func (s *Mem) TotalHighMemInKB() uint64 {
	return convertBToKB(s.TotalHighMem)
}

func (s *Mem) TotalHighMemInMB() float64 {
	return convertBToMB(s.TotalHighMem)
}

func (s *Mem) TotalHighMemInGB() float64 {
	return convertBToGB(s.TotalHighMem)
}

// GetFreeRAMInKB returns the memory not being used by the system
func (s *Mem) FreeMemInKB() uint64 {
	return convertBToKB(s.FreeMem)
}

func (s *Mem) FreeMemInMB() float64 {
	return convertBToMB(s.FreeMem)
}

func (s *Mem) FreeMemInGB() float64 {
	return convertBToGB(s.FreeMem)
}

func (s *Mem) FreeHighMemInKB() uint64 {
	return convertBToKB(s.FreeHighMem)
}

func (s *Mem) FreeHighMemInMB() float64 {
	return convertBToMB(s.FreeHighMem)
}

func (s *Mem) FreeHighMemInGB() float64 {
	return convertBToGB(s.FreeHighMem)
}

// GetAvailRAMInKB returns available memory
// that can be immediately used by processes
func (s *Mem) AvailMemInKB() uint64 {
	return convertBToKB(s.AvailMem)
}

func (s *Mem) AvailMemInMB() float64 {
	return convertBToMB(s.AvailMem)
}

func (s *Mem) AvailMemInGB() float64 {
	return convertBToGB(s.AvailMem)
}

func (s *Mem) UsedMemInKB() uint64 {
	return convertBToKB(s.UsedMem)
}

func (s *Mem) UsedMemInMB() float64 {
	return convertBToMB(s.UsedMem)
}

func (s *Mem) UsedMemInGB() float64 {
	return convertBToGB(s.UsedMem)
}

func (s *Mem) TotalUsedMemInKB() uint64 {
	return convertBToKB(s.TotalUsedMem)
}

func (s *Mem) TotalUsedMemInMB() float64 {
	return convertBToMB(s.TotalUsedMem)
}

func (s *Mem) TotalUsedMemInGB() float64 {
	return convertBToGB(s.TotalUsedMem)
}

func (s *Mem) TotalUsedInKB() uint64 {
	return convertBToKB(s.TotalUsed)
}

func (s *Mem) TotalUsedInMB() float64 {
	return convertBToMB(s.TotalUsed)
}

func (s *Mem) TotalUsedInGB() float64 {
	return convertBToGB(s.TotalUsed)
}

func (s *Mem) BufferMemInKB() uint64 {
	return convertBToKB(s.BufferMem)
}

func (s *Mem) BufferMemInMB() float64 {
	return convertBToMB(s.BufferMem)
}

func (s *Mem) BufferMemInGB() float64 {
	return convertBToGB(s.BufferMem)
}

func (s *Mem) SharedMemInKB() uint64 {
	return convertBToKB(s.SharedMem)
}

func (s *Mem) SharedMemInMB() float64 {
	return convertBToMB(s.SharedMem)
}

func (s *Mem) SharedMemInGB() float64 {
	return convertBToGB(s.SharedMem)
}

func (s *Mem) TotalSwapInKB() uint64 {
	return convertBToKB(s.TotalSwap)
}

func (s *Mem) TotalSwapInMB() float64 {
	return convertBToMB(s.TotalSwap)
}

func (s *Mem) TotalSwapInGB() float64 {
	return convertBToGB(s.TotalSwap)
}

func (s *Mem) FreeSwapInKB() uint64 {
	return convertBToKB(s.FreeSwap)
}

func (s *Mem) FreeSwapInMB() float64 {
	return convertBToMB(s.FreeSwap)
}

func (s *Mem) FreeSwapInGB() float64 {
	return convertBToGB(s.FreeSwap)
}

func (s *Mem) UsedSwapInKB() uint64 {
	return convertBToKB(s.UsedSwap)
}

func (s *Mem) UsedSwapInMB() float64 {
	return convertBToMB(s.UsedSwap)
}

func (s *Mem) UsedSwapInGB() float64 {
	return convertBToGB(s.UsedSwap)
}

func (s *Mem) Get() error {
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

// read /proc/meminfo and returns an anonymous struct
// with available and cached memory in bytes
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
