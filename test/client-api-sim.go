package main

import (
	"log"
	"sync"
	"net"
	"time"
)

//Connection constants
const NETWORK_PROTOCOL = "tcp"
const TRACKER_ADDR = "127.0.0.1"
const TRACKER_PORT = ":9595"

//Message constants
const HEARTBEAT_MSG = "1\n"

//Sync variables
var wg sync.WaitGroup

func clientApi1() {
	var bytes_written int
	var io_err error
	set_name_msg := []byte("2clientOne\n")
	first_message := true

	//Connecting to server
	log.Printf("[CLIENT_1] [INFO] Connecting to tracker server of address: %s.\n", TRACKER_ADDR)
	api1_conn, err := net.Dial(NETWORK_PROTOCOL, TRACKER_ADDR + TRACKER_PORT)
	
	if (err != nil) {
		log.Printf("[CLIENT_1] [ERROR] Failed to connect to tracker server of address: %s.\n", TRACKER_ADDR)
		return
	}
	
	//Connection loop
	for {
		if (first_message) {
			//Setting username
			api1_conn.Write(set_name_msg)
		}else{
			//Writting heartbeat message to the server
			bytes_written, io_err = api1_conn.Write([]byte(HEARTBEAT_MSG))
		}

		//Handling IO error
		if (io_err != nil) {
			log.Printf("[CLIENT_1] [INFO] Failed to write message to the server of address: %s\n", TRACKER_ADDR)
			io_err = api1_conn.Close()

			if (io_err != nil) {
				log.Println("[CLIENT_1] [INFO] Failed to close connection.")
			}
			
			wg.Done()
			return
		}
		
		//Log after writting message
		if (first_message) {
			log.Printf("[CLIENT_1] [INFO] Setting username to %s.\n", string(set_name_msg))
			first_message = false
		}else{
			log.Println("[CLIENT_1] [INFO] Sending heartbeat message to keep the connection alive.")
		}

		log.Printf("[CLIENT_1] [INFO] %d bytes were sent to the server of address: %s.\n", bytes_written, TRACKER_ADDR)
		
		//Wait before continuing
		time.Sleep(3*time.Second)
	}

	wg.Done()
}

func clientApi2() {
	var bytes_written int
	var io_err error
	first_message := true
	set_name_msg := []byte("2clientTwo\n")

	//Connecting to server
	log.Printf("[CLIENT_2] [INFO] Connecting to tracker server on address: %s.\n", TRACKER_ADDR)
	api2_conn, err := net.Dial(NETWORK_PROTOCOL, TRACKER_ADDR + TRACKER_PORT)

	if (err != nil) {
		log.Printf("[CLIENT_2] [ERROR] Failed to connect to tracker server of address: %s.\n", TRACKER_ADDR)
		return
	}	

	//Connection loop
	for {
		if (first_message) {
			//Setting username
			api2_conn.Write(set_name_msg)
		}else{
			//Writting heartbeat message to the server
			bytes_written, io_err = api2_conn.Write([]byte(HEARTBEAT_MSG))
		}

		//Handling IO error
		if (io_err != nil) {
			log.Printf("[CLIENT_2] [INFO] Failed to write message to the server of address: %s\n", TRACKER_ADDR)
			io_err = api2_conn.Close()

			if (io_err != nil) {
				log.Println("[CLIENT_2] [INFO] Failed to close connection.")
			}

			wg.Done()
			return
		}
		
		//Log after writing message
		if (first_message) {
			log.Printf("[CLIENT_2] [INFO] Setting username to %s.\n", string(set_name_msg))
			first_message = false
		}else{
			log.Println("[CLIENT_2] [INFO] Sending heartbeat message to keep the connection alive.")
		}

		log.Printf("[CLIENT_2] [INFO] %d bytes were sent to the server of address: %s.\n", bytes_written, TRACKER_ADDR)
		
		//Wait before continuing
		time.Sleep(3*time.Second)
	}

	wg.Done()
}

func main() {
	wg.Add(2)

	go clientApi1()
	go clientApi2()

	wg.Wait()
}
