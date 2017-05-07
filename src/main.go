//
// main -
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-07 12:08:51
//

package main

import (
	"fmt"
	"unique"
)

func main() {
	fmt.Println("this is really unique, no?")

	fmt.Println(unique.CreateULID())
}
