package uptime_test

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/alexdreptu/sysinfo/uptime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

// test values
const (
	uptime int64 = 186072

	uptimeSeconds = 12
	uptimeMinutes = 41
	uptimenHours  = 3
	uptimeDays    = 2

	sep byte = '-'
)

// mock function
func sysinfo(info *unix.Sysinfo_t) error {
	info.Uptime = uptime
	return nil
}

func TestUptimeNewWithSeparator(t *testing.T) {
	uptime, err := New(sep)
	require.NoError(t, err)
	assert.True(t, strings.Contains(uptime.String(), string(sep)))

	uptime, err = New()
	require.NoError(t, err)
	assert.True(t, strings.Contains(uptime.String(), ":"))
}

func TestUptime(t *testing.T) {
	t.Run("fetch with custom separator", func(t *testing.T) {
		uptime := &Uptime{Sep: sep}
		uptime.F = sysinfo

		uptimeFormat := fmt.Sprintf("%dd%s%dh%s%dm%s%ds", uptimeDays, string(sep),
			uptimenHours, string(sep), uptimeMinutes, string(sep), uptimeSeconds)

		require.NoError(t, uptime.Fetch())
		assert.Equal(t, uptimeFormat, uptime.String())
	})

	t.Run("fetch with default separator", func(t *testing.T) {
		uptime := &Uptime{}
		uptime.F = sysinfo

		uptimeFormat := fmt.Sprintf("%dd:%dh:%dm:%ds", uptimeDays, uptimenHours,
			uptimeMinutes, uptimeSeconds)

		require.NoError(t, uptime.Fetch())
		assert.Equal(t, uptimeFormat, uptime.String())
	})
}
