package mosquito

import (
	"os"
log "github.com/Sirupsen/logrus"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Subscriber struct {
	ClientOptions *mqtt.ClientOptions
	Client        mqtt.Client
	Chan          chan []byte
	Topic         string
	Routing       func([]byte) error
}

func New(routing func([]byte) error) *Subscriber {

	subscriber := Subscriber{}
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://" + os.Getenv("MQTT_HOST"))
	opts.SetUsername(os.Getenv("MQTT_USER"))
	opts.SetPassword(os.Getenv("MQTT_PASS"))
	opts.SetClientID("sub")
	subscriber.ClientOptions = opts
	subscriber.Topic = os.Getenv("MQTT_TOPIC")
	subscriber.Routing = routing

	subscriber.Chan = make(chan []byte)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {

		subscriber.Chan <- msg.Payload()
	})

	subscriber.Client = mqtt.NewClient(opts)

	return &subscriber
}

func (subscriber *Subscriber) Run(Error chan bool) {

	if subscriber.Routing == nil {
		Error <- true
	} else {

		client := subscriber.Client
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			Error <- true
		} else {
			if token := client.Subscribe(subscriber.Topic, byte(0), nil); token.Wait() && token.Error() != nil {
				Error <- true
			} else {
				for {
					incoming := <-subscriber.Chan
					RouteError := subscriber.Routing(incoming)
					if RouteError != nil {
                        log.Error(RouteError)
						Error <- true
						break
					}
				}
			}

		}
	}

}
