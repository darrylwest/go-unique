//
// simple verion file
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-07 12:20:40
//

package unique

import "fmt"

const (
	major = 20
	minor = 03
	patch = 15
)

// Version - return the version number as a single string
func Version() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
