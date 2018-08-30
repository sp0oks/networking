package main

import (
        "net"
        "log"
        "fmt"
)

func handleConn(c net.Conn) {
        msg := make([]byte, 4096)
        fmt.Println("Receiving connection from", c.RemoteAddr())
        fmt.Fprintf(c, "Hello, %v\n", c.RemoteAddr())

        for {
                length, err := c.Read(msg)
                if err != nil {
                        log.Println("Got", err, "from", c.RemoteAddr())
                        c.Close()
                        break
                }
                if length == 1 && msg[0] == '\n' {
                        fmt.Println(c.RemoteAddr(), "is quitting...")
                        c.Close()
                        break
                }
                fmt.Println(c.RemoteAddr(),"says:", string(msg))
                fmt.Fprintf(c, "Got %d bytes\n", length)
        }
}

func main() {
        fmt.Println("Starting server...")
        l, err := net.Listen("tcp", ":1999")
        if err != nil {
                log.Fatal(err)
        }
        defer l.Close()
        for {
                conn, err := l.Accept()
                if err != nil {
                        log.Fatal(err)
                }
                go handleConn(conn)
        }
}

