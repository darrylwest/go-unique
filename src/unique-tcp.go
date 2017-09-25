//
// main - start the tcp server
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-09-25 08:27:43
//

package main

import (
	"fmt"
    "flag"
    "net"
    "os"
    "path"
    "time"
	"unique"
)

type Server struct {
    Port int
    IdleTimeout time.Duration
}

func readRequest(conn net.Conn) ([]byte, error) {
    var (
        response []byte
        err error
        ccount int
    )

    buf := make([]byte, 32)
    ccount, err = conn.Read(buf)
    if err == nil && ccount > 0{
        response = buf[:ccount]
    }
    return response, err
}

func (svr Server) handleClient(conn net.Conn) {
    defer conn.Close()
    for {
        buf, err := readRequest(conn)
        if err != nil {
            fmt.Printf("connection logs: %s\n", err)
            break
        }

        fmt.Println(buf)

        // parse the request
        id := unique.CreateULID()
        fmt.Fprintf(conn, id)
    }
}

func main() {
    svr := parseArgs()
    if svr == nil {
        return 
    }

    // start the server
    host := fmt.Sprintf("0.0.0.0:%d", svr.Port)
    ss, err := net.Listen("tcp", host)
    if err != nil {
        fmt.Printf("error opening host %s...\n", host)
        os.Exit(1)
    }

    fmt.Printf("listening on host: %s\n", host)

    defer ss.Close()
    for {
        conn, err := ss.Accept()
        if err != nil {
            fmt.Println("Accept error: ", err.Error())
        }

        // create a client struct and add to the list
        go svr.handleClient(conn)
    }
}

func showVersion() {
    fmt.Printf("%s Version: %s\n", path.Base(os.Args[0]), unique.Version())
}

func parseArgs() *Server {
    svr := Server{ Port: 3001, IdleTimeout: 5 * time.Minute }

    vers := flag.Bool("version", false, "show the version and exit")
    flag.Parse();

    // show the version
    if *vers == true {
        return nil
    }

    // show the port and idle timeout
    return &svr
}

