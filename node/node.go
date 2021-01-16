package node

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

type fetchFunc func(buf *unix.Utsname) error

type Node struct {
	DomainName string
	Machine    string
	NodeName   string
	Release    string
	SysName    string
	Version    string
	OSName     string // from /etc/os-release
	F          fetchFunc
}

// Fetch updates the Node struct woth new values
func (n *Node) Fetch() error {
	utsname := unix.Utsname{}

	var osName string
	var err error

	if n.F == nil {
		n.F = unix.Uname
		osName, err = readOSName()
		if err != nil {
			return err
		}
		n.OSName = osName
	}

	if err = n.F(&utsname); err != nil {
		return err
	}

	n.DomainName = string(utsname.Domainname[:bytes.IndexByte(utsname.Domainname[:], 0)])
	n.Machine = string(utsname.Machine[:bytes.IndexByte(utsname.Machine[:], 0)])
	n.NodeName = string(utsname.Nodename[:bytes.IndexByte(utsname.Nodename[:], 0)])
	n.Release = string(utsname.Release[:bytes.IndexByte(utsname.Release[:], 0)])
	n.SysName = string(utsname.Sysname[:bytes.IndexByte(utsname.Sysname[:], 0)])
	n.Version = string(utsname.Version[:bytes.IndexByte(utsname.Version[:], 0)])

	return nil
}

func readOSName() (string, error) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var osName string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "NAME") {
			osName = strings.Trim(strings.Split(line, "=")[1], "\"")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return osName, nil
}
