package main

import (
	"context"
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

const (
	ApiKey          = ""
	ApiSecret       = ""
	TargetTimestamp = int64(1709768000)
	OrderSymbol     = "TIAUSDT"
	OrderSize       = 1
	OrderPrice      = 22
)

func main() {
	client := binance.NewFuturesClient(ApiKey, ApiSecret)

	defer func(c *futures.Client) {
		err := c.NewCancelAllOpenOrdersService().Symbol(OrderSymbol).Do(context.Background())
		if err != nil {
			panic(err)
		}

		fmt.Println("all orders cancelled")
	}(client)

	var sum int64
	results := make([]int64, 0, 20)

	client.NewListAccountTradeService().Do(context.Background())

	//var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		//wg.Add(1)
		//
		//go func() {
		//	defer wg.Done()

		res := sendOrder(client)
		sum += res
		results = append(results, res)

		time.Sleep(time.Second)
		fmt.Println(res)
		//}()
	}

	//wg.Wait()

	fmt.Printf("average time difference is (ms): %f\n", float64(sum)/20.0)
	fmt.Println(results)
}

func sendOrder(client *futures.Client) int64 {
	quantity := fmt.Sprintf("%d", OrderSize)
	price := fmt.Sprintf("%d", OrderPrice)
	now := time.Now()
	nowTs := now.UnixMilli()
	qb := client.NewCreateOrderService().
		Symbol(OrderSymbol).
		Side(futures.SideTypeSell).
		Type(futures.OrderTypeLimit).
		TimeInForce(futures.TimeInForceTypeGTC).
		Quantity(quantity).
		Price(price)
	order, err := qb.Do(context.Background())
	if err != nil {
		panic(err)
	}

	return order.UpdateTime - nowTs
}
