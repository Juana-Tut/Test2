//Filename: main.go - Echo server upgraded
// Description: This is an upgraded version of the Echo server that uses concurrency to handle multiple requests simultaneously,
// logging, command handling and other features. It is designed to be more efficient and robust than the original version.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)
const inactivityTimeout = 30 * time.Second // Define the inactivity timeout
const maxMessageLength int= 1024 // Define the maximum message length

func handleConnection(conn net.Conn) {
	clientAddr := conn.RemoteAddr().String()
	clientIP := strings.Split(clientAddr, ":")[0] // Extract the client IP address
	fmt.Printf("[%s] Client connected: %s\n", time.Now().Format(time.RFC3339), clientAddr)
	
	defer func() {
		fmt.Printf("[%s] Client disconnected: %s\n", time.Now().Format(time.RFC3339), clientAddr)
		conn.Close()
	}() // Close the connection when the function returns
	
	// Open or create a log file for the client 
	logFileName := fmt.Sprintf("%s.log", clientIP)
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("[%s] Error opening log file for client %s: %v\n", time.Now().Format(time.RFC3339), clientAddr, err)
		return
	}
	defer logFile.Close() // Ensure the log file is closed when done

	reader := bufio.NewReader(conn) // Create a buffered reader for the connection

	for {
		
		conn.SetDeadline(time.Now().Add(inactivityTimeout)) // Set a deadline for inactivity
		input, err := reader.ReadString('\n') // Read until a new line from the client
		if err != nil {
			if err == io.EOF {
                // Client closed the connection
                fmt.Printf("[%s] Client %s closed the connection\n", time.Now().Format(time.RFC3339), clientAddr)
            } else if os.IsTimeout(err) {
				// Inactivity timeout occurred
                fmt.Printf("[%s] Client %s disconnected due to inactivity\n", time.Now().Format(time.RFC3339), clientAddr)
			}else {
                // Handle other errors gracefully
                fmt.Printf("[%s] Error reading from client %s: %v\n", time.Now().Format(time.RFC3339), clientAddr, err)
            }
			return
		}
		//Trim the input to remove any leading or trailing whitespace
		cleanInput := strings.TrimSpace(input)

		//Handle long messages
		if len(cleanInput) > maxMessageLength{
			// Reject or truncate the message
			truncateMessage := cleanInput[:maxMessageLength]
			rejectionMessage := fmt.Sprintf("Message too long. Truncated to: %s\n", truncateMessage)
			conn.Write([]byte(rejectionMessage)) // Send rejection message to the client
		

			// Log the truncated message to the client's log file
			logMessage := fmt.Sprintf("[%s] %s:%s\n", time.Now().Format(time.RFC3339), clientAddr, truncateMessage)
			if _, err := logFile.WriteString(logMessage); err != nil { 
				fmt.Printf("[%s] Error writing to log file for client %s: %v\n", time.Now().Format(time.RFC3339), clientAddr, err)
				return
			}
		}

		// Log the message to the client's log file
        logMessage := fmt.Sprintf("[%s] %s\n", time.Now().Format(time.RFC3339), cleanInput)
        if _, err := logFile.WriteString(logMessage); err != nil {
            fmt.Printf("[%s] Error writing to log file for client %s: %v\n", time.Now().Format(time.RFC3339), clientAddr, err)
            return
        }

		// Echo the message back to the client
		_, err = conn.Write([]byte(cleanInput + "\n")) // Send the cleaned input back to the client
		if err != nil {
			fmt.Printf("[%s] Error writing to client %s: %v\n", time.Now().Format(time.RFC3339), clientAddr, err)
			return
		}
	} 
}

func main() {
	// Define the port flag
	port := flag.String("port", "4000", "Port to listen on")
	flag.Parse() // Parse the command line flags

    // define the target host and port we want to connect to
	address := fmt.Sprintf(":%s", *port)
    listener, err := net.Listen("tcp", address) // Listen on the specified port
	if err != nil {
		panic(err)
	}
    defer listener.Close()
	fmt.Printf("Server listening on %s\n", address)

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

	
   