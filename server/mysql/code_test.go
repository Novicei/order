package mysql

import "testing"

func TestQueryByUserID(t *testing.T) {
	InitDB("root", "123456", "192.168.11.239", "order")
	QueryByUserID("8630603564438221213", "orderT")
}

func TestQueryBySkuID(t *testing.T) {
	InitDB("root", "123456", "192.168.11.239", "order")
	res, _ := QueryBySkuID("1", "orderT")
	println(res)
}

func TestQueryByTimeTop10(t *testing.T) {
	InitDB("root", "123456", "192.168.11.239", "order")
	res, _ := QueryByTimeTop10(1636012444, 1636012475, "orderT")
	for _, value := range res {
		println(value)
	}
}
