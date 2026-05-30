//go:build darwin || freebsd || netbsd || openbsd || windows

package utils

import "syscall"

func setProxyListenerSocketOptions(_ string, _ string, _ syscall.RawConn) error {
	return nil
}
