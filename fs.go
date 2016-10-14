package sysinfo

import "golang.org/x/sys/unix"

type FS struct {
	Total uint64
	Free  uint64
	Used  uint64
}

func (f *FS) GetTotalSizeInKB() uint64 {
	return ConvertBToKB(f.Total)
}

func (f *FS) GetTotalSizeInMB() float64 {
	return ConvertBToMB(f.Total)
}

func (f *FS) GetTotalSizeInGB() float64 {
	return ConvertBToGB(f.Total)
}

func (f *FS) GetFreeSpaceInKB() uint64 {
	return ConvertBToKB(f.Free)
}

func (f *FS) GetFreeSpaceInMB() float64 {
	return ConvertBToMB(f.Free)
}

func (f *FS) GetFreeSpaceInGB() float64 {
	return ConvertBToGB(f.Free)
}

func (f *FS) GetUsedSpaceInKB() uint64 {
	return ConvertBToKB(f.Used)
}

func (f *FS) GetUsedSpaceInMB() float64 {
	return ConvertBToMB(f.Used)
}

func (f *FS) GetUsedSpaceInGB() float64 {
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
