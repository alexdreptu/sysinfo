package node_test

import (
	"testing"

	. "github.com/alexdreptu/sysinfo/node"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/unix"
)

// test values
const (
	domainName = "localhost"
	machine    = "x86_64"
	nodeNname  = "SplinterProject"
	release    = "5.10-7.arch1-1"
	sysName    = "Linux"
	version    = "#1 SMP PREEMPT Wed, 13 Jan 2021 12:02:01 +0000"
	osName     = "Arch Linux"
)

// mock function
func uname(buf *unix.Utsname) error {
	copy(buf.Domainname[:], []byte(domainName))
	copy(buf.Machine[:], []byte(machine))
	copy(buf.Nodename[:], []byte(nodeNname))
	copy(buf.Release[:], []byte(release))
	copy(buf.Sysname[:], []byte(sysName))
	copy(buf.Version[:], []byte(version))
	return nil
}

func TestNode(t *testing.T) {
	node := &Node{}
	node.F = uname
	node.OSName = osName

	assert.NoError(t, node.Fetch())
	assert.Equal(t, domainName, node.DomainName)
	assert.Equal(t, machine, node.Machine)
	assert.Equal(t, nodeNname, node.NodeName)
	assert.Equal(t, release, node.Release)
	assert.Equal(t, sysName, node.SysName)
	assert.Equal(t, version, node.Version)
	assert.Equal(t, osName, node.OSName)
}
