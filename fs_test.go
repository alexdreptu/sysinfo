package sysinfo_test

import (
	"testing"

	"github.com/alexdreptu/sysinfo"
)

func TestDisplayRootFSInfo(t *testing.T) {
	rootfs := sysinfo.FS{}
	if err := rootfs.Get("/"); err != nil {
		t.Fatal(err)
	}

	t.Logf("Root Total Size: %d KB\n", rootfs.TotalSizeInKB())
	t.Logf("Root Total Size: %.1f MB\n", rootfs.TotalSizeInMB())
	t.Logf("Root Total Size: %.1f GB\n", rootfs.TotalSizeInGB())

	t.Logf("Root Free Space: %d KB\n", rootfs.FreeSpaceInKB())
	t.Logf("Root Free Space: %.1f MB\n", rootfs.FreeSpaceInMB())
	t.Logf("Root Free Space: %.1f GB\n", rootfs.FreeSpaceInGB())

	t.Logf("Root Used Space: %d KB\n", rootfs.UsedSpaceInKB())
	t.Logf("Root Used Space: %.1f MB\n", rootfs.UsedSpaceInMB())
	t.Logf("Root Used Space: %.1f GB\n", rootfs.UsedSpaceInGB())
}

func TestDisplayHomeFSInfo(t *testing.T) {
	homefs := sysinfo.FS{}
	if err := homefs.Get("/home"); err != nil {
		t.Fatal(err)
	}

	t.Logf("Home Total Size: %d KB\n", homefs.TotalSizeInKB())
	t.Logf("Home Total Size: %.1f MB\n", homefs.TotalSizeInMB())
	t.Logf("Home Total Size: %.1f GB\n", homefs.TotalSizeInGB())

	t.Logf("Home Free Space: %d KB\n", homefs.FreeSpaceInKB())
	t.Logf("Home Free Space: %.1f MB\n", homefs.FreeSpaceInMB())
	t.Logf("Home Free Space: %.1f GB\n", homefs.FreeSpaceInGB())

	t.Logf("Home Used Space: %d KB\n", homefs.UsedSpaceInKB())
	t.Logf("Home Used Space: %.1f MB\n", homefs.UsedSpaceInMB())
	t.Logf("Home Used Space: %.1f GB\n", homefs.UsedSpaceInGB())
}
