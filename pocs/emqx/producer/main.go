package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	client mqtt.Client
	opts   *mqtt.ClientOptions
}

// NewClient return a new connection of broker
func NewClient(id string) *Client {

	opts := &mqtt.ClientOptions{
		ClientID:      id,
		CleanSession:  true,
		AutoReconnect: true,
		//MaxReconnectInterval: 120 * time.Second,
		//KeepAlive:            30,
	}

	opts.AddBroker("tcp://localhost:1883")
	client := mqtt.NewClient(opts)

	return &Client{client, opts}
}

// Connect to broker
func (c *Client) Connect() error {
	token := c.client.Connect()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Publish publishes a message on a specific topic.
func (c *Client) Publish(topic, msg string) error {
	token := c.client.Publish(topic, 0, false, msg)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Subscribe creates a subscription for the passed topic.
func (c *Client) Subsribe(topic string, f mqtt.MessageHandler) error {
	token := c.client.Subscribe(topic, 0, f)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// SetPublishHandler set a default handler for message
func (c *Client) SetPublishHandler(f mqtt.MessageHandler) {
	c.opts.SetDefaultPublishHandler(f)
	c.client = mqtt.NewClient(c.opts)
}

func main() {
	topic := "test/message"

	client := NewClient("producer")

	if err := client.Connect(); err != nil {
		fmt.Printf("Got error from : %s\n", err.Error())
		os.Exit(1)
	}

	if err := client.Subsribe(topic, nil); err != nil {
		fmt.Printf("Got error from : %s\n", err.Error())
		os.Exit(1)
	}

	for i := 0; i < 100; i++ {
		fmt.Println("Send message #", i)
		message := fmt.Sprintf("hello world from #%d", i)
		if err := client.Publish(topic, message); err != nil {
			fmt.Printf("Got error from : %s\n", err.Error())
			os.Exit(1)
		}
		time.Sleep(time.Duration(2) * time.Second)
	}

	select {}

}
