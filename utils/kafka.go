package utils

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Shopify/sarama"

	"github.com/mmc-projects/mmc_core/schemas"
)

func ProduceLogMessage(level schemas.LogLevel, message string, service string, producer sarama.SyncProducer) error {
	kafka_topic := "mmclogs"

	// log the health check
	logMesg := schemas.LogMessage{
		Level:     level,
		Message:   message,
		CreatedAt: time.Now(),
		Service:   service,
	}

	messageJSON, err := json.Marshal(logMesg)
	if err != nil {
		log.Println("Error encoding JSON: ", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: kafka_topic,
		Key:   nil,
		Value: sarama.StringEncoder(messageJSON),
	}
	_, _, _ = producer.SendMessage(msg)
	return nil
}


func ProduceNotificationMsg(notiEvent schemas.NotificationEvent, producer sarama.SyncProducer) error {
	kafka_topic := "mmcnotifications"
	messageJSON, err := json.Marshal(notiEvent)
	if err != nil {
		log.Println("Error encoding JSON: ", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: kafka_topic,
		Key:   nil,
		Value: sarama.StringEncoder(messageJSON),
	}
	_, _, _ = producer.SendMessage(msg)
	return nil
}
