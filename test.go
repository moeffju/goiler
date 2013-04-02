package main

import (
	"fmt"
	irc "github.com/fluffle/goirc/client"
)

func main() {
	c := irc.SimpleClient("goilertest", "goiler", "goiler_in ist goiler als noiler")
	c.SSL = true
	c.EnableStateTracking()

	// join channel on connect
	c.AddHandler(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) { conn.Join("#furanzentest") })
	// emit a signal on disconnect
	quit := make(chan bool)
	c.AddHandler(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) { quit <- true })

	if err := c.Connect("irc.freenode.net"); err != nil {
		fmt.Printf("Connection error: %s\n", err)
	}

	// Wait for disconnect
	<-quit
}
