package main

import (
	"io/ioutil"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}

	log.Println("Listening at :3000...")
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		log.Println("Accepting...")
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		defer c.Close()
		bs, err := ioutil.ReadAll(c)
		if err != nil {
			panic(err)
		}
		log.Printf("> %s", string(bs))
		if _, err := c.Write(bs); err != nil {
			panic(err)
		}
		if err := c.Close(); err != nil {
			panic(err)
		}
	}
}
