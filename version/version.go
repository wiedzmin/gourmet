package version

import "fmt"

const major = 0
const minor = 0
const patch = 1
const relInfo = "master"

const (
	// Description provides tidy naming service
	Description = "Fancy bookmarks manager"
	// Usage describes what aims it will be used for
	Usage = "Bookmarks maintainance"
)

// Version returns version info for Gourmet
func Version() string {
	return fmt.Sprintf("v%d.%d.%d#%s", major, minor, patch, relInfo)
}
