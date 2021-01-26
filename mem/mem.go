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

const memInfoPath = "/proc/meminfo"

func (m *Mem) TotalMemInKibibytes() float64 {
	return (convert.Size(m.TotalMem) * convert.Byte).ToKibibytes()
}

func (m *Mem) TotalMemInMebibytes() float64 {
	return (convert.Size(m.TotalMem) * convert.Byte).ToMebibytes()
}

func (m *Mem) TotalMemInGibibytes() float64 {
	return (convert.Size(m.TotalMem) * convert.Byte).ToGibibytes()
}

func (m *Mem) TotalHighMemInKibibytes() float64 {
	return (convert.Size(m.TotalHighMem) * convert.Byte).ToKibibytes()
}

func (m *Mem) TotalHighMemInMebibytes() float64 {
	return (convert.Size(m.TotalHighMem) * convert.Byte).ToMebibytes()
}

func (m *Mem) TotalHighMemInGibibytes() float64 {
	return (convert.Size(m.TotalHighMem) * convert.Byte).ToGibibytes()
}

func (m *Mem) FreeMemInKibibytes() float64 {
	return (convert.Size(m.FreeMem) * convert.Byte).ToKibibytes()
}

func (m *Mem) FreeMemInMebibytes() float64 {
	return (convert.Size(m.FreeMem) * convert.Byte).ToMebibytes()
}

func (m *Mem) FreeMemInGibibytes() float64 {
	return (convert.Size(m.FreeMem) * convert.Byte).ToGibibytes()
}

func (m *Mem) FreeHighMemInKibibytes() float64 {
	return (convert.Size(m.FreeHighMem) * convert.Byte).ToKibibytes()
}

func (m *Mem) FreeHighMemInMebibytes() float64 {
	return (convert.Size(m.FreeHighMem) * convert.Byte).ToMebibytes()
}

func (m *Mem) FreeHighMemInGibibytes() float64 {
	return (convert.Size(m.FreeHighMem) * convert.Byte).ToGibibytes()
}

func (m *Mem) AvailMemInKibibytes() float64 {
	return (convert.Size(m.AvailMem) * convert.Byte).ToKibibytes()
}

func (m *Mem) AvailMemInMebibytes() float64 {
	return (convert.Size(m.AvailMem) * convert.Byte).ToMebibytes()
}

func (m *Mem) AvailMemInGibibytes() float64 {
	return (convert.Size(m.AvailMem) * convert.Byte).ToGibibytes()
}

func (m *Mem) CachedMemInKibibytes() float64 {
	return (convert.Size(m.CachedMem) * convert.Byte).ToKibibytes()
}

func (m *Mem) CachedMemInMebibytes() float64 {
	return (convert.Size(m.CachedMem) * convert.Byte).ToMebibytes()
}

func (m *Mem) CachedMemInGibibytes() float64 {
	return (convert.Size(m.CachedMem) * convert.Byte).ToGibibytes()
}

func (m *Mem) UsedMemInKibibytes() float64 {
	return (convert.Size(m.UsedMem) * convert.Byte).ToKibibytes()
}

func (m *Mem) UsedMemInMebibytes() float64 {
	return (convert.Size(m.UsedMem) * convert.Byte).ToMebibytes()
}

func (m *Mem) UsedMemInGibibytes() float64 {
	return (convert.Size(m.UsedMem) * convert.Byte).ToGibibytes()
}

func (m *Mem) TotalUsedMemInKibibytes() float64 {
	return (convert.Size(m.TotalUsedMem) * convert.Byte).ToKibibytes()
}

func (m *Mem) TotalUsedMemInMebibytes() float64 {
	return (convert.Size(m.TotalUsedMem) * convert.Byte).ToMebibytes()
}

func (m *Mem) TotalUsedMemInGibibytes() float64 {
	return (convert.Size(m.TotalUsedMem) * convert.Byte).ToGibibytes()
}

func (m *Mem) TotalUsedInKibibytes() float64 {
	return (convert.Size(m.TotalUsed) * convert.Byte).ToKibibytes()
}

func (m *Mem) TotalUsedInMebibytes() float64 {
	return (convert.Size(m.TotalUsed) * convert.Byte).ToMebibytes()
}

func (m *Mem) TotalUsedInGibibytes() float64 {
	return (convert.Size(m.TotalUsed) * convert.Byte).ToGibibytes()
}

func (m *Mem) BufferMemInKibibytes() float64 {
	return (convert.Size(m.BufferMem) * convert.Byte).ToKibibytes()
}

func (m *Mem) BufferMemInMebibytes() float64 {
	return (convert.Size(m.BufferMem) * convert.Byte).ToMebibytes()
}

func (m *Mem) BufferMemInGibibytes() float64 {
	return (convert.Size(m.BufferMem) * convert.Byte).ToGibibytes()
}

func (m *Mem) SharedMemInKibibytes() float64 {
	return (convert.Size(m.SharedMem) * convert.Byte).ToKibibytes()
}

func (m *Mem) SharedMemInMebibytes() float64 {
	return (convert.Size(m.SharedMem) * convert.Byte).ToMebibytes()
}

func (m *Mem) SharedMemInGibibytes() float64 {
	return (convert.Size(m.SharedMem) * convert.Byte).ToGibibytes()
}

func (m *Mem) TotalSwapInKibibytes() float64 {
	return (convert.Size(m.TotalSwap) * convert.Byte).ToKibibytes()
}

func (m *Mem) TotalSwapInMebibytes() float64 {
	return (convert.Size(m.TotalSwap) * convert.Byte).ToMebibytes()
}

func (m *Mem) TotalSwapInGibibytes() float64 {
	return (convert.Size(m.TotalSwap) * convert.Byte).ToGibibytes()
}

func (m *Mem) FreeSwapInKibibytes() float64 {
	return (convert.Size(m.FreeSwap) * convert.Byte).ToKibibytes()
}

func (m *Mem) FreeSwapInMebibytes() float64 {
	return (convert.Size(m.FreeSwap) * convert.Byte).ToMebibytes()
}

func (m *Mem) FreeSwapInGibibytes() float64 {
	return (convert.Size(m.FreeSwap) * convert.Byte).ToGibibytes()
}

func (m *Mem) UsedSwapInKibibytes() float64 {
	return (convert.Size(m.UsedSwap) * convert.Byte).ToKibibytes()
}

func (m *Mem) UsedSwapInMebibytes() float64 {
	return (convert.Size(m.UsedSwap) * convert.Byte).ToMebibytes()
}

func (m *Mem) UsedSwapInGibibytes() float64 {
	return (convert.Size(m.UsedSwap) * convert.Byte).ToGibibytes()
}

// Fetch updates the Mem struct woth new values
func (m *Mem) Fetch() error {
	if m.F == nil {
		m.F = unix.Sysinfo
	}

	si := unix.Sysinfo_t{}
	if err := m.F(&si); err != nil {
		return err
	}

	if m.AvailMem == 0 && m.CachedMem == 0 {
		if err := m.readMemInfo(); err != nil {
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

func (m *Mem) readMemInfo() error {
	file, err := os.Open(memInfoPath)
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
			m.AvailMem = uint64((convert.Size(n) * convert.Kibibyte).ToBytes())

		case "Cached":
			n, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			m.CachedMem = uint64((convert.Size(n) * convert.Kibibyte).ToBytes())
		}
	}

	return scanner.Err()
}

func New() (*Mem, error) {
	mem := &Mem{}

	if err := mem.Fetch(); err != nil {
		return &Mem{}, err
	}

	return mem, nil
}
