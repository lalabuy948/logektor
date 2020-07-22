package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"gopkg.in/ini.v1"
)

// Event - structure of json body
type Event struct {
	ObjectID      int64  `json:"objectId"`
	ObjectType    string `json:"objectType"`
	Action        string `json:"action"`
	SubjectID     int64  `json:"subjectId"`
	SubjectType   string `json:"subjectType"`
	DatetimeStamp int64  `json:"datetimeStamp"`
}

var (
	kafkaURL string
	topic    string
	writer   *kafka.Writer
)

func initProducer() error {
	cfg, err := ini.Load("env.ini")
	if err != nil {
		return err
	}

	kafkaURL = cfg.Section("kafka").Key("host").String()
	topic = "tracker"

	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return nil
}

func track(rawEvent []byte) error {

	// Would be nice to do a proper validation
	var event Event
	err := json.Unmarshal(rawEvent, &event)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(fmt.Sprint("event")),
		Value: rawEvent,
	}

	err = writer.WriteMessages(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
