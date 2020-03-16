//
// main - then standard cli
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-07 12:08:51
//

package main

import (
	"flag"
	"fmt"
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
	size := flag.Int("size", 24, "the number of random bytes to user (only for bytes)")
	ulid := flag.Bool("ulid", false, "generate a ulid")
	uuid := flag.Bool("uuid", false, "generate a uuid")
	guid := flag.Bool("guid", false, "generate a guid")
	tsid := flag.Bool("tsid", false, "generate a tsid")
	txid := flag.Bool("txid", false, "generate a txid")
	cuid := flag.Bool("cuid", false, "generate a cuid")
	xuid := flag.Bool("xuid", false, "generate a xuid")
	bytes := flag.Bool("bytes", false, "generate a 48 character byte stream")

	flag.Parse()

	if *bytes == true {
		if buf, err := unique.RandomBytes(*size); err == nil {
			fmt.Printf("%x\n", buf)
			return
		}
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

	if *cuid == true || strings.HasSuffix(nm, "cuid") {
		fmt.Println(unique.CreateCUID())
		return
	}

	if *xuid == true || strings.HasSuffix(nm, "xuid") {
		fmt.Println(unique.CreateXUID())
		return
	}

	if *vers == true || len(os.Args) == 1 {
		showVersion()
		return
	}
}

func main() {
	parseArgs()
}
