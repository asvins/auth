package main

import (
	"log"

	"github.com/asvins/common_io"
	"github.com/asvins/utils/config"
)

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
