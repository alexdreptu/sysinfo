package fs

import (
	"errors"

	"github.com/alexdreptu/sysinfo/convert"
	"golang.org/x/sys/unix"
)

var ErrMissingPath = errors.New("path is missing")

type fetchFunc func(path string, buf *unix.Statfs_t) error

type FS struct {
	Total uint64
	Free  uint64
	Used  uint64
	Path  string
	F     fetchFunc
}

func (f *FS) TotalSpaceInKibibytes() float64 {
	return (convert.Size(f.Total) * convert.Byte).ToKibibytes()
}

func (f *FS) TotalSpaceInMebibytes() float64 {
	return (convert.Size(f.Total) * convert.Byte).ToMebibytes()
}

func (f *FS) TotalSpaceInGibibytes() float64 {
	return (convert.Size(f.Total) * convert.Byte).ToGibibytes()
}

func (f *FS) FreeSpaceInKibibytes() float64 {
	return (convert.Size(f.Free) * convert.Byte).ToKibibytes()
}

func (f *FS) FreeSpaceInMebibytes() float64 {
	return (convert.Size(f.Free) * convert.Byte).ToMebibytes()
}

func (f *FS) FreeSpaceInGibibytes() float64 {
	return (convert.Size(f.Free) * convert.Byte).ToGibibytes()
}

func (f *FS) UsedSpaceInKibibytes() float64 {
	return (convert.Size(f.Used) * convert.Byte).ToKibibytes()
}

func (f *FS) UsedSpaceInMebibytes() float64 {
	return (convert.Size(f.Used) * convert.Byte).ToMebibytes()
}

func (f *FS) UsedSpaceInGibibytes() float64 {
	return (convert.Size(f.Used) * convert.Byte).ToGibibytes()
}

// Fetch updates the FS struct woth new values
func (f *FS) Fetch() error {
	if f.F == nil {
		f.F = unix.Statfs
	}

	if f.Path == "" {
		return ErrMissingPath
	}

	statfs := unix.Statfs_t{}
	if err := f.F(f.Path, &statfs); err != nil {
		return err
	}

	f.Total = uint64(statfs.Bsize) * statfs.Blocks
	f.Free = uint64(statfs.Bsize) * statfs.Bavail
	f.Used = (uint64(statfs.Bsize) * statfs.Blocks) -
		(uint64(statfs.Bsize) * statfs.Bfree)

	return nil
}

func New(path string) (*FS, error) {
	fs := &FS{Path: path}

	if err := fs.Fetch(); err != nil {
		return &FS{}, err
	}

	return fs, nil
}
