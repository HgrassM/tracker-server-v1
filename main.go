package main

import (
	"net"
	"log"
	"sync"
	"tracker-server/client"
)

//Constants
const NETWORK_PROTOCOL = "tcp"
const LISTENER_ADDR = "127.0.0.1:9595"

//Global variables
var connMap *sync.Map
var total_conn_num *uint64

/*Main execution loop that listens to new connection requests */
func main() {
	//Creating tcp listener
	con_listener, err := net.Listen(NETWORK_PROTOCOL, LISTENER_ADDR)

	if err != nil {
		log.Fatal(err)
	}
	
	//Closes at the end of execution
	defer con_listener.Close()
	log.Printf("[INFO] Initializing server on %s\n", LISTENER_ADDR)

	//Handling connection requests
	for {
		conn, err := con_listener.Accept()

		if err != nil {
			log.Printf("[ERROR] Listener couldn't handle new connection request: %s \n", err)
		}
		
		//Gets client address and uses it as key for "connMap"
		conn_addr := conn.RemoteAddr().String()
		connMap.Store(conn_addr, conn)

		//Runs routine
		go client.ClientRoutine(conn_addr, connMap, total_conn_num) 
	}
}
