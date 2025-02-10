package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	sendCommand("*3\r\n$3\r\nSET\r\n$5\r\nmyVar\r\n$15\r\nHello, World!!!\r\n")

	sendCommand("*2\r\n$4\r\nINCR\r\n$5\r\ncount\r\n")

	sendCommand("*2\r\n$3\r\nGET\r\n$5\r\ncount\r\n")

	sendCommand("*3\r\n$6\r\nINCRBY\r\n$5\r\ncount\r\n:3\r\n")

	sendCommand("*3\r\n$11\r\nINCRBYFLOAT\r\n$5\r\ncount\r\n,0.003\r\n")
}

func sendCommand(command string) {
	con, _ := net.Dial("tcp", "localhost:6379")
	defer con.Close()

	writer := bufio.NewWriter(con)
	reader := bufio.NewReader(con)

	writer.Write([]byte(command))
	writer.Flush()

	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Printf("%s", bytes)
}
