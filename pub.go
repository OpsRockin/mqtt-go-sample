package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

import MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"

func main() {
	stdin := bufio.NewReader(os.Stdin)
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

	for {
		message, err := stdin.ReadString('\n')
		if err == io.EOF {
			os.Exit(0)
		}
		r := client.Publish(MQTT.QoS(*qos), *topic, []byte(strings.TrimSpace(message)))
		<-r
		//fmt.Print("Message Sent")
	}
}
