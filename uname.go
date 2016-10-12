package sysinfo

import (
	"bytes"
	"encoding/gob"

	"golang.org/x/sys/unix"
)

type Uname unix.Utsname

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

	return nil
}
