package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/asvins/auth/models"
	"github.com/asvins/common_io"
	"github.com/asvins/utils/config"
)

const (
	EVENT_CREATED = iota
	EVENT_UPDATED
	EVENT_DELETED
)

func topic(event int, prefix string) (string, error) {
	var sufix string

	switch event {
	case EVENT_CREATED:
		sufix = "_created"
	case EVENT_UPDATED:
		sufix = "_updated"
	case EVENT_DELETED:
		sufix = "_deleted"
	default:
		return "", errors.New("[ERROR] Event not found")
	}

	return prefix + sufix, nil
}

func fireEvent(event int, usr *models.User) {
	b, err := json.Marshal(usr)
	if err != nil {
		// TODO tratar erro...
		return
	}

	topic, err := topic(event, "user")
	if err != nil {
		// TODO tratar erro...
		return
	}

	producer.Publish(topic, b)
}

func setupCommonIo() {
	cfg := common_io.Config{}

	err := config.Load("common_io_config.gcfg", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	/*
	*	Producer
	 */
	producer, err = common_io.NewProducer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	/*
	*	Consumer
	 */
	//consumer = common_io.NewConsumer(cfg)
	//consumer.HandleTopic("", nil)

	//if err = consumer.StartListening(); err != nil {
	//	log.Fatal(err)
	//}

}

/*
*	Here can be added the handlers for kafka topics
 */
