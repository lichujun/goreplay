package main

import "fmt"

type ZmqOutput struct {
	address string
	config  *ZmqOutputConfig
	msg     chan *Message
}

type ZmqOutputConfig struct {
}

func NewZmqOutput(address string, config *ZmqOutputConfig) PluginWriter {
	msg := make(chan *Message, 100)
	zmqOutput := &ZmqOutput{
		address: address,
		config:  config,
		msg:     msg,
	}
	go consumeMsg(msg)
	return zmqOutput
}

func (o *ZmqOutput) PluginWrite(msg *Message) (n int, err error) {
	if !isOriginPayload(msg.Meta) {
		return len(msg.Data), nil
	}
	o.msg <- msg
	return len(msg.Data) + len(msg.Meta), nil
}

func consumeMsg(msg <-chan *Message) {
	for message := range msg {
		fmt.Println("data:", string(message.Data))
		fmt.Println("meta:", string(message.Meta))
	}
}
