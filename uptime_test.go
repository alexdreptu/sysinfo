package sysinfo

import "testing"

func TestDisplayUptimeInfo(t *testing.T) {
	uptime := NewUptime()
	if err := uptime.Get(); err != nil {
		t.Fatal(err)
	}

	t.Logf("Uptime: %s\n", uptime.GetUptime())
	t.Logf("Uptime: %.1f seconds\n", uptime.GetUptime().Seconds())
	t.Logf("Uptime: %.1f minutes\n", uptime.GetUptime().Minutes())
	t.Logf("Uptime: %.1f hours\n", uptime.GetUptime().Hours())
}
