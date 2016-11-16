package main

import (
	"bufio"
	"fmt"
	"net"
)

//accept-main
//read
//write

func main() {
	ln, _ := net.Listen("tcp", ":8087")
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go readMessage(conn)
	}
}

func readMessage(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		//[]byte array
		fmt.Println("Message read successfully:-")
		fmt.Println(string(message))
		go writeMessage(conn, string(message))
	}
}

func writeMessage(conn net.Conn, message string) {
	retVal := helper(message)
	conn.Write([]byte(retVal))
}

func helper(message string) string {

	switch message {
	case "Hello\n":
		return "Hi\n"
	case "Name\n":
		return "Chitty\n"
	case "GoodBye\n":
		return "Bye\n"
	default:
		return "Could not identify this command.\n"
	}

}

// func handleRequest(conn net.Conn) {
// 	for {
// 		message, _ := bufio.NewReader(conn).ReadString('\n')
// 		fmt.Println(string(message))
// 		conn.Write([]byte("Message received: " + message))
// 	}

// }
