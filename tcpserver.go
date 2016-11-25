package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8086")

	if err != nil {
		fmt.Println("%v: ", err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("%v: ", err)
		}
		go readMessage(conn)
	}
}

func readMessage(conn net.Conn) error {
	for {
		_, err := read(conn)
		if err != nil {
			return errors.New("error occured in read connection")
		}
	}
}

func read(conn net.Conn) (int, error) {
	b := make([]byte, 1024)
	n, err := conn.Read(b)
	if err != nil {
		fmt.Println("error")
	}
	request := string(b[0:n])
	n, err = writeMessage(conn, string(request))
	return n, err
}

func writeMessage(conn net.Conn, request string) (int, error) {
	response := helper(request)
	n, err := conn.Write([]byte(response))
	return n, err
}

func helper(request string) string {
	switch request {
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
