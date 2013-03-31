package main

import irc "github.com/fluffle/goirc/client"
import "fmt"

func main() {
    c := irc.SimpleClient("goilertest")
    c.SSL = true

    // join channel on connect
    c.AddHandler(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) { conn.Join("#nodrama.de") })
    // emit a signal on disconnect
    quit := make(chan bool)
    c.AddHandler(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) { quit <- true })

    err := c.Connect("irc.freenode.net")
    if err != nil {
        fmt.Printf("Connection error: %s\n", err)
    }

    // Wait for disconnect
    <-quit
}
