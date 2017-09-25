//
// main - start the tcp server
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-09-25 08:27:43
//

package main

import (
    "context"
	"fmt"
    "flag"
    "net"
    "os"
    "path"
    "strings"
    "time"
	"unique"
)

type Server struct {
    Port int
    IdleTimeout time.Duration
}

type Client struct {
    id string
    IdleTimeout time.Duration
    Requests int
    buffer [64]byte
}

func (cli *Client) readRequest(conn net.Conn) (string, error) {
    fmt.Println("read from client...")

    ctx, cancel := context.WithTimeout(context.Background(), cli.IdleTimeout)
    defer cancel()

    var (
        response string
        err error
        ccount int
    )

    complete := make(chan bool, 1)

    go func() {
        buf := cli.buffer[:]
        ccount, err = conn.Read(buf)
        // fmt.Println("ccount:", ccount)
        if err == nil && ccount > 0 {
            response = strings.TrimSpace(string(buf[:ccount]))
        }
        cli.Requests++

        complete <- true
    }()

    select {
    case <-ctx.Done():
        return response, ctx.Err()
    case <-complete:
        return response, err
    }
}

func (cli Client) handleClient(conn net.Conn) {
    defer conn.Close()
    for {
        buf, err := cli.readRequest(conn)
        if err != nil {
            fmt.Printf("connection lost from client: %s; %s\n", cli.id, err)
            break
        } else {
            // parse the request
            id := unique.CreateULID()
            fmt.Printf("client: %s request: %s, response: %s\n", cli.id, buf, id)

            fmt.Fprintf(conn, "%s\n\r", id)
        }
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
        client := Client{ id:unique.CreateTXID(), IdleTimeout:svr.IdleTimeout }
        go client.handleClient(conn)
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

