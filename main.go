package main

import (
	"net"
	"log"
	"sync"
	"tracker-server/client"
)

//Constants
const NETWORK_PROTOCOL = "tcp"
const LISTENER_ADDR = ":9595"

//Global variables
var total_conn_num *uint64
var connMap *sync.Map
var counter_mutex *sync.Mutex

/*Main execution loop that listens to new connection requests */
func main() {
	//Creating tcp listener
	con_listener, err := net.Listen(NETWORK_PROTOCOL, LISTENER_ADDR)

	if err != nil {
		log.Fatal(err)
	}
	
	//Closes at the end of execution
	defer con_listener.Close()
	log.Printf("[INFO] Initializing server on 127.0.0.1%s\n", LISTENER_ADDR)

	//Handling connection requests
	for {
		conn, err := con_listener.Accept()

		if err != nil {
			log.Printf("[ERROR] Listener couldn't handle new connection request: %s \n", err)
			continue
		}

		//Runs routine
		go client.ClientRoutine(conn, connMap, total_conn_num, counter_mutex) 
	}
}
