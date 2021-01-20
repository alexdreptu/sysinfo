package fs_test

import (
	"testing"

	"github.com/alexdreptu/sysinfo/fs"
	. "github.com/alexdreptu/sysinfo/fs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

// test values
const (
	bsize  int64  = 4096
	blocks uint64 = 14261319
	bavail uint64 = 1313231
	bfree  uint64 = 2618397
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
	buf.Bsize = bsize
	buf.Blocks = blocks
	buf.Bavail = bavail
	buf.Bfree = bfree
	return nil
}

func TestFSNew(t *testing.T) {
	_, err := fs.New("")
	require.Error(t, err)

	_, err = fs.New("/")
	require.NoError(t, err)
}

func TestFS(t *testing.T) {
	fs := &FS{Path: ""}
	require.Error(t, fs.Fetch())

	fs = &FS{Path: "/"}
	fs.F = statfs

	require.NoError(t, fs.Fetch())

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
