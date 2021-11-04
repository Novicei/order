package mock

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

type Order struct {
	OrderID   int64   `json:"orderId"`
	OrderTime int64   `json:"orderTime"`
	SkuID     int64   `json:"skuId"`
	UserID    int64   `json:"userID"`
	Status    int     `json:"status"`
	Price     float64 `json:"price"`
	Pay       float64 `json:"pay"`
}

type OrderManage struct {
	CommodityHash map[int64]float64 //保存商品价格
	OrderHash     map[int64]int     //保存订单状态
	UserHash      map[int64]int64   //key:orderID value:userID
}

func Init() *OrderManage {
	return &OrderManage{
		CommodityHash: map[int64]float64{},
		OrderHash:     map[int64]int{},
		UserHash:      map[int64]int64{},
	}
}

func (this *OrderManage) Mock(prob float64, orderID int64, commodityID int64) (used bool, data string) {
	if rand.Float64() > prob {
		if this.OrderHash[orderID] < 3 {
			var order *Order
			order = &Order{
				OrderID:   orderID,
				OrderTime: time.Now().Unix(),
				SkuID:     commodityID,
				UserID:    this.UserHash[orderID],
				Status:    this.OrderHash[orderID],
				Price:     this.CommodityHash[commodityID],
				Pay:       0,
			}
			if order.UserID == 0 {
				order.UserID = rand.Int63()
				this.UserHash[orderID] = order.UserID
			}
			if this.OrderHash[orderID] == 1 {
				order.Pay = order.Price
			}
			data, err := json.Marshal(order)
			if err != nil {
				log.Printf("order marshal err!err=%v\n", err)
				return false, ""
			}
			this.OrderHash[orderID]++
			return true, string(data)
		}
		return true, ""
	}
	return this.UserHash[orderID] > 0, ""
}
