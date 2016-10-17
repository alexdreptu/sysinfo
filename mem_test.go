package sysinfo

import "testing"

func TestPrintNumOfProcs(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Number of procs: %d\n", mem.Procs)
}

func TestDisplayMemInfo(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}

	t.Logf("Total Mem: %d KB\n", mem.TotalMemInKB())
	t.Logf("Total Mem: %.1f MB\n", mem.TotalMemInMB())
	t.Logf("Total Mem: %.1f GB\n", mem.TotalMemInGB())

	t.Logf("Total High Mem: %d KB\n", mem.TotalHighMemInKB())
	t.Logf("Total High Mem: %.1f MB\n", mem.TotalHighMemInMB())
	t.Logf("Total High Mem: %.1f GB\n", mem.TotalHighMemInGB())

	t.Logf("Free Mem: %d KB\n", mem.FreeMemInKB())
	t.Logf("Free Mem: %.1f MB\n", mem.FreeMemInMB())
	t.Logf("Free Mem: %.1f GB\n", mem.FreeMemInGB())

	t.Logf("Free High Mem: %d KB\n", mem.FreeHighMemInKB())
	t.Logf("Free High Mem: %.1f MB\n", mem.FreeHighMemInMB())
	t.Logf("Free High Mem: %.1f GB\n", mem.FreeHighMemInGB())

	t.Logf("Avail Mem: %d KB\n", mem.AvailMemInKB())
	t.Logf("Avail Mem: %.1f MB\n", mem.AvailMemInMB())
	t.Logf("Avail Mem: %.1f GB\n", mem.AvailMemInGB())

	t.Logf("Used Mem: %d KB\n", mem.UsedMemInKB())
	t.Logf("Used Mem: %.1f MB\n", mem.UsedMemInMB())
	t.Logf("Used Mem: %.1f GB\n", mem.UsedMemInGB())

	t.Logf("Total Used Mem: %d KB\n", mem.TotalUsedMemInKB())
	t.Logf("Total Used Mem: %1.f MB\n", mem.TotalUsedMemInMB())
	t.Logf("Total Used Mem: %.1f GB\n", mem.TotalUsedMemInGB())

	t.Logf("Total Used: %d KB\n", mem.TotalUsedInKB())
	t.Logf("Total Used: %.1f MB\n", mem.TotalUsedInMB())
	t.Logf("Total Used: %.1f GB\n", mem.TotalUsedInGB())

	t.Logf("Buffer Mem: %d KB\n", mem.BufferMemInKB())
	t.Logf("Buffer Mem: %.1f MB\n", mem.BufferMemInMB())
	t.Logf("Buffer Mem: %.1f GB\n", mem.BufferMemInGB())

	t.Logf("Shared Mem: %d KB\n", mem.SharedMemInKB())
	t.Logf("Shared Mem: %.1f MB\n", mem.SharedMemInMB())
	t.Logf("Shared Mem: %.1f GB\n", mem.SharedMemInGB())

	t.Logf("Total Swap: %d KB\n", mem.TotalSwapInKB())
	t.Logf("Total Swap: %.1f MB\n", mem.TotalSwapInMB())
	t.Logf("Total Swap: %.1f GB\n", mem.TotalSwapInGB())

	t.Logf("Free Swap: %d KB\n", mem.FreeSwapInKB())
	t.Logf("Free Swap: %.1f MB\n", mem.FreeSwapInMB())
	t.Logf("Free Swap: %.1f GB\n", mem.FreeSwapInGB())

	t.Logf("Used Swap: %d KB\n", mem.UsedSwapInKB())
	t.Logf("Used Swap: %.1f MB\n", mem.UsedSwapInMB())
	t.Logf("Used Swap: %.1f GB\n", mem.UsedSwapInGB())
}
