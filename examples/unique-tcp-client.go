package main

import (
	"fmt"
	"net"
	"os"
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

	defer conn.Close()
	count := 0

    buf := make([]byte, 512)

	for {
		count++
        if count % 10 == 0 {
            time.Sleep(5 * time.Second)
        }

		text := "ulid\n\r"
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
