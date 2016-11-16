package sysinfo_test

import (
	"testing"
	"time"

	"github.com/alexdreptu/sysinfo"
)

func TestDisplayUptimeInfo(t *testing.T) {
	uptime := sysinfo.Uptime{}
	if err := uptime.Get(); err != nil {
		t.Fatal(err)
	}

	t.Logf("Uptime: %s\n", uptime)
	t.Logf("Up since (RFC1123): %s\n", uptime.UpSince())
	t.Logf("Up since (UnixDate): %s\n", uptime.UpSinceFormat(time.UnixDate))
}
