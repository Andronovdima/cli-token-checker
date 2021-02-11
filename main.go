package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
)

func main() {
	flag.Parse()
	flags := flag.Args()

	client := Client{}
	err := client.ParseArgs(flags)
	if err != nil {
		log.Println("Error parse arguments")
		return
	}

	client.ReqID = rand.Int31()
	client.Conn = &TCPConn{}

	response, err := client.Request()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(response.Output())

	return
}
