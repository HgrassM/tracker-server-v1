package main

import (
	"net"
	"log"
)

//Constants
const NETWORK_PROTOCOL = "tcp"
const LISTENER_ADDR = "127.0.0.1:9595"

//Global variables


/*Main execution loop that listens to new connection requests */
func main() {
	//Creating tcp listener
	con_listener, err := net.Listen(NETWORK_PROTOCOL, LISTENER_ADDR)

	if err != nil {
		log.Fatal(err)
	}
	
	//Closes at the end of execution
	defer con_listener.Close()

	//Handling connection requests
	for {
		conn, err := con_listener.Accept()

		if err != nil {
			log.Printf("[ERROR] Listener couldn't handle new connection request: %s \n", err)
		}

		//TODO pass the new connection to a go routine that handles the connection
	}
}
