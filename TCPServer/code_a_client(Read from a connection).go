package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main(){
	conn,err:=net.Dial("tcp",":8080")
	if err!=nil{
		log.Panic(err)
	}
	defer conn.Close()

	b,err:=ioutil.ReadAll(conn)
	if err!=nil{
		log.Println(err)
	}
	fmt.Println(string(b))
}
