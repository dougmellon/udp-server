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

func sendToBoard() {

}

func main() {
	s, err := net.ResolveUDPAddr("udp4", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("UDP available on port 8080")
	}

	defer conn.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	// connect to local Python client
	connect := "174.16.236.178:6789"

	raddr, err := net.ResolveUDPAddr("udp4", connect)
	conn2, err := net.DialUDP("udp4", nil, raddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn2.Close()

	for {
		n, _, err := conn.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server")
			return
		}

		data := []byte(buffer[0 : n-1])
		fmt.Printf("data: %s\n", string(data))
		_, err = conn.WriteToUDP(data, raddr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
