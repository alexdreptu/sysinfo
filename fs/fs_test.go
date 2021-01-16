package fs_test

import (
	"testing"

	. "github.com/alexdreptu/sysinfo/fs"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/unix"
)

// test values
const (
	bsize  float64 = 4096
	blocks float64 = 14261319
	bavail float64 = 1313231
	bfree  float64 = 2618397
)

const (
	totalSpaceInKibibytes float64 = 57045276.000000
	totalSpaceInMebibytes float64 = 55708.277344
	totalSpaceInGibibytes float64 = 54.402615

	freeSpaceInKibibytes float64 = 5252924.000000
	freeSpaceInMebibytes float64 = 5129.808594
	freeSpaceInGibibytes float64 = 5.009579

	usedSpaceInKibibytes float64 = 46571688.000000
	usedSpaceInMebibytes float64 = 45480.164062
	usedSpaceInGibibytes float64 = 44.414223
)

const delta = 0.01

// mock function
func statfs(path string, buf *unix.Statfs_t) error {
	buf.Bsize = int64(bsize)
	buf.Blocks = uint64(blocks)
	buf.Bavail = uint64(bavail)
	buf.Bfree = uint64(bfree)
	return nil
}

func TestFS(t *testing.T) {
	fs := &FS{}
	fs.F = statfs

	assert.NoError(t, fs.Fetch(""))

	assert.InDelta(t, totalSpaceInKibibytes, fs.TotalSpaceInKibibytes(), delta)
	assert.InDelta(t, totalSpaceInMebibytes, fs.TotalSpaceInMebibytes(), delta)
	assert.InDelta(t, totalSpaceInGibibytes, fs.TotalSpaceInGibibytes(), delta)

	assert.InDelta(t, freeSpaceInKibibytes, fs.FreeSpaceInKibibytes(), delta)
	assert.InDelta(t, freeSpaceInMebibytes, fs.FreeSpaceInMebibytes(), delta)
	assert.InDelta(t, freeSpaceInGibibytes, fs.FreeSpaceInGibibytes(), delta)

	assert.InDelta(t, usedSpaceInKibibytes, fs.UsedSpaceInKibibytes(), delta)
	assert.InDelta(t, usedSpaceInMebibytes, fs.UsedSpaceInMebibytes(), delta)
	assert.InDelta(t, usedSpaceInGibibytes, fs.UsedSpaceInGibibytes(), delta)
}
