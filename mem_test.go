package sysinfo

import "testing"

func TestGetUptime(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Uptime: %s\n", mem.GetUptime())
}

func TestPrintNumOfProcs(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Number of procs: %d\n", mem.Procs)
}

func TestTotalMemInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Mem: %d KB\n", mem.TotalMemInKB())
}

func TestTotalMemInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Mem: %.1f MB\n", mem.TotalMemInMB())
}

func TestTotalMemInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Mem: %.1f GB\n", mem.TotalMemInGB())
}

func TestTotalHighMemInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total High Mem: %d KB\n", mem.TotalHighMemInKB())
}

func TestTotalHighMemInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total High Mem: %.1f MB\n", mem.TotalHighMemInMB())
}

func TestTotalHighMemInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total High Mem: %.1f GB\n", mem.TotalHighMemInGB())
}

func TestFreeMemInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free Mem: %d KB\n", mem.FreeMemInKB())
}

func TestFreeMemInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free Mem: %.1f MB\n", mem.FreeMemInMB())
}

func TestFreeMemInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free Mem: %.1f GB\n", mem.FreeMemInGB())
}

func TestFreeHighMemInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free High Mem: %d KB\n", mem.FreeHighMemInKB())
}

func TestFreeHighMemInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free High Mem: %.1f MB\n", mem.FreeHighMemInMB())
}

func TestFreeHighMemInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free High Mem: %.1f GB\n", mem.FreeHighMemInGB())
}

func TestAvailMemInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Avail Mem: %d KB\n", mem.AvailMemInKB())
}

func TestAvailMemInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Avail Mem: %.1f MB\n", mem.AvailMemInMB())
}

func TestAvailMemInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Avail Mem: %.1f GB\n", mem.AvailMemInGB())
}

func TestUsedMemInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Used Mem: %d KB\n", mem.UsedMemInKB())
}

func TestUsedMemInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Used Mem: %.1f MB\n", mem.UsedMemInMB())
}

func TestUsedMemInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Used Mem: %.1f GB\n", mem.UsedMemInGB())
}

func TestTotalUsedMemInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Used Mem: %d KB\n", mem.TotalUsedMemInKB())
}

func TestTotalUsedMemInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Used Mem: %1.f MB\n", mem.TotalUsedMemInMB())
}

func TestTotalUsedMemInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Used Mem: %.1f GB\n", mem.TotalUsedMemInGB())
}

func TestTotalUsedInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Used: %d KB\n", mem.TotalUsedInKB())
}

func TestTotalUsedInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Used: %.1f MB\n", mem.TotalUsedInMB())
}

func TestTotalUsedInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Used: %.1f GB\n", mem.TotalUsedInGB())
}

func TestBufferMemInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Buffer Mem: %d KB\n", mem.BufferMemInKB())
}

func TestBufferMemInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Buffer Mem: %.1f MB\n", mem.BufferMemInMB())
}

func TestBufferMemInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Buffer Mem: %.1f GB\n", mem.BufferMemInGB())
}

func TestSharedMemInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Shared Mem: %d KB\n", mem.SharedMemInKB())
}

func TestSharedMemInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Shared Mem: %.1f MB\n", mem.SharedMemInMB())
}

func TestSharedMemInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Shared Mem: %.1f GB\n", mem.SharedMemInGB())
}

func TestTotalSwapInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Swap: %d KB\n", mem.TotalSwapInKB())
}

func TestTotalSwapInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Swap: %.1f MB\n", mem.TotalSwapInMB())
}

func TestTotalSwapInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Swap: %.1f GB\n", mem.TotalSwapInGB())
}

func TestFreeSwapInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free Swap: %d KB\n", mem.FreeSwapInKB())
}

func TestFreeSwapInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free Swap: %.1f MB\n", mem.FreeSwapInMB())
}

func TestFreeSwapInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free Swap: %.1f GB\n", mem.FreeSwapInGB())
}

func TestUsedSwapInKB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Used Swap: %d KB\n", mem.UsedSwapInKB())
}

func TestUsedSwapInMB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Used Swap: %.1f MB\n", mem.UsedSwapInMB())
}

func TestUsedSwapInGB(t *testing.T) {
	mem := NewMem()
	if err := mem.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Used Swap: %.1f GB\n", mem.UsedSwapInGB())
}
