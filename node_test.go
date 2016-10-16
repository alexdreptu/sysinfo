package sysinfo

import "testing"

func TestDomainName(t *testing.T) {
	node := NewNode()
	if err := node.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Domainname: %s\n", node.DomainName)
}

func TestMachine(t *testing.T) {
	node := NewNode()
	if err := node.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Machine: %s\n", node.Machine)
}

func TestNodeName(t *testing.T) {
	node := NewNode()
	if err := node.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Nodename: %s\n", node.NodeName)
}

func TestRelease(t *testing.T) {
	node := NewNode()
	if err := node.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Release: %s\n", node.Release)
}

func TestSysName(t *testing.T) {
	node := NewNode()
	if err := node.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Sysname: %s\n", node.SysName)
}

func TestVersion(t *testing.T) {
	node := NewNode()
	if err := node.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("Version: %s\n", node.Version)
}

func TestOSName(t *testing.T) {
	node := NewNode()
	if err := node.Get(); err != nil {
		t.Fatal(err)
	}
	t.Logf("OS Name: %s\n", node.OSName)
}
