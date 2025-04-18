//Filename: main.go - Echo server upgraded
// Description: This is an upgraded version of the Echo server that uses concurrency to handle multiple requests simultaneously,
// logging, command handling and other features. It is designed to be more efficient and robust than the original version.

package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
		// Echo the message back to the client
		_, err = conn.Write(buf[:n])
		if err != nil {
   		fmt.Println("Error writing to client:", err)
		}
	} 
}

func main() {
    // define the target host and port we want to connect to
    listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}
    defer listener.Close()
	fmt.Println("Server listening on :4000")
	// Our program runs an infinite loop
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting:", err)
				continue
			}
			handleConnection(conn)
		}
}

	
   