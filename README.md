# Tracker server for chat API

The purpose of the server is to track down the connected users and allow for them to be visible to each other. \
With this, the clients chat API will be able to establish a P2P connection with each other for communication.

## Implementation

* The server will listen to a TCP port for connections
* Data will be retrieved and sent to clients using a very simple text protocol on top of the transport layer
* To keep track of the connected users, the server will expect a heartbeat package every few seconds
