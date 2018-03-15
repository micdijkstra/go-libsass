package libsass

import "github.com/micdijkstra/go-libsass/libs"

// Version reports libsass version information
func Version() string {
	return libs.Version()
}
