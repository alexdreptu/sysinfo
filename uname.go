package sysinfo

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

type Uname struct {
	unix.Utsname
	OSName string // from /etc/os-release
}

func (u *Uname) GetDomainName() string {
	return ConvertInt8ArrayToString(u.Domainname)
}

func (u *Uname) GetMachine() string {
	return ConvertInt8ArrayToString(u.Machine)
}

func (u *Uname) GetNodeName() string {
	return ConvertInt8ArrayToString(u.Nodename)
}

func (u *Uname) GetRelease() string {
	return ConvertInt8ArrayToString(u.Release)
}

func (u *Uname) GetSysName() string {
	return ConvertInt8ArrayToString(u.Sysname)
}

func (u *Uname) GetVersion() string {
	return ConvertInt8ArrayToString(u.Version)
}

// GetOSName is a conformity method
func (u *Uname) GetOSName() string {
	return u.OSName
}

func (u *Uname) Get() error {
	var err error
	utsname := &unix.Utsname{}
	if err := unix.Uname(utsname); err != nil {
		return err
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	if err = enc.Encode(utsname); err != nil {
		return err
	}

	if err = dec.Decode(&u); err != nil {
		return err
	}

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
