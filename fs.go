package sysinfo

import "golang.org/x/sys/unix"

type FSInfo struct {
	Total uint64
	Free  uint64
	Used  uint64
}

func (f *FSInfo) TotalSizeInKB() uint64 {
	return ConvertBToKB(f.Total)
}

func (f *FSInfo) TotalSizeInMB() float64 {
	return ConvertBToMB(f.Total)
}

func (f *FSInfo) TotalSizeInGB() float64 {
	return ConvertBToGB(f.Total)
}

func (f *FSInfo) FreeSpaceInKB() uint64 {
	return ConvertBToKB(f.Free)
}

func (f *FSInfo) FreeSpaceInMB() float64 {
	return ConvertBToMB(f.Free)
}

func (f *FSInfo) FreeSpaceInGB() float64 {
	return ConvertBToGB(f.Free)
}

func (f *FSInfo) UsedSpaceInKB() uint64 {
	return ConvertBToKB(f.Used)
}

func (f *FSInfo) UsedSpaceInMB() float64 {
	return ConvertBToMB(f.Used)
}

func (f *FSInfo) UsedSpaceInGB() float64 {
	return ConvertBToGB(f.Used)
}

func (f *FSInfo) Get(path string) error {
	statfs := unix.Statfs_t{}
	if err := unix.Statfs(path, &statfs); err != nil {
		return err
	}

	f.Total = uint64(statfs.Bsize) * statfs.Blocks
	f.Free = uint64(statfs.Bsize) * statfs.Bavail
	f.Used = (uint64(statfs.Bsize) * statfs.Blocks) - (uint64(statfs.Bsize) * statfs.Bfree)

	return nil
}

func NewFS() *FSInfo { return &FSInfo{} }
