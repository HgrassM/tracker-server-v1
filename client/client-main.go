/*
CLIENT MESSAGE FORMAT:

Each field represents a javascript string character count

option    data      '\n'
  ^        ^         ^
  |        |         | 
------------------------
| 1 |      N       | 1 |
------------------------

option = 1 -> Heartbeat message
option = 2 -> Set the username of the client
option = 3 -> Get peer's adress to start P2P connection

data -> If option == 1, this field will be ignored
data -> If option == 2, this field needs to contain the desired username
data -> If option == 3, this field needs to contain the peer's nickname whose client wants to connect

'\n' -> A javascript single character string to indicate the end of the message

SERVER RESPONSE MESSAGE FORMAT:

The format is the same as the client message. There is no response to "option = 1" messages.

option = 1 -> Request for client to resend the message
option = 2 -> Response to set username
option = 3 -> Response to get peer's adress

data -> If option == 1, this field will be empty
data -> If option == 2, this field will contain a "OK" javascript string on sucess and a "FAIL" on failure
data -> If option == 3, this field will contain a string of the requested ipv4 adress

'\n' -> A javascript single character string to indicate the end of the message 

*/

package client

import (
	"net"
	"log"
	"sync"
	"errors"
	"os"
	"time"
)

func CheckOption

func ClientRoutine(client_conn net.Conn, connMap *sync.Map, counter_ptr *uint64, counter_mutex *sync.Mutex) {
	//Closes connection at the end of the routine
	defer client_conn.Close()
	
	//Variables
	var bytes_read int
	var byter_written int
	var io_err error
	var received_option string
	var received_data string
	var message_len int
	received_message_buffer := make([]byte, 1024)
	client_addr := client_conn.RemoteAddr()
	timeout := time.Time{}.Add(Duration.Seconds(5)) 

	//Setting timeout for connection and adding counter
	err := client_conn.SetReadDeadLine(timeout)
	counter_mutex.Lock()
	*counter_ptr = *counter_ptr + 1
	counter_mutex.Unlock()

	if err != nil {
		log.Printf("[ERROR] Couldn't set timeout for client tcp conn of address "%s". Message: %s \n", client_addr, err)
		return 
	}

	for {	
		//Tries to read message from client
		bytes_read, io_err = client_conn.Read(message_buffer)
		message_len = bytes_read - 1

		if (io_err) != nil {
			log.Printf("[ERROR] Error while trying to read from client tpc conn of address "%s". Message: %s \n", client_addr, io_err)
			return
		}
		
		//If the received message has no delimiter, the client is asked to resend it
		if (message_buffer[message_len] != "\n") {
			write_message = []byte("1\n")

			bytes_written, io_err = client_conn.Write(write_message)

			if (io_err) != nil {
				log.Printf("[ERROR] Error while trying to read from client tpc conn of address "%s". Message: %s \n", client_addr, io_err)
				return
			}
		}else{
			//Getting message fields
			received_option = string(received_message_buffer[0])
			 
			if (bytes_read > 2) {
				received_data = string(received_message_buffer[1:message_len])
			}else{
				received_data = nil	
			}

			switch (received_option) {
				//TODO implement option pattern matching
			}
		}

	}
}
