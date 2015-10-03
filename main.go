package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/bitly/go-nsq"
	"github.com/optonaut/cron-server/config"
	"github.com/optonaut/shared-server/utils"
	"github.com/optonaut/shared-server/workerdata"
)

func job() {
	fmt.Println("Every hour")

}

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	cfg := config.Parse()

	producer, err := nsq.NewProducer(cfg.NSQd.Host+":"+cfg.NSQd.TCPPort, nsq.NewConfig())
	if err != nil {
		logrus.WithField("error", err).Fatal("Could not connect to nsq")
	}

	producer.SetLogger(utils.Logger{logrus.WithFields(logrus.Fields{})}, nsq.LogLevelInfo)

	for {
		if err := producer.Ping(); err != nil {
			logrus.WithField("error", err).Error("Trying to connect to nsq...")
			time.Sleep(500 * time.Millisecond)
		} else {
			break
		}
	}

	logrus.Info("Connected to nsq")

	t := time.NewTicker(time.Hour)
	for {

		payload := workerdata.CronHourlyPayload{
			Time: time.Now(),
		}
		data, _ := json.Marshal(payload)

		if err := producer.Publish(workerdata.CronHourlyTopic, data); err != nil {
			logrus.WithField("error", err).Error("couldn't write message to queue")
		}
		logrus.WithField("payload", payload).Info("cron hourly")
		<-t.C
	}
}
