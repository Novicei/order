package config

type Conf struct {
	KafkaConf   `ini:"kafka"`
	OrderClient `ini:"orderClient"`
}

type KafkaConf struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:"chan_max_size"`
}

type OrderClient struct {
	Topic        string  `ini:"topic"`
	Prob         float64 `ini:"prob"`
	CommodityNum int     `ini:"commodityNum"`
}
