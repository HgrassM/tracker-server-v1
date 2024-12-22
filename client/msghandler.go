package client

import (
	"log"
	"errors"
	"sync"
	"net"
	"fmt"
)

const HEARTBEAT_RESPONSE = "1\n"  
const ASK_RESEND_RESPONSE = "4\n"
const OPTION_2_SUCCESS = "2OK\n"
const OPTION_3_FAILURE = "3FAIL\n"

func setClientUsername(username string, client_conn net.Conn connMap *sync.Map) string {
	connMap.Store(username, client_conn)
	
	response := OPTION_2_SUCCESS_DATA

	return response
}

func getPeerAddr(username string, connMap *sync.Map) string {
	var response string
	peer_addr, ok := connMap.Load(username)

	if ok {
		response = fmt.Sprintf("3%s\n", peer_addr)
	}else{
		response = OPTION_3_FAILURE
	}

	return response
}

func getResponse(message_buffer []byte, message_bytes_num int, client_conn net.Conn, connMap *sync.Map) []byte {
	var io_err error
	var received_option string
	var received_data string
	var message_to_write string
	message_len := message_bytes_num - 1;

	//If the received message has no delimiter, the client is asked to resend it
	if (message_buffer[message_len] != "\n") {
		message_to_write = ASK_RESEND_RESPONSE
	}else{
		//Getting message fields
		received_option = string(received_message_buffer[0])
		 
		if (bytes_read > 2) {
			received_data = string(received_message_buffer[1:message_len])
		}else{
			received_data = nil	
		}

		switch (received_option) {
			case 1:
				//Respond to heartbeat
				messate_to_write = HEARTBEAT_RESPONSE
			case 2:
				//Set username
				message_to_write = setClientUsername(received_data, connMap)
			case 3:
				//Get the desired user's ip
				message_to_write = getPeerAddr(received_data, connMap)		
			default:
				//Ask to resend message if option doesn't exist
				message_to_write = ASK_RESEND_RESPONSE
		}
	}

	return []byte(message_to_write)
}
