[![Build Status](https://travis-ci.org/alexdreptu/sysinfo.svg?branch=master)](https://travis-ci.org/alexdreptu/sysinfo)
[![GoDoc](https://godoc.org/github.com/alexdreptu/sysinfo?status.svg)](https://godoc.org/github.com/alexdreptu/sysinfo)
[![Platform](https://img.shields.io/badge/platform-Linux-5272b4.svg)](https://www.linuxfoundation.org/)
[![License](https://img.shields.io/badge/license-MIT-5272b4.svg)](https://github.com/alexdreptu/sysinfo/blob/master/LICENSE)

#### In development

### Examples

#### CPU Info

```Go
cpu := sysinfo.CPU{}
if err := cpu.Get(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}

fmt.Printf("Name: %s\n", cpu.Name)
fmt.Printf("Model: %d\n", cpu.Model)
fmt.Printf("Vendor: %s\n", cpu.Vendor)
fmt.Printf("Family: %d\n", cpu.Family)
fmt.Printf("Stepping: %d\n", cpu.Stepping)
fmt.Printf("Cores: %d\n", cpu.Cores)
fmt.Printf("Cache Size: %d KB\n", cpu.CacheSize)
fmt.Printf("Cache Size: %.1f MB\n", cpu.CacheSizeInMB())
fmt.Printf("Min Frequency: %d Khz\n", cpu.MinFreq)
fmt.Printf("Max Frequency: %d Khz\n", cpu.MaxFreq)
fmt.Printf("Min Frequency: %d Mhz\n", cpu.MinFreqInMhz())
fmt.Printf("Max Frequency: %d Mhz\n", cpu.MaxFreqInMhz())
fmt.Printf("Min Frequency: %.1f Ghz\n", cpu.MinFreqInGhz())
fmt.Printf("Max Frequency: %.1f Ghz\n", cpu.MaxFreqInGhz())
fmt.Printf("Flags %s\n", cpu.Flags)
```

#### Filesystem Info

```Go
// get disk usage for root
rootfs := sysinfo.FS{}
if err := rootfs.Get("/"); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}

fmt.Printf("Root Total Size: %d KB\n", rootfs.TotalSizeInKB())
fmt.Printf("Root Total Size: %.1f MB\n", rootfs.TotalSizeInMB())
fmt.Printf("Root Total Size: %.1f GB\n", rootfs.TotalSizeInGB())

fmt.Printf("Root Free Space: %d KB\n", rootfs.FreeSpaceInKB())
fmt.Printf("Root Free Space: %.1f MB\n", rootfs.FreeSpaceInMB())
fmt.Printf("Root Free Space: %.1f GB\n", rootfs.FreeSpaceInGB())

fmt.Printf("Root Used Space: %d KB\n", rootfs.UsedSpaceInKB())
fmt.Printf("Root Used Space: %.1f MB\n", rootfs.UsedSpaceInMB())
fmt.Printf("Root Used Space: %.1f GB\n", rootfs.UsedSpaceInGB()
```

#### Memory Info

```Go
mem := sysinfo.Mem{}
if err := mem.Get(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}

fmt.Printf("Total Mem: %d KB\n", mem.TotalMemInKB())
fmt.Printf("Total Mem: %.1f MB\n", mem.TotalMemInMB())
fmt.Printf("Total Mem: %.1f GB\n", mem.TotalMemInGB())

fmt.Printf("Total High Mem: %d KB\n", mem.TotalHighMemInKB())
fmt.Printf("Total High Mem: %.1f MB\n", mem.TotalHighMemInMB())
fmt.Printf("Total High Mem: %.1f GB\n", mem.TotalHighMemInGB())

fmt.Printf("Free Mem: %d KB\n", mem.FreeMemInKB())
fmt.Printf("Free Mem: %.1f MB\n", mem.FreeMemInMB())
fmt.Printf("Free Mem: %.1f GB\n", mem.FreeMemInGB())

fmt.Printf("Free High Mem: %d KB\n", mem.FreeHighMemInKB())
fmt.Printf("Free High Mem: %.1f MB\n", mem.FreeHighMemInMB())
fmt.Printf("Free High Mem: %.1f GB\n", mem.FreeHighMemInGB())

fmt.Printf("Avail Mem: %d KB\n", mem.AvailMemInKB())
fmt.Printf("Avail Mem: %.1f MB\n", mem.AvailMemInMB())
fmt.Printf("Avail Mem: %.1f GB\n", mem.AvailMemInGB())

fmt.Printf("Used Mem: %d KB\n", mem.UsedMemInKB())
fmt.Printf("Used Mem: %.1f MB\n", mem.UsedMemInMB())
fmt.Printf("Used Mem: %.1f GB\n", mem.UsedMemInGB())

fmt.Printf("Total Used Mem: %d KB\n", mem.TotalUsedMemInKB())
fmt.Printf("Total Used Mem: %1.f MB\n", mem.TotalUsedMemInMB())
fmt.Printf("Total Used Mem: %.1f GB\n", mem.TotalUsedMemInGB())

fmt.Printf("Total Used: %d KB\n", mem.TotalUsedInKB())
fmt.Printf("Total Used: %.1f MB\n", mem.TotalUsedInMB())
fmt.Printf("Total Used: %.1f GB\n", mem.TotalUsedInGB())

fmt.Printf("Buffer Mem: %d KB\n", mem.BufferMemInKB())
fmt.Printf("Buffer Mem: %.1f MB\n", mem.BufferMemInMB())
fmt.Printf("Buffer Mem: %.1f GB\n", mem.BufferMemInGB())

fmt.Printf("Shared Mem: %d KB\n", mem.SharedMemInKB())
fmt.Printf("Shared Mem: %.1f MB\n", mem.SharedMemInMB())
fmt.Printf("Shared Mem: %.1f GB\n", mem.SharedMemInGB())

fmt.Printf("Total Swap: %d KB\n", mem.TotalSwapInKB())
fmt.Printf("Total Swap: %.1f MB\n", mem.TotalSwapInMB())
fmt.Printf("Total Swap: %.1f GB\n", mem.TotalSwapInGB())

fmt.Printf("Free Swap: %d KB\n", mem.FreeSwapInKB())
fmt.Printf("Free Swap: %.1f MB\n", mem.FreeSwapInMB())
fmt.Printf("Free Swap: %.1f GB\n", mem.FreeSwapInGB())

fmt.Printf("Used Swap: %d KB\n", mem.UsedSwapInKB())
fmt.Printf("Used Swap: %.1f MB\n", mem.UsedSwapInMB())
fmt.Printf("Used Swap: %.1f GB\n", mem.UsedSwapInGB())
```

#### Node Info

```Go
node := sysinfo.Node{}
if err := node.Get(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}

fmt.Printf("Domainname: %s\n", node.DomainName)
fmt.Printf("Machine: %s\n", node.Machine)
fmt.Printf("Nodename: %s\n", node.NodeName)
fmt.Printf("Release: %s\n", node.Release)
fmt.Printf("Sysname: %s\n", node.SysName)
fmt.Printf("Version: %s\n", node.Version)
fmt.Printf("OS Name: %s\n", node.OSName)
```

#### Uptime Info

```Go
uptime := sysinfo.Uptime{}
if err := uptime.Get(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}

fmt.Printf("Uptime: %s\n", uptime)
fmt.Printf("Up since (RFC1123): %s\n", uptime.UpSince())
fmt.Printf("Up since (Custom): %s\n", uptime.UpSinceFormat("%d/%m/%y %T")) // strftime format
```
