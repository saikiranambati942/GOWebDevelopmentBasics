package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\n Hello from TCP Server")
		fmt.Fprintln(conn, "\n I hope you are good")
		fmt.Fprintln(conn,"\n My name is Sai:")
		conn.Close()
	}
}
