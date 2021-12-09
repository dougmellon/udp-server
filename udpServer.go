package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	s, err := net.ResolveUDPAddr("udp4", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("UDP server available on port 8081")
	}

	defer conn.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	// send setup
	connect := "127.0.0.1:8080"
	raddr, err := net.ResolveUDPAddr("udp4", connect)
	connTwo, err := net.DialUDP("udp4", nil, raddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connTwo.Close()

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server")
			return
		}

		data := []byte(buffer[0 : n-1])
		fmt.Printf("data: %s\n", string(data))
		_, err = conn.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = connTwo.Write(data)
	}
}
