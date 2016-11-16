package sysinfo

import "golang.org/x/sys/unix"

type FS struct {
	Total uint64
	Free  uint64
	Used  uint64
}

func (f *FS) TotalSizeInKB() uint64 {
	return convertBToKB(f.Total)
}

func (f *FS) TotalSizeInMB() float64 {
	return convertBToMB(f.Total)
}

func (f *FS) TotalSizeInGB() float64 {
	return convertBToGB(f.Total)
}

func (f *FS) FreeSpaceInKB() uint64 {
	return convertBToKB(f.Free)
}

func (f *FS) FreeSpaceInMB() float64 {
	return convertBToMB(f.Free)
}

func (f *FS) FreeSpaceInGB() float64 {
	return convertBToGB(f.Free)
}

func (f *FS) UsedSpaceInKB() uint64 {
	return convertBToKB(f.Used)
}

func (f *FS) UsedSpaceInMB() float64 {
	return convertBToMB(f.Used)
}

func (f *FS) UsedSpaceInGB() float64 {
	return convertBToGB(f.Used)
}

func (f *FS) Get(path string) error {
	statfs := unix.Statfs_t{}
	if err := unix.Statfs(path, &statfs); err != nil {
		return err
	}

	f.Total = uint64(statfs.Bsize) * statfs.Blocks
	f.Free = uint64(statfs.Bsize) * statfs.Bavail
	f.Used = (uint64(statfs.Bsize) * statfs.Blocks) -
		(uint64(statfs.Bsize) * statfs.Bfree)

	return nil
}
