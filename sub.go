package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

import MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"

func onMessageReceived(client *MQTT.MqttClient, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\n", message.Topic())
	fmt.Printf("Message: %s\n", message.Payload())
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("signal received, existing")
		os.Exit(0)
	}()

	hostname, _ := os.Hostname()

	server := flag.String("server", "tcp://127.0.0.1:1883", "The endpoint of the MQTT")
	topic := flag.String("topic", hostname, "The topic to publish the messages on")
	qos := flag.Int("qos", 0, "The QoS to send the messages at")
	client_id := flag.String("client_id", hostname+strconv.Itoa(time.Now().Second()), "A client ID for the connection")
	username := flag.String("username", "", "A username to authenticate to the MQTT server")
	password := flag.String("password", "", "Password to match username")
	flag.Parse()

	connOpts := MQTT.NewClientOptions().AddBroker(*server).SetClientId(*client_id).SetCleanSession(true)
	if *username != "" {
		connOpts.SetUsername(*username)
		if *password != "" {
			connOpts.SetPassword(*password)
		}
	}

	client := MQTT.NewClient(connOpts)
	_, err := client.Start()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Connected to %s\n", *server)
	}

	filter, e := MQTT.NewTopicFilter(*topic, byte(*qos))
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	client.StartSubscription(onMessageReceived, filter)

	for {
		time.Sleep(1 * time.Second)
	}
}
