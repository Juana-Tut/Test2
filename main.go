//Filename: main.go - Echo server upgraded
// Description: This is an upgraded version of the Echo server that uses concurrency to handle multiple requests simultaneously,
// logging, command handling and other features. It is designed to be more efficient and robust than the original version.

package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("[%s] Client connected: %s\n", time.Now().Format(time.RFC3339), clientAddr)
	
	defer func() {
		fmt.Printf("[%s] Client disconnected: %s\n", time.Now().Format(time.RFC3339), clientAddr)
		conn.Close()
	}() // Close the connection when the function returns

	buf := make([]byte, 1024) // Buffer to hold incoming data

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
                // Client closed the connection
                fmt.Printf("[%s] Client %s closed the connection\n", time.Now().Format(time.RFC3339), clientAddr)
            } else {
                // Handle other errors gracefully
                fmt.Printf("[%s] Error reading from client %s: %v\n", time.Now().Format(time.RFC3339), clientAddr, err)
            }
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
			fmt.Printf("[%s] Error accepting: %v\n", time.Now().Format(time.RFC3339),err)
			continue
		}

		go handleConnection(conn) // Handle the connection in a separate goroutine
	} 
}

	
   