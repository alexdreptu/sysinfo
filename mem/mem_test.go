package mem_test

import (
	"testing"

	. "github.com/alexdreptu/sysinfo/mem"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/unix"
)

// test values
const (
	totalMemInBytes     uint64  = 8589934592
	totalMemInKibibytes float64 = 8388608
	totalMemInMebibytes float64 = 8192
	totalMemInGibibytes float64 = 8

	totalHighMemInBytes     uint64  = 2147483648
	totalHighMemInKibibytes float64 = 2097152
	totalHighMemInMebibytes float64 = 2048
	totalHighMemInGibibytes float64 = 2

	freeMemInBytes     uint64  = 1073741824
	freeMemInKibibytes float64 = 1048576
	freeMemInMebibytes float64 = 1024
	freeMemInGibibytes float64 = 1

	freeHighMemInBytes     uint64  = 536870912
	freeHighMemInKibibytes float64 = 524288
	freeHighMemInMebibytes float64 = 512
	freeHighMemInGibibytes float64 = 0.5

	availMemInBytes     uint64  = 2684354560
	availMemInKibibytes float64 = 2621440
	availMemInMebibytes float64 = 2560
	availMemInGibibytes float64 = 2.5

	bufferMemInBytes     uint64  = 2684354560
	bufferMemInKibibytes float64 = 2621440
	bufferMemInMebibytes float64 = 2560
	bufferMemInGibibytes float64 = 2.5

	cachedMemInBytes     uint64  = 536870912
	cachedMemInKibibytes float64 = 524288
	cachedMemInMebibytes float64 = 512
	cachedMemInGibibytes float64 = 0.5

	sharedMemInBytes     uint64  = 1073741824
	sharedMemInKibibytes float64 = 1048576
	sharedMemInMebibytes float64 = 1024
	sharedMemInGibibytes float64 = 1

	totalSwapInBytes     uint64  = 12884901888
	totalSwapInKibibytes float64 = 12582912
	totalSwapInMebibytes float64 = 12288
	totalSwapInGibibytes float64 = 12

	freeSwapInBytes     uint64  = 7516192768
	freeSwapInKibibytes float64 = 7340032
	freeSwapInMebibytes float64 = 7168
	freeSwapInGibibytes float64 = 7
)

const (
	usedMemInKibibytes float64 = totalMemInKibibytes - freeMemInKibibytes -
		bufferMemInKibibytes - cachedMemInKibibytes
	usedMemInMebibytes float64 = totalMemInMebibytes - freeMemInMebibytes -
		bufferMemInMebibytes - cachedMemInMebibytes
	usedMemInGibibytes float64 = totalMemInGibibytes - freeMemInGibibytes -
		bufferMemInGibibytes - cachedMemInGibibytes

	totalUsedMemInKibibytes float64 = totalMemInKibibytes - freeMemInKibibytes
	totalUsedMemInMebibytes float64 = totalMemInMebibytes - freeMemInMebibytes
	totalUsedMemInGibibytes float64 = totalMemInGibibytes - freeMemInGibibytes

	usedSwapInKibibytes float64 = totalSwapInKibibytes - freeSwapInKibibytes
	usedSwapInMebibytes float64 = totalSwapInMebibytes - freeSwapInMebibytes
	usedSwapInGibibytes float64 = totalSwapInGibibytes - freeSwapInGibibytes

	totalUsedInKibibytes float64 = totalMemInKibibytes - freeMemInKibibytes +
		totalSwapInKibibytes - freeSwapInKibibytes
	totalUsedInMebibytes float64 = totalMemInMebibytes - freeMemInMebibytes +
		totalSwapInMebibytes - freeSwapInMebibytes
	totalUsedInGibibytes float64 = totalMemInGibibytes - freeMemInGibibytes +
		totalSwapInGibibytes - freeSwapInGibibytes
)

const (
	procs        uint16 = 1520
	totalMem     uint64 = totalMemInBytes
	totalHighMem uint64 = totalHighMemInBytes
	freeMem      uint64 = freeMemInBytes
	freeHighMem  uint64 = freeHighMemInBytes
	availMem     uint64 = availMemInBytes
	bufferMem    uint64 = bufferMemInBytes
	cachedMem    uint64 = cachedMemInBytes
	sharedMem    uint64 = sharedMemInBytes
	totalSwap    uint64 = totalSwapInBytes
	freeSwap     uint64 = freeSwapInBytes
)

const delta = 0.01

// mock function
func sysinfo(info *unix.Sysinfo_t) error {
	info.Procs = procs
	info.Totalram = totalMem
	info.Totalhigh = totalHighMem
	info.Freeram = freeMem
	info.Freehigh = freeHighMem
	info.Bufferram = bufferMem
	info.Sharedram = sharedMem
	info.Totalswap = totalSwap
	info.Freeswap = freeSwap
	return nil
}

func TestMemConversion(t *testing.T) {
	mem := &Mem{}
	mem.F = sysinfo
	mem.AvailMem = availMem
	mem.CachedMem = cachedMem

	assert.NoError(t, mem.Fetch())

	assert.Equal(t, procs, mem.Procs)
	assert.Equal(t, totalMem, mem.TotalMem)
	assert.Equal(t, totalHighMem, mem.TotalHighMem)
	assert.Equal(t, freeMem, mem.FreeMem)
	assert.Equal(t, freeHighMem, mem.FreeHighMem)
	assert.Equal(t, availMem, mem.AvailMem)
	assert.Equal(t, bufferMem, mem.BufferMem)
	assert.Equal(t, cachedMem, mem.CachedMem)
	assert.Equal(t, sharedMem, mem.SharedMem)
	assert.Equal(t, totalSwap, mem.TotalSwap)
	assert.Equal(t, freeSwap, mem.FreeSwap)

	t.Run("total mem", func(t *testing.T) {
		assert.InDelta(t, totalMemInKibibytes, mem.TotalMemInKibibytes(), delta)
		assert.InDelta(t, totalMemInMebibytes, mem.TotalMemInMebibytes(), delta)
		assert.InDelta(t, totalMemInGibibytes, mem.TotalMemInGibibytes(), delta)
	})

	t.Run("total high mem", func(t *testing.T) {
		assert.InDelta(t, totalHighMemInKibibytes, mem.TotalHighMemInKibibytes(), delta)
		assert.InDelta(t, totalHighMemInMebibytes, mem.TotalHighMemInMebibytes(), delta)
		assert.InDelta(t, totalHighMemInGibibytes, mem.TotalHighMemInGibibytes(), delta)
	})

	t.Run("free mem", func(t *testing.T) {
		assert.InDelta(t, freeMemInKibibytes, mem.FreeMemInKibibytes(), delta)
		assert.InDelta(t, freeMemInMebibytes, mem.FreeMemInMebibytes(), delta)
		assert.InDelta(t, freeMemInGibibytes, mem.FreeMemInGibibytes(), delta)
	})

	t.Run("free high mem", func(t *testing.T) {
		assert.InDelta(t, freeHighMemInKibibytes, mem.FreeHighMemInKibibytes(), delta)
		assert.InDelta(t, freeHighMemInMebibytes, mem.FreeHighMemInMebibytes(), delta)
		assert.InDelta(t, freeHighMemInGibibytes, mem.FreeHighMemInGibibytes(), delta)
	})

	t.Run("avail mem", func(t *testing.T) {
		assert.InDelta(t, availMemInKibibytes, mem.AvailMemInKibibytes(), delta)
		assert.InDelta(t, availMemInMebibytes, mem.AvailMemInMebibytes(), delta)
		assert.InDelta(t, availMemInGibibytes, mem.AvailMemInGibibytes(), delta)
	})

	t.Run("used mem", func(t *testing.T) {
		assert.InDelta(t, usedMemInKibibytes, mem.UsedMemInKibibytes(), delta)
		assert.InDelta(t, usedMemInMebibytes, mem.UsedMemInMebibytes(), delta)
		assert.InDelta(t, usedMemInGibibytes, mem.UsedMemInGibibytes(), delta)
	})

	t.Run("total high mem", func(t *testing.T) {
		assert.InDelta(t, totalHighMemInKibibytes, mem.TotalHighMemInKibibytes(), delta)
		assert.InDelta(t, totalHighMemInMebibytes, mem.TotalHighMemInMebibytes(), delta)
		assert.InDelta(t, totalHighMemInGibibytes, mem.TotalHighMemInGibibytes(), delta)
	})

	t.Run("buffer mem", func(t *testing.T) {
		assert.InDelta(t, bufferMemInKibibytes, mem.BufferMemInKibibytes(), delta)
		assert.InDelta(t, bufferMemInMebibytes, mem.BufferMemInMebibytes(), delta)
		assert.InDelta(t, bufferMemInGibibytes, mem.BufferMemInGibibytes(), delta)
	})

	t.Run("cached mem", func(t *testing.T) {
		assert.InDelta(t, cachedMemInKibibytes, mem.CachedMemInKibibytes(), delta)
		assert.InDelta(t, cachedMemInMebibytes, mem.CachedMemInMebibytes(), delta)
		assert.InDelta(t, cachedMemInGibibytes, mem.CachedMemInGibibytes(), delta)
	})

	t.Run("shared mem", func(t *testing.T) {
		assert.InDelta(t, sharedMemInKibibytes, mem.SharedMemInKibibytes(), delta)
		assert.InDelta(t, sharedMemInMebibytes, mem.SharedMemInMebibytes(), delta)
		assert.InDelta(t, sharedMemInGibibytes, mem.SharedMemInGibibytes(), delta)
	})

	t.Run("total swap", func(t *testing.T) {
		assert.InDelta(t, totalSwapInKibibytes, mem.TotalSwapInKibibytes(), delta)
		assert.InDelta(t, totalSwapInMebibytes, mem.TotalSwapInMebibytes(), delta)
		assert.InDelta(t, totalSwapInGibibytes, mem.TotalSwapInGibibytes(), delta)
	})

	t.Run("free swap", func(t *testing.T) {
		assert.InDelta(t, freeSwapInKibibytes, mem.FreeSwapInKibibytes(), delta)
		assert.InDelta(t, freeSwapInMebibytes, mem.FreeSwapInMebibytes(), delta)
		assert.InDelta(t, freeSwapInGibibytes, mem.FreeSwapInGibibytes(), delta)
	})

	t.Run("used swap", func(t *testing.T) {
		assert.InDelta(t, usedSwapInKibibytes, mem.UsedSwapInKibibytes(), delta)
		assert.InDelta(t, usedSwapInMebibytes, mem.UsedSwapInMebibytes(), delta)
		assert.InDelta(t, usedSwapInGibibytes, mem.UsedSwapInGibibytes(), delta)
	})

	t.Run("total used", func(t *testing.T) {
		assert.InDelta(t, totalUsedInKibibytes, mem.TotalUsedInKibibytes(), delta)
		assert.InDelta(t, totalUsedInMebibytes, mem.TotalUsedInMebibytes(), delta)
		assert.InDelta(t, totalUsedInGibibytes, mem.TotalUsedInGibibytes(), delta)
	})

	t.Run("total used", func(t *testing.T) {
		assert.InDelta(t, totalUsedInKibibytes, mem.TotalUsedInKibibytes(), delta)
		assert.InDelta(t, totalUsedInMebibytes, mem.TotalUsedInMebibytes(), delta)
		assert.InDelta(t, totalUsedInGibibytes, mem.TotalUsedInGibibytes(), delta)
	})
}
