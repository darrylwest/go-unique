//
// main -
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-07 12:08:51
//

package main

import (
	"fmt"
    "flag"
    "os"
    "path"
	"unique"
)

func showVersion() {
    fmt.Printf("%s Version: %s\n", path.Base(os.Args[0]), unique.Version())
}

func parseArgs() {
    vers := flag.Bool("version", false, "show the version and exit")
    ulid := flag.Bool("ulid", false, "generate a ulid")
    uuid := flag.Bool("uuid", false, "generate a uuid")
    guid := flag.Bool("guid", false, "generate a guid")
    tsid := flag.Bool("tsid", false, "generate a tsid")
    txid := flag.Bool("txid", false, "generate a txid")
    bytes := flag.Bool("bytes", false, "generate a 48 character byte stream")

    flag.Parse()

    showVersion()
    if *vers == true {
        os.Exit(0)
    }

    if *ulid == true {
        fmt.Println(unique.CreateULID())
    }

    if *uuid == true {
        fmt.Println(unique.CreateUUID())
    }

    if *guid == true {
        fmt.Println(unique.CreateGUID())
    }

    if *tsid == true {
        fmt.Println(unique.CreateTSID())
    }

    if *txid == true {
        fmt.Println(unique.CreateTXID())
    }

    if *bytes == true {
        if buf, err := unique.RandomBytes(24); err == nil {
            str := fmt.Sprintf("%x", buf)
            fmt.Printf("%s (%d)\n", str, len(str))
        }
    }
}

func main() {
    parseArgs()
}
