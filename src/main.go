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
    "strings"
	"unique"
)

func showVersion() {
    fmt.Printf("%s Version: %s\n", path.Base(os.Args[0]), unique.Version())
}

func parseArgs() {
    vers := flag.Bool("version", false, "show the version and exit")
    ulid := flag.Bool("ulid", false, "generate a ulid")
    uuid := flag.Bool("uuid", true, "generate a uuid")
    guid := flag.Bool("guid", false, "generate a guid")
    tsid := flag.Bool("tsid", false, "generate a tsid")
    txid := flag.Bool("txid", false, "generate a txid")
    bytes := flag.Bool("bytes", false, "generate a 48 character byte stream")

    flag.Parse()

    if *vers == true {
        os.Exit(0)
    }

    nm := os.Args[0]

    if *ulid == true || strings.HasSuffix(nm, "ulid") {
        fmt.Println(unique.CreateULID())
        return
    }

    if *uuid == true || strings.HasSuffix(nm, "uuid") {
        fmt.Println(unique.CreateUUID())
        return
    }

    if *guid == true || strings.HasSuffix(nm, "guid") {
        fmt.Println(unique.CreateGUID())
        return
    }

    if *tsid == true || strings.HasSuffix(nm, "tsid") {
        fmt.Println(unique.CreateTSID())
        return
    }

    if *txid == true || strings.HasSuffix(nm, "txid") {
        fmt.Println(unique.CreateTXID())
        return
    }

    if *bytes == true {
        if buf, err := unique.RandomBytes(24); err == nil {
            str := fmt.Sprintf("%x", buf)
            fmt.Printf("%s (%d)\n", str, len(str))
            return
        }
    }

    if *vers == true {
        showVersion()
    }
}

func main() {
    parseArgs()
}
