package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/asvins/auth/models"
	"github.com/asvins/common_io"
	"github.com/asvins/notification/mailer"
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

}

/*
*	Senders
 */
func sendUserCreated(usr *models.User) {
	topic, _ := common_io.BuildTopicFromCommonEvent(common_io.EVENT_CREATED, "user")
	b, err := json.Marshal(usr)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}

	m := mailer.Mail{
		To:      []string{usr.Email},
		Subject: "Bem Vindo!",
		Body:    mailer.TemplateWelcome,
	}

	mMsg, _ json.Marshal(m)

	producer.Publish("send_mail", mMsg)
	producer.Publish(topic, b)
}
