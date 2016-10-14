package sysinfo

import "testing"

func TestGetUptime(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Uptime: %s\n", sys.GetUptime())
}

func TestNumberOfProcs(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Number of procs: %d\n", sys.Procs)
}

func TestGetTotalRAMInKB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total RAM: %d KB\n", sys.TotalRAMInKB())
}

func TestGetTotalRAMInMB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total RAM: %.1f MB\n", sys.TotalRAMInMB())
}

func TestGetTotalRAMInGB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total RAM: %.1f GB\n", sys.TotalRAMInGB())
}

func TestGetTotalHighRAMInKB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total High RAM: %d KB\n", sys.TotalHighRAMInKB())
}

func TestGetTotalHighRAMInMB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total High RAM: %.1f MB\n", sys.TotalHighRAMInMB())
}

func TestGetTotalHighRAMInGB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total High RAM: %.1f GB\n", sys.TotalHighRAMInGB())
}

func TestGetFreeRAMInKB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free RAM: %d KB\n", sys.FreeRAMInKB())
}

func TestGetFreeRAMInMB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free RAM: %.1f MB\n", sys.FreeRAMInMB())
}

func TestGetFreeRAMInGB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free RAM: %.1f GB\n", sys.FreeRAMInGB())
}

func TestGetFreeHighRAMInKB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free High RAM: %d KB\n", sys.FreeHighRAMInKB())
}

func TestGetFreeHighRAMInMB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free High RAM: %.1f MB\n", sys.FreeHighRAMInMB())
}

func TestGetFreeHighRAMInGB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free High RAM: %.1f GB\n", sys.FreeHighRAMInGB())
}

func TestGetAvailRAMInKB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Avail RAM: %d KB\n", sys.AvailRAMInKB())
}

func TestGetAvailRAMInMB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Avail RAM: %.1f MB\n", sys.AvailRAMInMB())
}

func TestGetAvailRAMInGB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Avail RAM: %.1f GB\n", sys.AvailRAMInGB())
}

func TestGetBufferRAMInKB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Buffer RAM: %d KB\n", sys.BufferRAMInKB())
}

func TestGetBufferRAMInMB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Buffer RAM: %.1f MB\n", sys.BufferRAMInMB())
}

func TestGetBufferRAMInGB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Buffer RAM: %.1f GB\n", sys.BufferRAMInGB())
}

func TestSharedRAMInKB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Shared RAM: %d KB\n", sys.SharedRAMInKB())
}

func TestSharedRAMInMB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Shared RAM: %.1f MB\n", sys.SharedRAMInMB())
}

func TestSharedRAMInGB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Shared RAM: %.1f GB\n", sys.SharedRAMInGB())
}

func TestTotalSwapInKB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Swap: %d KB\n", sys.TotalSwapInKB())
}

func TestTotalSwapInMB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Swap: %.1f MB\n", sys.TotalSwapInMB())
}

func TestTotalSwapInGB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Total Swap: %.1f GB\n", sys.TotalSwapInGB())
}

func TestFreeSwapInKB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free Swap: %d KB\n", sys.FreeSwapInKB())
}

func TestFreeSwapInMB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free Swap: %.1f MB\n", sys.FreeSwapInMB())
}

func TestFreeSwapInGB(t *testing.T) {
	sys := &Sysinfo{}
	if err := sys.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Free Swap: %.1f GB\n", sys.FreeSwapInGB())
}
