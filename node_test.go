package sysinfo

import "testing"

func TestNodeInfo(t *testing.T) {
	node := NewNode()
	if err := node.Get(); err != nil {
		t.Fatal(err)
	}

	t.Logf("Domainname: %s\n", node.DomainName)
	t.Logf("Machine: %s\n", node.Machine)
	t.Logf("Nodename: %s\n", node.NodeName)
	t.Logf("Release: %s\n", node.Release)
	t.Logf("Sysname: %s\n", node.SysName)
	t.Logf("Version: %s\n", node.Version)
	t.Logf("OS Name: %s\n", node.OSName)
}
