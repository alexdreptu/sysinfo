package mem

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/alexdreptu/sysinfo/convert"
	"golang.org/x/sys/unix"
)

type fetchFunc func(info *unix.Sysinfo_t) error

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
	F            fetchFunc
}

func (m *Mem) TotalMemInKibibytes() float64 {
	return (convert.Unit(m.TotalMem) * convert.Byte).Kibibytes()
}

func (m *Mem) TotalMemInMebibytes() float64 {
	return (convert.Unit(m.TotalMem) * convert.Byte).Mebibytes()
}

func (m *Mem) TotalMemInGibibytes() float64 {
	return (convert.Unit(m.TotalMem) * convert.Byte).Gibibytes()
}

func (m *Mem) TotalHighMemInKibibytes() float64 {
	return (convert.Unit(m.TotalHighMem) * convert.Byte).Kibibytes()
}

func (m *Mem) TotalHighMemInMebibytes() float64 {
	return (convert.Unit(m.TotalHighMem) * convert.Byte).Mebibytes()
}

func (m *Mem) TotalHighMemInGibibytes() float64 {
	return (convert.Unit(m.TotalHighMem) * convert.Byte).Gibibytes()
}

func (m *Mem) FreeMemInKibibytes() float64 {
	return (convert.Unit(m.FreeMem) * convert.Byte).Kibibytes()
}

func (m *Mem) FreeMemInMebibytes() float64 {
	return (convert.Unit(m.FreeMem) * convert.Byte).Mebibytes()
}

func (m *Mem) FreeMemInGibibytes() float64 {
	return (convert.Unit(m.FreeMem) * convert.Byte).Gibibytes()
}

func (m *Mem) FreeHighMemInKibibytes() float64 {
	return (convert.Unit(m.FreeHighMem) * convert.Byte).Kibibytes()
}

func (m *Mem) FreeHighMemInMebibytes() float64 {
	return (convert.Unit(m.FreeHighMem) * convert.Byte).Mebibytes()
}

func (m *Mem) FreeHighMemInGibibytes() float64 {
	return (convert.Unit(m.FreeHighMem) * convert.Byte).Gibibytes()
}

func (m *Mem) AvailMemInKibibytes() float64 {
	return (convert.Unit(m.AvailMem) * convert.Byte).Kibibytes()
}

func (m *Mem) AvailMemInMebibytes() float64 {
	return (convert.Unit(m.AvailMem) * convert.Byte).Mebibytes()
}

func (m *Mem) AvailMemInGibibytes() float64 {
	return (convert.Unit(m.AvailMem) * convert.Byte).Gibibytes()
}

func (m *Mem) CachedMemInKibibytes() float64 {
	return (convert.Unit(m.CachedMem) * convert.Byte).Kibibytes()
}

func (m *Mem) CachedMemInMebibytes() float64 {
	return (convert.Unit(m.CachedMem) * convert.Byte).Mebibytes()
}

func (m *Mem) CachedMemInGibibytes() float64 {
	return (convert.Unit(m.CachedMem) * convert.Byte).Gibibytes()
}

func (m *Mem) UsedMemInKibibytes() float64 {
	return (convert.Unit(m.UsedMem) * convert.Byte).Kibibytes()
}

func (m *Mem) UsedMemInMebibytes() float64 {
	return (convert.Unit(m.UsedMem) * convert.Byte).Mebibytes()
}

func (m *Mem) UsedMemInGibibytes() float64 {
	return (convert.Unit(m.UsedMem) * convert.Byte).Gibibytes()
}

func (m *Mem) TotalUsedMemInKibibytes() float64 {
	return (convert.Unit(m.TotalUsedMem) * convert.Byte).Kibibytes()
}

func (m *Mem) TotalUsedMemInMebibytes() float64 {
	return (convert.Unit(m.TotalUsedMem) * convert.Byte).Mebibytes()
}

func (m *Mem) TotalUsedMemInGibibytes() float64 {
	return (convert.Unit(m.TotalUsedMem) * convert.Byte).Gibibytes()
}

func (m *Mem) TotalUsedInKibibytes() float64 {
	return (convert.Unit(m.TotalUsed) * convert.Byte).Kibibytes()
}

func (m *Mem) TotalUsedInMebibytes() float64 {
	return (convert.Unit(m.TotalUsed) * convert.Byte).Mebibytes()
}

func (m *Mem) TotalUsedInGibibytes() float64 {
	return (convert.Unit(m.TotalUsed) * convert.Byte).Gibibytes()
}

func (m *Mem) BufferMemInKibibytes() float64 {
	return (convert.Unit(m.BufferMem) * convert.Byte).Kibibytes()
}

func (m *Mem) BufferMemInMebibytes() float64 {
	return (convert.Unit(m.BufferMem) * convert.Byte).Mebibytes()
}

func (m *Mem) BufferMemInGibibytes() float64 {
	return (convert.Unit(m.BufferMem) * convert.Byte).Gibibytes()
}

func (m *Mem) SharedMemInKibibytes() float64 {
	return (convert.Unit(m.SharedMem) * convert.Byte).Kibibytes()
}

func (m *Mem) SharedMemInMebibytes() float64 {
	return (convert.Unit(m.SharedMem) * convert.Byte).Mebibytes()
}

func (m *Mem) SharedMemInGibibytes() float64 {
	return (convert.Unit(m.SharedMem) * convert.Byte).Gibibytes()
}

func (m *Mem) TotalSwapInKibibytes() float64 {
	return (convert.Unit(m.TotalSwap) * convert.Byte).Kibibytes()
}

func (m *Mem) TotalSwapInMebibytes() float64 {
	return (convert.Unit(m.TotalSwap) * convert.Byte).Mebibytes()
}

func (m *Mem) TotalSwapInGibibytes() float64 {
	return (convert.Unit(m.TotalSwap) * convert.Byte).Gibibytes()
}

func (m *Mem) FreeSwapInKibibytes() float64 {
	return (convert.Unit(m.FreeSwap) * convert.Byte).Kibibytes()
}

func (m *Mem) FreeSwapInMebibytes() float64 {
	return (convert.Unit(m.FreeSwap) * convert.Byte).Mebibytes()
}

func (m *Mem) FreeSwapInGibibytes() float64 {
	return (convert.Unit(m.FreeSwap) * convert.Byte).Gibibytes()
}

func (m *Mem) UsedSwapInKibibytes() float64 {
	return (convert.Unit(m.UsedSwap) * convert.Byte).Kibibytes()
}

func (m *Mem) UsedSwapInMebibytes() float64 {
	return (convert.Unit(m.UsedSwap) * convert.Byte).Mebibytes()
}

func (m *Mem) UsedSwapInGibibytes() float64 {
	return (convert.Unit(m.UsedSwap) * convert.Byte).Gibibytes()
}

// Fetch updates the Mem struct woth new values
func (m *Mem) Fetch() error {
	si := unix.Sysinfo_t{}

	if m.F == nil {
		m.F = unix.Sysinfo
	}

	if err := m.F(&si); err != nil {
		return err
	}

	if m.AvailMem == 0 && m.CachedMem == 0 {
		if err := m.readMeminfo(); err != nil {
			return err
		}
	}

	m.Procs = si.Procs
	m.TotalMem = si.Totalram
	m.TotalHighMem = si.Totalhigh
	m.FreeMem = si.Freeram
	m.FreeHighMem = si.Freehigh
	m.BufferMem = si.Bufferram
	m.SharedMem = si.Sharedram
	m.TotalSwap = si.Totalswap
	m.FreeSwap = si.Freeswap

	m.UsedMem = m.TotalMem - m.FreeMem - m.BufferMem - m.CachedMem
	m.TotalUsedMem = m.TotalMem - m.FreeMem
	m.TotalUsed = m.TotalMem - m.FreeMem + m.TotalSwap - m.FreeSwap
	m.UsedSwap = m.TotalSwap - m.FreeSwap

	return nil
}

func (m *Mem) readMeminfo() error {
	file, err := os.Open("/proc/meminfo")
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
		value := strings.Fields(fields[1])[0]

		switch key {
		case "MemAvailable":
			n, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			m.AvailMem = uint64((convert.Unit(n) * convert.Kibibyte).Bytes())

		case "Cached":
			n, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			m.CachedMem = uint64((convert.Unit(n) * convert.Kibibyte).Bytes())
		}
	}

	return scanner.Err()
}

func New() *Mem {
	mem := &Mem{}
	mem.Fetch()
	return mem
}
