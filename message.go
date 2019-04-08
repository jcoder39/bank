package main

type MessageStatus int16

const (
	OK                 MessageStatus = 0
	NOT_FOUND          MessageStatus = 2
	INSUFFICIENT_FUNDS MessageStatus = 3
	ERROR              MessageStatus = 666
)

type Message struct {
	Status MessageStatus          `json:"status"`
	Params map[string]interface{} `json:"params"`
}

func NewMessage() *Message {
	message := new(Message)
	message.Status = OK
	message.Params = make(map[string]interface{})
	return message
}

func (this *Message) AddParam(key string, val interface{}) {
	this.Params[key] = val
}

func (this *Message) GetParam(key string) interface{} {
	return this.Params[key]
}

func (this *Message) ClearParams() {
	this.Params = make(map[string]interface{})
}
