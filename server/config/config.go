package config

type Conf struct {
	KafkaConf `ini:"kafka"`
	MysqlConf `ini:"mysql"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type MysqlConf struct {
	Address   string `ini:"address"`
	UserName  string `ini:"username"`
	PassWord  string `ini:"password"`
	DataBase  string `ini:"database"`
	TableName string `ini:"tableName"`
}
