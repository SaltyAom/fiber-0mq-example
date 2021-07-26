package bridge

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	zmq "github.com/pebbe/zmq4"
)

type Message struct {
	Content  string
	Response chan string
}

func set_identity(socket *zmq.Socket) {
	identity := fmt.Sprintf("%04X-%04X", rand.Intn(0x10000), rand.Intn(0x10000))
	socket.SetIdentity(identity)
}

func send(message Message) {
	socket, err := zmq.NewSocket(zmq.DEALER)
	set_identity(socket)

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

func bridge(messages chan Message) {
	for message := range messages {
		go send(message)
	}
}

func CreateBridge() chan Message {
	concurrent := int(math.Pow(2, 16))
	bridgeQueue := make(chan Message, concurrent)

	go bridge(bridgeQueue)

	return bridgeQueue
}
