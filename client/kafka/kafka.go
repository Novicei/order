package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

var (
	client        sarama.SyncProducer
	orderDataChan chan *orderData
)

type orderData struct {
	topic string
	data  string
}

func Init(addr []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return err
	}
	orderDataChan = make(chan *orderData, maxSize)
	go sendToKafka()
	return nil
}

func sendToKafka() {
	for {
		select {
		case orderData := <-orderDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = orderData.topic
			msg.Value = sarama.StringEncoder(orderData.data)
			_, _, err := client.SendMessage(msg)
			if err != nil {
				log.Println("send msg failed, err:", err)
			}
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

func SendToChan(topic, data string) {
	msg := &orderData{
		topic: topic,
		data:  data,
	}
	orderDataChan <- msg
}
