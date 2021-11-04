package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sort"
	"strconv"
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

var db *sql.DB

func InitDB(username, password, ip, database string) (err error) {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", username, password, ip, database)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func Insert(tableName string, data []byte) {
	order := Order{}
	err := json.Unmarshal(data, &order)
	if err != nil {
		log.Printf("json unmarshal error!err=%v", err)
		return
	}
	sql := fmt.Sprintf("insert into %s(orderID,orderTime,skuId,userId,status,price,pay) values (?,?,?,?,?,?,?)", tableName)
	_, err = db.Exec(sql, strconv.FormatInt(order.OrderID, 10), strconv.FormatInt(order.OrderTime, 10), strconv.FormatInt(order.SkuID, 10), strconv.FormatInt(order.UserID, 10), strconv.Itoa(order.Status), order.Price, order.Pay)
	if err != nil {
		log.Printf("insert failed,err:%v\n", err)
		return
	}
}

func QueryByUserID(userID string, tableName string) ([]string, error) {
	res := []string{}
	sql := fmt.Sprintf("select * from %s where userId = ?", tableName)
	rows, err := db.Query(sql, userID)
	if err != nil {
		log.Printf("query failed!err=%v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order Order
		rows.Scan(&order.OrderID, &order.OrderTime, &order.SkuID, &order.UserID, &order.Status, &order.Price, &order.Pay)
		t, err := json.Marshal(order)
		if err != nil {
			log.Printf("QueryByUserID_json_marshal_err!err=%v", err)
			return nil, err
		}
		res = append(res, string(t))
	}
	return res, nil
}

func QueryBySkuID(SkuID string, tableName string) (float64, error) {
	status := [3]int{}
	sql := fmt.Sprintf("select * from %s where SkuID = ?", tableName)
	rows, err := db.Query(sql, SkuID)
	if err != nil {
		log.Printf("query failed!err=%v", err)
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var order Order
		rows.Scan(&order.OrderID, &order.OrderTime, &order.SkuID, &order.UserID, &order.Status, &order.Price, &order.Pay)
		status[order.Status]++
	}
	return float64(status[1]) / float64(status[0]+status[2]), nil
}

func QueryByTimeTop10(start, end int64, tableName string) ([]int64, error) {
	hash := map[int64]bool{}
	count := map[int64]int{}
	res := []int64{}
	sql := fmt.Sprintf("select * from %s where orderTime >= ? and orderTime <= ?", tableName)
	rows, err := db.Query(sql, start, end)
	if err != nil {
		log.Printf("query failed!err=%v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order Order
		rows.Scan(&order.OrderID, &order.OrderTime, &order.SkuID, &order.UserID, &order.Status, &order.Price, &order.Pay)
		hash[order.SkuID] = true
		count[order.SkuID]++
	}
	for index, _ := range hash {
		res = append(res, index)
	}
	sort.Slice(res, func(i, j int) bool {
		if count[res[i]] == count[res[j]] {
			return res[i] < res[j]
		}
		return count[res[i]] > count[res[j]]
	})
	return res[:10], nil
}
