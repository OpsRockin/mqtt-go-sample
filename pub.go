package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

import MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"

type Config struct {
	Host     string `toml:"Host"`
	Port     int    `toml:"Port"`
	Topic    string `toml:"Topic"`
	QoS      int    `toml:"QoS"`
	Username string `toml:"Username"`
	Password string `toml:"Password"`
}

func main() {
	var config Config
	var err error

	_, err = toml.DecodeFile("config.tml", &config)
	if err != nil {
		panic(err)
	}

	stdin := bufio.NewReader(os.Stdin)
	hostname, _ := os.Hostname()

	server := flag.String("server", config.Host+":"+strconv.Itoa(config.Port), "The endpoint of the MQTT")
	topic := flag.String("topic", hostname, "The topic to publish the messages on")
	qos := flag.Int("qos", config.QoS, "The QoS to send the messages at")
	client_id := flag.String("client_id", hostname+strconv.Itoa(time.Now().Second()), "A client ID for the connection")
	username := flag.String("username", config.Username, "A username to authenticate to the MQTT server")
	password := flag.String("password", config.Password, "Password to match username")
	flag.Parse()

	connOpts := MQTT.NewClientOptions().AddBroker(*server).SetClientId(*client_id).SetCleanSession(true)
	if *username != "" {
		connOpts.SetUsername(*username)
		if *password != "" {
			connOpts.SetPassword(*password)
		}
	}

	client := MQTT.NewClient(connOpts)
	_, err = client.Start()
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
		fmt.Print("Message Sent\n")
	}
}
