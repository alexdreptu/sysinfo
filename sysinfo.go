package sysinfo

import (
	"bytes"
	"encoding/gob"

	"golang.org/x/sys/unix"
)

type SysInfo unix.Sysinfo_t

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
)

func (s *SysInfo) GetTotalRAMinKB() uint64 {
	return s.Totalram / uint64(KB)
}

func (s *SysInfo) GetTotalRAMinMB() uint64 {
	return s.Totalram / uint64(MB)
}

func (s *SysInfo) GetTotalRAMinGB() float64 {
	return float64(s.Totalram) / GB
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
