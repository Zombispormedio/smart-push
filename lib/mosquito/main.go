package mosquito

import (
	
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/yosssi/gmq/mqtt"
	mqttClient "github.com/yosssi/gmq/mqtt/client"
)

type Subscriber struct {
	Options *mqttClient.ConnectOptions
	Client  *mqttClient.Client
	Routing func([]byte) error
	Topic string
}

func New(routing func([]byte) error) *Subscriber {

	subscriber := Subscriber{}

	subscriber.Client = mqttClient.New(&mqttClient.Options{
		ErrorHandler: func(err error) {
			log.Error(err)
		},
	})

	subscriber.Options = &mqttClient.ConnectOptions{
		Network:  "tcp",
		Address:  os.Getenv("MQTT_HOST"),
		ClientID: []byte("sub"),
		UserName: []byte(os.Getenv("MQTT_USER")),
		Password: []byte(os.Getenv("MQTT_PASS")),
	}

	subscriber.Topic = os.Getenv("MQTT_TOPIC")
	subscriber.Routing = routing

	return &subscriber
}

func (subscriber *Subscriber) Run() error {
var Error error
	if subscriber.Routing != nil {

		client := subscriber.Client
		defer client.Terminate()
		err := client.Connect(subscriber.Options)
		if err != nil {
			return err
		}

		err = client.Subscribe(&mqttClient.SubscribeOptions{
			SubReqs: []*mqttClient.SubReq{
				&mqttClient.SubReq{
					TopicFilter: []byte(subscriber.Topic),
					QoS:         mqtt.QoS0,
					Handler: func(topicName, message []byte) {
					
						Error:=subscriber.Routing(message)
						if Error!=nil{
							log.Error(Error)
						}
					},
				},
			},
		})
		if err != nil {
			return err
		}

	}
	
	return Error

}
