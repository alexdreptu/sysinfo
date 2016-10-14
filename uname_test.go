package sysinfo

import "testing"

func TestDomainName(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Domainname: %s\n", uname.DomainName)
}

func TestMachine(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Machine: %s\n", uname.Machine)
}

func TestNodeName(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Nodename: %s\n", uname.NodeName)
}

func TestRelease(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Release: %s\n", uname.Release)
}

func TestSysName(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Sysname: %s\n", uname.SysName)
}

func TestVersion(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Version: %s\n", uname.Version)
}

func TestOSName(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("OS Name: %s\n", uname.OSName)
}
