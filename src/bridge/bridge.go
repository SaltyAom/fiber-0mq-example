package bridge

import (
	"log"
	"math"

	zmq "github.com/pebbe/zmq4"
)

type BridgeMessage struct {
	Content  string
	Response chan string
}

func bridgeSend(message BridgeMessage) {
	socket, err := zmq.NewSocket(zmq.REQ)

	if err != nil {
		log.Fatal(err)
		return
	}

	socket.Connect("tcp://0.0.0.0:5556")

	socket.Send(message.Content, 0)
	receive, _ := socket.Recv(0)

	message.Response <- receive

	socket.Close()
}

func bridge(messages chan BridgeMessage) {
	for message := range messages {
		bridgeSend(message)
	}
}

func CreateBridge() chan BridgeMessage {
	concurrent := int(math.Pow(2, 24))
	bridgeQueue := make(chan BridgeMessage, concurrent)

	go bridge(bridgeQueue)

	return bridgeQueue
}
