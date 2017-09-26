package main

import (
	"fmt"
	"net"
	"os"
    "strings"
	"time"
)

func main() {
	port := 3001
	host := fmt.Sprintf(":%d", port)
	fmt.Println("dailing: ", host)

	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("error connecting to host ", host)
		os.Exit(1)
	}

    cmds := strings.Split("uuid ulid guid tsid txid", " ")
	defer conn.Close()
	count := 0

    buf := make([]byte, 512)

	for {
        cmd := cmds[(count % len(cmds))]
		count++
        if count % 10 == 0 {
            time.Sleep(5 * time.Second)
        }

		text := fmt.Sprintf("%s\n\r", cmd)
		_, err := fmt.Fprintf(conn, text)
		if err != nil {
			fmt.Println("lost connection...")
			return
		}

		fmt.Printf("sent: %s", text)

        n, err := conn.Read(buf)
        if err != nil {
			fmt.Println("lost connection...")
			return
        }

        fmt.Printf("recd: %s", buf[:n]);
	}
}
