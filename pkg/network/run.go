package network

import (
	"runtime"
	"syscall"
)

func RunNetwork() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	// Make sure any files we create are readable ONLY by us
	syscall.Umask(0077)
}
