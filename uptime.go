package sysinfo

import (
	"time"

	"golang.org/x/sys/unix"
)

type UptimeInfo struct {
	Uptime time.Duration
}

func (u *UptimeInfo) GetUptime() time.Duration {
	return u.Uptime * time.Second
}

func (u *UptimeInfo) Get() error {
	si := unix.Sysinfo_t{}
	if err := unix.Sysinfo(&si); err != nil {
		return err
	}

	u.Uptime = time.Duration(si.Uptime)
	return nil
}

func NewUptime() *UptimeInfo {
	return &UptimeInfo{}
}
