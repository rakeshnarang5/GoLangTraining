package main

import (
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

func readMessage(conn net.Conn) (int, error) {
	for {
		//request, err := bufio.NewReader(conn).ReadString('\n')
		b := make([]byte, 1024)
		n, err := conn.Read(b)
		request := string(b[0:n])
		//[]byte array
		/*if err != nil {
			fmt.Println(err)
		}*/
		// fmt.Println("Message read successfully:-")
		// fmt.Println(string(message))
		go writeMessage(conn, string(request))
		return n, err
	}
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
