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
    fmt.Println("unique, Version:", unique.Version())
	fmt.Println("---------------------------------------------")

	fmt.Println(unique.CreateULID())
	fmt.Println(unique.CreateUUID())
	fmt.Println(unique.CreateGUID())
	fmt.Println(unique.CreateTSID())
	fmt.Println(unique.CreateTXID())

	if buf, err := unique.RandomBytes(24); err == nil {
		str := fmt.Sprintf("%x", buf)
		fmt.Printf("%s (%d)\n", str, len(str))
	}
}
