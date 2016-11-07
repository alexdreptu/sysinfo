package sysinfo

import (
	"testing"
	"time"
)

func TestDisplayUptimeInfo(t *testing.T) {
	uptime := NewUptime()
	if err := uptime.Get(); err != nil {
		t.Fatal(err)
	}

	t.Logf("Uptime: %s\n", uptime)
	t.Logf("Up since (RFC1123): %s\n", uptime.UpSince())
	t.Logf("Up since (UnixDate): %s\n", uptime.UpSinceFormat(time.UnixDate))
}
