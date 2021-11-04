package kafka

import (
	"Order/server/mysql"
	"github.com/Shopify/sarama"
	"log"
)

func Init(address []string, topic string, tableName string) (err error) {
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		log.Printf("fail to start consumer, err:%v\n", err)
		return err
	}

	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		log.Printf("fail to get list of partition:err%v\n", err)
		return err
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				if err != nil {
					log.Printf("kafka consumer unmarshal err,err:%v\n", err)
					continue
				}
				mysql.Insert(tableName, msg.Value)
			}
		}(pc)
	}
	select {}
	return nil
}
