package main

import (
        "net"
        "log"
        "fmt"
        "strings"
)

func handleConn(c net.Conn) {
        msg := make([]byte, 4096)
        fmt.Println("Receiving connection from", c.RemoteAddr())
        fmt.Fprintf(c, "Hello, %v\n", c.RemoteAddr())

        for {
                length, err := c.Read(msg)
                smsg := strings.Replace(string(msg), "\x00", "", -1)
                if err != nil {
                        log.Println("Got", err, "from", c.RemoteAddr())
                        c.Close()
                        break
                }
                if length == 1 && msg[0] == '\n' {
                        fmt.Fprintf(c, "Would you like to quit? [y/n]\n")
                        _, err := c.Read(msg)
                        if msg[0] == 'y' {
                            fmt.Println(c.RemoteAddr(), "is quitting...")
                            c.Close()
                            break
                        }
                        if err != nil {
                            log.Println("Got", err, "from", c.RemoteAddr())
                            c.Close()
                            break
                        }
                }
                msg = []byte(smsg[:len(smsg)-1])
                fmt.Println(c.RemoteAddr(),"says:", string(msg), ", message length", length, "bytes")
                fmt.Fprintf(c,"%v says: %s", c.RemoteAddr(), string(msg))
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

