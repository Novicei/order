package main

import (
	"Order/server/config"
	"Order/server/kafka"
	"Order/server/mysql"
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	var conf = new(config.Conf)
	err := ini.MapTo(conf, "server/config/config.ini")
	if err != nil {
		log.Printf("load ini failed!\n")
		return
	}
	mysql.InitDB(conf.MysqlConf.UserName, conf.MysqlConf.PassWord, conf.MysqlConf.Address, conf.MysqlConf.DataBase)
	err = kafka.Init([]string{conf.KafkaConf.Address}, conf.KafkaConf.Topic, conf.MysqlConf.TableName)
	if err != nil {
		log.Printf("kafka init failed!err:%v\n", err)
		return
	}
	select {}
}
