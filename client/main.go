package main

import (
	"Order/client/config"
	"Order/client/kafka"
	"Order/client/mock"
	"gopkg.in/ini.v1"
	"log"
	"math/rand"
	"time"
)

var (
	conf = new(config.Conf)
)

func main() {
	rand.Seed(time.Now().Unix())
	err := ini.MapTo(conf, "client/config/config.ini")
	if err != nil {
		log.Printf("load ini failed!\n")
		return
	}
	err = kafka.Init([]string{conf.Address}, conf.ChanMaxSize)
	if err != nil {
		log.Printf("init Kafka failed,err=%v\n", err)
		return
	}
	log.Printf("kafka init success\n")

	ordermanage := mock.Init()
	var orderID = int64(0)
	var commodityID = int64(0)
	for i := 0; i < conf.OrderClient.CommodityNum; i++ {
		ordermanage.CommodityHash[int64(i)] = float64(i + 1)
	}
	for {
		used, data := ordermanage.Mock(conf.OrderClient.Prob, orderID, commodityID)
		if data == "" {
			if used == true {
				orderID = orderID + 1
			}
			commodityID = commodityID + 1
			continue
		}
		kafka.SendToChan(conf.OrderClient.Topic, data)
		time.Sleep(1 * time.Second)
	}
}
