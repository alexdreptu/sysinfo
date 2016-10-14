package sysinfo

import "golang.org/x/sys/unix"

type FS struct {
	Total uint64
	Free  uint64
	Used  uint64
}

func (f *FS) TotalSizeInKB() uint64 {
	return ConvertBToKB(f.Total)
}

func (f *FS) TotalSizeInMB() float64 {
	return ConvertBToMB(f.Total)
}

func (f *FS) TotalSizeInGB() float64 {
	return ConvertBToGB(f.Total)
}

func (f *FS) FreeSpaceInKB() uint64 {
	return ConvertBToKB(f.Free)
}

func (f *FS) FreeSpaceInMB() float64 {
	return ConvertBToMB(f.Free)
}

func (f *FS) FreeSpaceInGB() float64 {
	return ConvertBToGB(f.Free)
}

func (f *FS) UsedSpaceInKB() uint64 {
	return ConvertBToKB(f.Used)
}

func (f *FS) UsedSpaceInMB() float64 {
	return ConvertBToMB(f.Used)
}

func (f *FS) UsedSpaceInGB() float64 {
	return ConvertBToGB(f.Used)
}

func (f *FS) Get(path string) error {
	statfs := &unix.Statfs_t{}
	if err := unix.Statfs(path, statfs); err != nil {
		return err
	}

	f.Total = uint64(statfs.Bsize) * statfs.Blocks
	f.Free = uint64(statfs.Bsize) * statfs.Bavail
	f.Used = f.Total - f.Free

	return nil
}
