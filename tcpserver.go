package main

import (
	"errors"
	"fmt"
	"net"
)

//accept-main
//read
//write

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
	//[]byte array
	/*if err != nil {
		fmt.Println(err)
	}*/
	//fmt.Println("Message read successfully:-")
	//fmt.Println(string(request))
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

// func handleRequest(conn net.Conn) {
// 	for {
// 		message, _ := bufio.NewReader(conn).ReadString('\n')
// 		fmt.Println(string(message))
// 		conn.Write([]byte("Message received: " + message))
// 	}

// }
