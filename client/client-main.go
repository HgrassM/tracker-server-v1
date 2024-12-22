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

The format is the same as the client message.

option = 1 -> Response to heartbeat
option = 2 -> Response to set username
option = 3 -> Response to get peer's adress
option = 4 -> Request for client to resend the message

data -> If option == 1, this field will be empty
data -> If option == 2, this field will contain a "OK" javascript string
data -> If option == 3, this field will contain a string of the requested ipv4 adress. Reuturns "FAIL" on failure
data -> If option == 4, this field will be empty

'\n' -> A javascript single character string to indicate the end of the message 

*/

package client

import (
	"net"
	"log"
	"sync"
	"errors"
	"time"
)

func ClientRoutine(client_conn net.Conn, connMap *sync.Map, counter_ptr *uint64, counter_mutex *sync.Mutex) {
	var bytes_read int
	var bytes_written int
	received_message_buffer := make([]byte, 1024)
	timeout := time.Time{}.Add(Duration.Seconds(5)) 

	//Setting timeout for connection
	err := client_conn.SetReadDeadLine(timeout)
	
	//Adding counter
	counter_mutex.Lock()
	*counter_ptr = *counter_ptr + 1
	counter_mutex.Unlock()

	if err != nil {
		log.Printf("[ERROR] Couldn't set timeout for client tcp conn of address "%s". Message: %s \n", client_addr, err)
		close_err := client_conn.Close()
		
		if (close_err != nil) {
			log.Printf("[ERROR] Couldn't close client tcp conn of address "%s". Message: %s \n", client_addr, close_err)		
		}

		return 
	}

	for {	
		//Tries to read message from client
		bytes_read, io_err = client_conn.Read(message_buffer)

		if (io_err) != nil {
			log.Printf("[ERROR] Error while trying to read from client tpc conn of address "%s". Message: %s \n", client_addr, io_err)
			continue
		}
		
		//Get response message
		message_to_write = getResponse(message_buffer, bytes_read, client_conn, connMap *sync.Map)

		bytes_written, io_err = client_conn.Write(message_to_write)

		if (io_err) != nil {
			log.Printf("[ERROR] Error while trying to read from client tpc conn of address "%s". Message: %s \n", client_addr, io_err)
			continue
		}

	}
}
