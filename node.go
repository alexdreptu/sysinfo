package sysinfo

import (
	"bufio"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

type Node struct {
	DomainName string
	Machine    string
	NodeName   string
	Release    string
	SysName    string
	Version    string
	OSName     string // from /etc/os-release
}

func (u *Node) Get() error {
	utsname := unix.Utsname{}
	if err := unix.Uname(&utsname); err != nil {
		return err
	}

	u.DomainName = convertInt8ArrayToString(utsname.Domainname)
	u.Machine = convertInt8ArrayToString(utsname.Machine)
	u.NodeName = convertInt8ArrayToString(utsname.Nodename)
	u.Release = convertInt8ArrayToString(utsname.Release)
	u.SysName = convertInt8ArrayToString(utsname.Sysname)
	u.Version = convertInt8ArrayToString(utsname.Version)

	var err error
	u.OSName, err = readOSName()
	if err != nil {
		return err
	}

	return nil
}

func readOSName() (string, error) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = scanner.Text()
		if strings.Contains(content, "NAME") {
			content = strings.Trim(strings.Split(content, "=")[1], "\"")
			break
		}
	}

	return content, nil
}
