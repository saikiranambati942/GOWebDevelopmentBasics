package main

import (
	"io"
	"log"
	"net"
)

func main(){
	conn,err:=net.Dial("tcp",":8080")
	if err!=nil{
		log.Panic(err)
	}
	defer conn.Close()
	io.WriteString(conn,"I dailed you:")
}
