# Echo Server
## How to run the server
1. **Ensure you have Go installed and setup on your system.**
2. **Clone this repository to your local machine.**
3. **Run the server:**
- To run the server with default port (4000)
  ```bash
  go run main.go
  ```
- To run the server with a custom port(e.g., 3000)
  ```bash
  go run main.go -port 3000
  ```
4. **Connect a client**
- Use a tool like telnet or nc to connect to server:
  ```sh
  nc localhost 4000
  ```
  - Send messages and they will be logged to a file named after the client's IP address followed by .log

## Link to video 
()

## Educationally Enriching Functionality
The most educationally enriching functionality was implementing the **per-client logging system**. It required understanding how to handle concurrent connections, manage file I/O efficiently, and ensure thread safety when multiple clients are connected.

## Functionality Requiring the Most Research
The functionality that required the most research was **creating and managing the log files for each client**. This involved:
- Learning how to dynamically create and name log files based on the client's IP address.
- Handling file I/O operations efficiently in Go.
- Ensuring proper cleanup and error handling when writing to log files.
