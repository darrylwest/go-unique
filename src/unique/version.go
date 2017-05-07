//
// simple verion file
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-07 12:20:40
//

package unique

import "fmt"

const (
	major = 1
	minor = 0
	patch = 0
)

// Version - return the version number as a single string
func Version() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
