package mem_test

import (
	"testing"

	. "github.com/alexdreptu/sysinfo/mem"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/unix"
)

// Test values
const (
	totalMemInBytes     float64 = 8589934592
	totalMemInKibibytes float64 = 8388608
	totalMemInMebibytes float64 = 8192
	totalMemInGibibytes float64 = 8

	totalHighMemInBytes     float64 = 2147483648
	totalHighMemInKibibytes float64 = 2097152
	totalHighMemInMebibytes float64 = 2048
	totalHighMemInGibibytes float64 = 2

	freeMemInBytes     float64 = 1073741824
	freeMemInKibibytes float64 = 1048576
	freeMemInMebibytes float64 = 1024
	freeMemInGibibytes float64 = 1

	freeHighMemInBytes     float64 = 536870912
	freeHighMemInKibibytes float64 = 524288
	freeHighMemInMebibytes float64 = 512
	freeHighMemInGibibytes float64 = 0.5

	availMemInBytes     float64 = 2684354560
	availMemInKibibytes float64 = 2621440
	availMemInMebibytes float64 = 2560
	availMemInGibibytes float64 = 2.5

	usedMemInBytes     float64 = totalMemInBytes - freeMemInBytes - bufferMemInBytes - cachedMemInBytes
	usedMemInKibibytes float64 = totalMemInKibibytes - freeMemInKibibytes - bufferMemInKibibytes - cachedMemInKibibytes
	usedMemInMebibytes float64 = totalMemInMebibytes - freeMemInMebibytes - bufferMemInMebibytes - cachedMemInMebibytes
	usedMemInGibibytes float64 = totalMemInGibibytes - freeMemInGibibytes - bufferMemInGibibytes - cachedMemInGibibytes

	totalUsedMemInBytes     float64 = totalMemInBytes - freeMemInBytes
	totalUsedMemInKibibytes float64 = totalMemInKibibytes - freeMemInKibibytes
	totalUsedMemInMebibytes float64 = totalMemInMebibytes - freeMemInMebibytes
	totalUsedMemInGibibytes float64 = totalMemInGibibytes - freeMemInGibibytes

	bufferMemInBytes     float64 = 2684354560
	bufferMemInKibibytes float64 = 2621440
	bufferMemInMebibytes float64 = 2560
	bufferMemInGibibytes float64 = 2.5

	cachedMemInBytes     float64 = 536870912
	cachedMemInKibibytes float64 = 524288
	cachedMemInMebibytes float64 = 512
	cachedMemInGibibytes float64 = 0.5

	sharedMemInBytes     float64 = 1073741824
	sharedMemInKibibytes float64 = 1048576
	sharedMemInMebibytes float64 = 1024
	sharedMemInGibibytes float64 = 1

	totalSwapInBytes     float64 = 12884901888
	totalSwapInKibibytes float64 = 12582912
	totalSwapInMebibytes float64 = 12288
	totalSwapInGibibytes float64 = 12

	freeSwapInBytes     float64 = 7516192768
	freeSwapInKibibytes float64 = 7340032
	freeSwapInMebibytes float64 = 7168
	freeSwapInGibibytes float64 = 7

	usedSwapInBytes     float64 = totalSwapInBytes - freeSwapInBytes
	usedSwapInKibibytes float64 = totalSwapInKibibytes - freeSwapInKibibytes
	usedSwapInMebibytes float64 = totalSwapInMebibytes - freeSwapInMebibytes
	usedSwapInGibibytes float64 = totalSwapInGibibytes - freeSwapInGibibytes

	totalUsedInBytes     float64 = totalMemInBytes - freeMemInBytes + totalSwapInBytes - freeSwapInBytes
	totalUsedInKibibytes float64 = totalMemInKibibytes - freeMemInKibibytes + totalSwapInKibibytes - freeSwapInKibibytes
	totalUsedInMebibytes float64 = totalMemInMebibytes - freeMemInMebibytes + totalSwapInMebibytes - freeSwapInMebibytes
	totalUsedInGibibytes float64 = totalMemInGibibytes - freeMemInGibibytes + totalSwapInGibibytes - freeSwapInGibibytes
)

const (
	procs        uint16 = 1520
	totalMem     uint64 = uint64(totalMemInBytes)
	totalHighMem uint64 = uint64(totalHighMemInBytes)
	freeMem      uint64 = uint64(freeMemInBytes)
	freeHighMem  uint64 = uint64(freeHighMemInBytes)
	availMem     uint64 = uint64(availMemInBytes)
	bufferMem    uint64 = uint64(bufferMemInBytes)
	cachedMem    uint64 = uint64(cachedMemInBytes)
	sharedMem    uint64 = uint64(sharedMemInBytes)
	totalSwap    uint64 = uint64(totalSwapInBytes)
	freeSwap     uint64 = uint64(freeSwapInBytes)
)

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
	mem.Fetch()

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
		assert.Equal(t, totalMemInKibibytes, mem.TotalMemInKibibytes())
		assert.Equal(t, totalMemInMebibytes, mem.TotalMemInMebibytes())
		assert.Equal(t, totalMemInGibibytes, mem.TotalMemInGibibytes())
	})

	t.Run("total high mem", func(t *testing.T) {
		assert.Equal(t, totalHighMemInKibibytes, mem.TotalHighMemInKibibytes())
		assert.Equal(t, totalHighMemInMebibytes, mem.TotalHighMemInMebibytes())
		assert.Equal(t, totalHighMemInGibibytes, mem.TotalHighMemInGibibytes())
	})

	t.Run("free mem", func(t *testing.T) {
		assert.Equal(t, freeMemInKibibytes, mem.FreeMemInKibibytes())
		assert.Equal(t, freeMemInMebibytes, mem.FreeMemInMebibytes())
		assert.Equal(t, freeMemInGibibytes, mem.FreeMemInGibibytes())
	})

	t.Run("free high mem", func(t *testing.T) {
		assert.Equal(t, freeHighMemInKibibytes, mem.FreeHighMemInKibibytes())
		assert.Equal(t, freeHighMemInMebibytes, mem.FreeHighMemInMebibytes())
		assert.Equal(t, freeHighMemInGibibytes, mem.FreeHighMemInGibibytes())
	})

	t.Run("avail mem", func(t *testing.T) {
		assert.Equal(t, availMemInKibibytes, mem.AvailMemInKibibytes())
		assert.Equal(t, availMemInMebibytes, mem.AvailMemInMebibytes())
		assert.Equal(t, availMemInGibibytes, mem.AvailMemInGibibytes())
	})

	t.Run("used mem", func(t *testing.T) {
		assert.Equal(t, usedMemInKibibytes, mem.UsedMemInKibibytes())
		assert.Equal(t, usedMemInMebibytes, mem.UsedMemInMebibytes())
		assert.Equal(t, usedMemInGibibytes, mem.UsedMemInGibibytes())
	})

	t.Run("total high mem", func(t *testing.T) {
		assert.Equal(t, totalHighMemInKibibytes, mem.TotalHighMemInKibibytes())
		assert.Equal(t, totalHighMemInMebibytes, mem.TotalHighMemInMebibytes())
		assert.Equal(t, totalHighMemInGibibytes, mem.TotalHighMemInGibibytes())
	})

	t.Run("buffer mem", func(t *testing.T) {
		assert.Equal(t, bufferMemInKibibytes, mem.BufferMemInKibibytes())
		assert.Equal(t, bufferMemInMebibytes, mem.BufferMemInMebibytes())
		assert.Equal(t, bufferMemInGibibytes, mem.BufferMemInGibibytes())
	})

	t.Run("cached mem", func(t *testing.T) {
		assert.Equal(t, cachedMemInKibibytes, mem.CachedMemInKibibytes())
		assert.Equal(t, cachedMemInMebibytes, mem.CachedMemInMebibytes())
		assert.Equal(t, cachedMemInGibibytes, mem.CachedMemInGibibytes())
	})

	t.Run("shared mem", func(t *testing.T) {
		assert.Equal(t, sharedMemInKibibytes, mem.SharedMemInKibibytes())
		assert.Equal(t, sharedMemInMebibytes, mem.SharedMemInMebibytes())
		assert.Equal(t, sharedMemInGibibytes, mem.SharedMemInGibibytes())
	})

	t.Run("total swap", func(t *testing.T) {
		assert.Equal(t, totalSwapInKibibytes, mem.TotalSwapInKibibytes())
		assert.Equal(t, totalSwapInMebibytes, mem.TotalSwapInMebibytes())
		assert.Equal(t, totalSwapInGibibytes, mem.TotalSwapInGibibytes())
	})

	t.Run("free swap", func(t *testing.T) {
		assert.Equal(t, freeSwapInKibibytes, mem.FreeSwapInKibibytes())
		assert.Equal(t, freeSwapInMebibytes, mem.FreeSwapInMebibytes())
		assert.Equal(t, freeSwapInGibibytes, mem.FreeSwapInGibibytes())
	})

	t.Run("used swap", func(t *testing.T) {
		assert.Equal(t, usedSwapInKibibytes, mem.UsedSwapInKibibytes())
		assert.Equal(t, usedSwapInMebibytes, mem.UsedSwapInMebibytes())
		assert.Equal(t, usedSwapInGibibytes, mem.UsedSwapInGibibytes())
	})

	t.Run("total used", func(t *testing.T) {
		assert.Equal(t, totalUsedInKibibytes, mem.TotalUsedInKibibytes())
		assert.Equal(t, totalUsedInMebibytes, mem.TotalUsedInMebibytes())
		assert.Equal(t, totalUsedInGibibytes, mem.TotalUsedInGibibytes())
	})

	t.Run("total used", func(t *testing.T) {
		assert.Equal(t, totalUsedInKibibytes, mem.TotalUsedInKibibytes())
		assert.Equal(t, totalUsedInMebibytes, mem.TotalUsedInMebibytes())
		assert.Equal(t, totalUsedInGibibytes, mem.TotalUsedInGibibytes())
	})
}
