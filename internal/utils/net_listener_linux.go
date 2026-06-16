package utils

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/unix"
)

const proxyListenerMSS = 128

func setProxyListenerSocketOptions(network, address string, conn syscall.RawConn) error {
	var sockErr error

	if err := conn.Control(func(fd uintptr) {
		sockErr = unix.SetsockoptInt(
			int(fd),
			unix.IPPROTO_TCP,
			unix.TCP_MAXSEG,
			proxyListenerMSS,
		