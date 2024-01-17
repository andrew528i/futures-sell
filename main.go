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
	//// deal with timestamp and sleep
	now := time.Now()
	//targetTime := time.Unix(TargetTimestamp, 0)
	//
	//if targetTime.Before(now) {
	//	fmt.Println("target timestamp has already been passed")
	//	return
	//}
	//
	//delta := targetTime.Sub(now)
	//fmt.Println("sleeping:", delta)
	//time.Sleep(delta)

	// place sell order
	client := binance.NewFuturesClient(ApiKey, ApiSecret)
	quantity := fmt.Sprintf("%d", OrderSize)
	price := fmt.Sprintf("%d", OrderPrice)
	order, err := client.NewCreateOrderService().
		Symbol(OrderSymbol).
		Side(futures.SideTypeSell).
		Type(futures.OrderTypeLimit).
		TimeInForce(futures.TimeInForceTypeGTC).
		Quantity(quantity).
		Price(price).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	// print time delta result
	orderUpdateTime := time.Unix(order.UpdateTime, 0)
	delta := orderUpdateTime.Sub(now) // targetTime
	fmt.Println("time difference in ms:", delta.Milliseconds(), order.UpdateTime, now.Unix())

	err = client.NewCancelAllOpenOrdersService().Symbol(OrderSymbol).Do(context.Background())
	if err != nil {
		panic(err)
	}
}
