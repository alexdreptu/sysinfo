package sysinfo

import "testing"

func TestGetDomainName(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Domainname: %s\n", uname.GetDomainName())
}

func TestGetMachine(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Machine: %s\n", uname.GetMachine())
}

func TestGetNodeName(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Nodename: %s\n", uname.GetNodeName())
}

func TestGetRelease(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Release: %s\n", uname.GetRelease())
}

func TestGetSysName(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Sysname: %s\n", uname.GetSysName())
}

func TestGetVersion(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Version: %s\n", uname.GetVersion())
}

func TestGetOSName(t *testing.T) {
	uname := &Uname{}
	if err := uname.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("OS Name: %s\n", uname.GetOSName())
}
