package network

// Network I/O buffer size
var IOSize uint64 = 65536

// Number of minutes of profile data to capture
// XXX Where should this be set? Config file??
const PROFILE_MINS = 30

// Interface for all proxies
type Proxy interface {
	Start()
	Stop()
}
