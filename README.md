[![Build Status](https://github.com/alexdreptu/sysinfo/workflows/Build/badge.svg)](https://github.com/alexdreptu/sysinfo/actions)
[![Coverage Status](https://coveralls.io/repos/github/alexdreptu/sysinfo/badge.svg?branch=master)](https://coveralls.io/github/alexdreptu/sysinfo?branch=master)
[![GoDev](https://pkg.go.dev/badge/github.com/alexdreptu/sysinfo)](https://pkg.go.dev/github.com/alexdreptu/sysinfo)
[![Platform](https://img.shields.io/badge/platform-Linux-5272b4.svg)](https://www.linuxfoundation.org/)
[![License](https://img.shields.io/badge/license-MIT-5272b4.svg)](https://github.com/alexdreptu/sysinfo/blob/master/LICENSE)

#### This module is not meant to be a complete solution, rather it's meant to get just the minimum system information required by [archey-go](https://github.com/alexdreptu/archey-go). If you're looking for something more complete, then you might want to take a look at [gopsutil](https://github.com/shirou/gopsutil)

### Examples

#### CPU Info

```Go
package main

import (
    "fmt"

    "github.com/alexdreptu/sysinfo/cpu"
)
```

```Go
cpu := &cpu.CPU{}
if err := cpu.Fetch(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}
```

or

```Go
cpu, err := cpu.New()
if err != nil {
    fmt.Printf(os.Stderr, err)
    os.Exit(1)
}
```

```Go
fmt.Printf("Name: %s\n", cpu.Name)
fmt.Printf("Model: %d\n", cpu.Model)
fmt.Printf("Vendor: %s\n", cpu.Vendor)
fmt.Printf("Family: %d\n", cpu.Family)
fmt.Printf("Stepping: %d\n", cpu.Stepping)
fmt.Printf("Cores: %d\n", cpu.Cores)
fmt.Printf("Cache Size: %d KiB\n", cpu.CacheSize)
fmt.Printf("Cache Size: %.2f MiB\n", cpu.CacheSizeInMebibytes())
fmt.Printf("Min Frequency: %d KHz\n", cpu.MinFreq)
fmt.Printf("Max Frequency: %d KHz\n", cpu.MaxFreq)
fmt.Printf("Min Frequency: %.2f MHz\n", cpu.MinFreqInMegahertz())
fmt.Printf("Min Frequency: %.1f GHz\n", cpu.MinFreqInGigahertz())
fmt.Printf("Max Frequency: %.2f MHz\n", cpu.MaxFreqInMegahertz())
fmt.Printf("Max Frequency: %.1f GHz\n", cpu.MaxFreqInGigahertz())
fmt.Printf("Flags %s\n", cpu.FlagsAsString())
```

#### Filesystem Info

```Go
package main

import (
    "fmt"

    "github.com/alexdreptu/sysinfo/fs"
)
```

```Go
rootfs := &fs.FS{Path: "/"}
if err := rootfs.Fetch(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}
```

or

```Go
rootfs, err := fs.New("/")
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}
```

```Go
fmt.Printf("Total Space on Root: %.2f KiB\n", rootfs.TotalSpaceInKibibytes())
fmt.Printf("Total Space on Root: %.2f MiB\n", rootfs.TotalSpaceInMebibytes())
fmt.Printf("Total Space on Root: %.2f GiB\n", rootfs.TotalSpaceInGibibytes())

fmt.Printf("Free Space on Root: %.2f KiB\n", rootfs.FreeSpaceInKibibytes())
fmt.Printf("Free Space on Root: %.2f MiB\n", rootfs.FreeSpaceInMebibytes())
fmt.Printf("Free Space on Root: %.2f GiB\n", rootfs.FreeSpaceInGibibytes())

fmt.Printf("Used Space on Root: %.2f KiB\n", rootfs.UsedSpaceInKibibytes())
fmt.Printf("Used Space on Root: %.2f MiB\n", rootfs.UsedSpaceInMebibytes())
fmt.Printf("Used Space On Root: %.2f GiB\n", rootfs.UsedSpaceInGibibytes())
```

#### Memory Info

```Go
package main

import (
    "fmt"

    "github.com/alexdreptu/sysinfo/mem"
)
```

```Go
mem := &mem.Mem{}
if err := mem.Fetch(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}
```

or

```Go
mem, err := mem.New()
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}
```

```Go
fmt.Printf("Total Mem: %.2f KiB\n", mem.TotalMemInKibibytes())
fmt.Printf("Total Mem: %.2f MiB\n", mem.TotalMemInMebibytes())
fmt.Printf("Total Mem: %.2f GiB\n", mem.TotalMemInGibibytes())

fmt.Printf("Total High Mem: %.2f KiB\n", mem.TotalHighMemInKibibytes())
fmt.Printf("Total High Mem: %.2f MiB\n", mem.TotalHighMemInMebibytes())
fmt.Printf("Total High Mem: %.2f GiB\n", mem.TotalHighMemInGibibytes())

fmt.Printf("Free Mem: %.2f KiB\n", mem.FreeMemInKibibytes())
fmt.Printf("Free Mem: %.2f MiB\n", mem.FreeMemInMebibytes())
fmt.Printf("Free Mem: %.2f GiB\n", mem.FreeMemInGibibytes())

fmt.Printf("Free High Mem: %.2f KiB\n", mem.FreeHighMemInKibibytes())
fmt.Printf("Free High Mem: %.2f MiB\n", mem.FreeHighMemInMebibytes())
fmt.Printf("Free High Mem: %.2f GiB\n", mem.FreeHighMemInGibibytes())

fmt.Printf("Avail Mem: %.2f KiB\n", mem.AvailMemInKibibytes())
fmt.Printf("Avail Mem: %.2f MiB\n", mem.AvailMemInMebibytes())
fmt.Printf("Avail Mem: %.2f GiB\n", mem.AvailMemInGibibytes())

fmt.Printf("Used Mem: %.2f KiB\n", mem.UsedMemInKibibytes())
fmt.Printf("Used Mem: %.2f MiB\n", mem.UsedMemInMebibytes())
fmt.Printf("Used Mem: %.2f GiB\n", mem.UsedMemInGibibytes())

fmt.Printf("Total Used Mem: %.2f KiB\n", mem.TotalUsedMemInKibibytes())
fmt.Printf("Total Used Mem: %.2f MiB\n", mem.TotalUsedMemInMebibytes())
fmt.Printf("Total Used Mem: %.2f GiB\n", mem.TotalUsedMemInGibibytes())

fmt.Printf("Total Used: %.2f KiB\n", mem.TotalUsedInKibibytes())
fmt.Printf("Total Used: %.2f MiB\n", mem.TotalUsedInMebibytes())
fmt.Printf("Total Used: %.2f GiB\n", mem.TotalUsedInGibibytes())

fmt.Printf("Buffer Mem: %.2f KiB\n", mem.BufferMemInKibibytes())
fmt.Printf("Buffer Mem: %.2f MiB\n", mem.BufferMemInMebibytes())
fmt.Printf("Buffer Mem: %.2f GiB\n", mem.BufferMemInGibibytes())

fmt.Printf("Shared Mem: %.2f KiB\n", mem.SharedMemInKibibytes())
fmt.Printf("Shared Mem: %.2f MiB\n", mem.SharedMemInMebibytes())
fmt.Printf("Shared Mem: %.2f GiB\n", mem.SharedMemInGibibytes())

fmt.Printf("Total Swap: %.2f KiB\n", mem.TotalSwapInKibibytes())
fmt.Printf("Total Swap: %.2f MiB\n", mem.TotalSwapInMebibytes())
fmt.Printf("Total Swap: %.2f GiB\n", mem.TotalSwapInGibibytes())

fmt.Printf("Free Swap: %.2f KiB\n", mem.FreeSwapInKibibytes())
fmt.Printf("Free Swap: %.2f MiB\n", mem.FreeSwapInMebibytes())
fmt.Printf("Free Swap: %.2f GiB\n", mem.FreeSwapInGibibytes())

fmt.Printf("Used Swap: %.2f KiB\n", mem.UsedSwapInKibibytes())
fmt.Printf("Used Swap: %.2f MiB\n", mem.UsedSwapInMebibytes())
fmt.Printf("Used Swap: %.2f GiB\n", mem.UsedSwapInGibibytes())
```

#### Node Info

```Go
package main

import (
    "fmt"

    "github.com/alexdreptu/sysinfo/node"
)
```

```Go
node := &node.Node{}
if err := node.Fetch(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}
```

or

```Go
node, err := node.New()
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}
```

```Go
fmt.Printf("Domainname: %s\n", node.DomainName)
fmt.Printf("Machine: %s\n", node.Machine)
fmt.Printf("Nodename: %s\n", node.NodeName)
fmt.Printf("OS Name: %s\n", node.OSName)
fmt.Printf("Release: %s\n", node.Release)
fmt.Printf("Sysname: %s\n", node.SysName)
fmt.Printf("Version: %s\n", node.Version)
```

#### Uptime Info

```Go
package main

import (
    "fmt"

    "github.com/alexdreptu/sysinfo/uptime"
)
```

```Go
uptime := &uptime.Uptime{}
if err := uptime.Fetch(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}
```

or

```Go
uptime, err := uptime.New()
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
)
```

```Go
fmt.Printf("Uptime: %s\n", uptime)
fmt.Printf("Up since (RFC1123): %s\n", uptime.Since())
fmt.Printf("Up since (Custom): %s\n", uptime.SinceFormat("%d/%m/%y %T")) // strftime format
```

#### Convert units

```Go
package main

import (
    "fmt"

    "github.com/alexdreptu/sysinfo/convert"
)

func main() {
     // convert bytes
    bytes := 8267812044.8
    fmt.Printf("%.2f Bytes -> %.2f Kibibytes\n",
        bytes, (convert.Unit(bytes) * convert.Byte).Kibibytes())
    fmt.Printf("%.2f Bytes -> %.2f Mebibytes\n",
        bytes, (convert.Unit(bytes) * convert.Byte).Mebibytes())
    fmt.Printf("%.2f Bytes -> %.2f Gibibytes\n",
        bytes, (convert.Unit(bytes) * convert.Byte).Gibibytes())

    // convert kibibytes
    kibibytes := 8074035.2
    fmt.Printf("%.2f Kibibytes -> %.2f Bytes\n",
        kibibytes, (convert.Unit(kibibytes) * convert.Kibibyte).Bytes())
    fmt.Printf("%.2f Kibibytes -> %.2f Mebibytes\n",
        kibibytes, (convert.Unit(kibibytes) * convert.Kibibyte).Mebibytes())
    fmt.Printf("%.2f Kibibytes -> %.2f Gibibytes\n",
        kibibytes, (convert.Unit(kibibytes) * convert.Kibibyte).Gibibytes())

    // convert mebibytes
    mebibytes := 7884.8
    fmt.Printf("%.2f Mebibytes -> %.2f Bytes\n",
        mebibytes, (convert.Unit(mebibytes) * convert.Mebibyte).Bytes())
    fmt.Printf("%.2f Mebibytes -> %.2f Kibibytes\n",
        mebibytes, (convert.Unit(mebibytes) * convert.Mebibyte).Kibibytes())
    fmt.Printf("%.2f Mebibytes -> %.2f Gibibytes\n",
        mebibytes, (convert.Unit(mebibytes) * convert.Mebibyte).Gibibytes())

    // convert gibibytes
    gibibytes := 7.7
    fmt.Printf("%.2f Gibibytes -> %.2f Bytes\n",
        gibibytes, (convert.Unit(gibibytes) * convert.Gibibyte).Bytes())
    fmt.Printf("%.2f Gibibytes -> %.2f Kibibytes\n",
        gibibytes, (convert.Unit(gibibytes) * convert.Gibibyte).Kibibytes())
    fmt.Printf("%.2f Gibibytes -> %.2f Mebibytes\n",
        gibibytes, (convert.Unit(gibibytes) * convert.Gibibyte).Mebibytes())
    
    // convert kilohertz
    kilohertz := 3900000
    fmt.Printf("%v Kilohertz -> %v Megahertz\n",
        kilohertz, (convert.Unit(kilohertz) * convert.Kilohertz).Megahertz())
    fmt.Printf("%v Kilohertz -> %v Gigahertz\n",
        kilohertz, (convert.Unit(kilohertz) * convert.Kilohertz).Gigahertz())
}
```
