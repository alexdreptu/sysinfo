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
	if n.F == nil {
		n.F = unix.Uname
	}

	utsname := unix.Utsname{}
	if err := n.F(&utsname); err != nil {
		return err
	}

	if n.OSName == "" {
		if err := n.readOSName(); err != nil {
			return err
		}
	}

	n.DomainName = string(utsname.Domainname[:bytes.IndexByte(utsname.Domainname[:], 0)])
	n.Machine = string(utsname.Machine[:bytes.IndexByte(utsname.Machine[:], 0)])
	n.NodeName = string(utsname.Nodename[:bytes.IndexByte(utsname.Nodename[:], 0)])
	n.Release = string(utsname.Release[:bytes.IndexByte(utsname.Release[:], 0)])
	n.SysName = string(utsname.Sysname[:bytes.IndexByte(utsname.Sysname[:], 0)])
	n.Version = string(utsname.Version[:bytes.IndexByte(utsname.Version[:], 0)])

	return nil
}

func (n *Node) readOSName() error {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "NAME") {
			n.OSName = strings.Trim(strings.Split(line, "=")[1], "\"")
			break
		}
	}

	return scanner.Err()
}

func New() (*Node, error) {
	node := &Node{}

	if err := node.Fetch(); err != nil {
		return &Node{}, err
	}

	return node, nil
}
