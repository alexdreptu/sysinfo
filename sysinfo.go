package sysinfo

import (
	"bytes"
	"encoding/gob"
	"time"

	"golang.org/x/sys/unix"
)

type SysInfo unix.Sysinfo_t

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
)

// GetUptime returns uptime in human readable form
func (s *SysInfo) GetUptime() time.Duration {
	return time.Duration(s.Uptime) * time.Second
}

func (s *SysInfo) GetTotalRAMInKB() uint64 {
	return s.Totalram / uint64(KB)
}

func (s *SysInfo) GetTotalRAMInMB() uint64 {
	return s.Totalram / uint64(MB)
}

func (s *SysInfo) GetTotalRAMInGB() float64 {
	return float64(s.Totalram) / GB
}

func (s *SysInfo) GetTotalHighRAMInKB() uint64 {
	return s.Totalhigh / uint64(KB)
}

func (s *SysInfo) GetTotalHighRAMInMB() uint64 {
	return s.Totalhigh / uint64(MB)
}

func (s *SysInfo) GetTotalHighRAMInGB() float64 {
	return float64(s.Totalhigh) / KB
}

func (s *SysInfo) GetFreeRAMInKB() uint64 {
	return s.Freeram / uint64(KB)
}

func (s *SysInfo) GetFreeRAMInMB() uint64 {
	return s.Freeram / uint64(MB)
}

func (s *SysInfo) GetFreeRAMInGB() float64 {
	return float64(s.Freeram) / GB
}

func (s *SysInfo) GetFreeHighRAMInKB() uint64 {
	return s.Freehigh / uint64(KB)
}

func (s *SysInfo) GetFreeHighRAMInMB() uint64 {
	return s.Freehigh / uint64(MB)
}

func (s *SysInfo) GetFreeHighRAMInGB() float64 {
	return float64(s.Freehigh) / GB
}

func (s *SysInfo) GetBufferRAMInKB() uint64 {
	return s.Bufferram / uint64(KB)
}

func (s *SysInfo) GetBufferRAMInMB() uint64 {
	return s.Bufferram / uint64(MB)
}

func (s *SysInfo) GetBufferRAMInGB() float64 {
	return float64(s.Bufferram) / GB
}

func (s *SysInfo) GetSharedRAMInKB() uint64 {
	return s.Sharedram / uint64(KB)
}

func (s *SysInfo) GetSharedRAMInMB() uint64 {
	return s.Sharedram / uint64(MB)
}

func (s *SysInfo) GetSharedRAMInGB() float64 {
	return float64(s.Sharedram) / GB
}

func (s *SysInfo) GetTotalSwapInKB() uint64 {
	return s.Totalswap / uint64(KB)
}

func (s *SysInfo) GetTotalSwapInMB() uint64 {
	return s.Totalswap / uint64(MB)
}

func (s *SysInfo) GetTotalSwapInGB() float64 {
	return float64(s.Totalswap) / GB
}

func (s *SysInfo) GetFreeSwapInKB() uint64 {
	return s.Freeswap / uint64(KB)
}

func (s *SysInfo) GetFreeSwapInMB() uint64 {
	return s.Freeswap / uint64(MB)
}

func (s *SysInfo) GetFreeSwapInGB() float64 {
	return float64(s.Freeswap) / GB
}

func GetSysInfo() (*SysInfo, error) {
	sys := &SysInfo{}
	info := &unix.Sysinfo_t{}
	if err := unix.Sysinfo(info); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	if err := enc.Encode(info); err != nil {
		return nil, err
	}

	if err := dec.Decode(&sys); err != nil {
		return nil, err
	}

	return sys, nil
}
