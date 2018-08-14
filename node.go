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

func (u *Node) Get() (err error) {
	utsname := unix.Utsname{}
	if err = unix.Uname(&utsname); err != nil {
		return err
	}

	u.DomainName = string(utsname.Domainname[:])
	u.Machine = string(utsname.Machine[:])
	u.NodeName = string(utsname.Nodename[:])
	u.Release = string(utsname.Release[:])
	u.SysName = string(utsname.Sysname[:])
	u.Version = string(utsname.Version[:])

	if u.OSName, err = readOSName(); err != nil {
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
