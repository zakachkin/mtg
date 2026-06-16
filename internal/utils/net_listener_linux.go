package utils

import (
	"syscall"

	"golang.org/x/sys/unix"
)

const proxyListenerMSS = 128

func setProxyListenerSocketOptions(_, _ string, conn syscall.RawConn) error {
	var sockErr error
	ctrlErr := conn.Control(func(fd uintptr) {
		sockErr = unix.SetsockoptInt(int(fd), unix.IPPROTO_TCP, unix.TCP_MAXSEG, proxyListenerMSS)
	})
	if ctrlErr != nil {
		return ctrlErr //nolint: wrapcheck
	}
	return sockErr //nolint: wrapcheck
}
