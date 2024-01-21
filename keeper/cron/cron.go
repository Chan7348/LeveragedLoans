package cron

import (
	"limit/contract"
	"limit/subgraph"

	"log"

	"github.com/robfig/cron/v3"
	"github.com/shopspring/decimal"
)

var latestPrice decimal.Decimal

func Init() {
	log.Println("Starting cron ....")
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 */12 * * * ?", updateLatestPrice)
	c.AddFunc("0 */12 * * * ?", handleLimitOrders)
	c.Start()
}

func updateLatestPrice() {
	log.Printf("Ready to get latestPrice")
	latestPrice = contract.GetLatestPrice()
	log.Printf("price: %d", latestPrice)
}

func handleLimitOrders() {
	keeper := contract.GetWalletAddress()
	orders := subgraph.QueryLimitOrders(keeper)
	for _, order := range orders {
		log.Printf("Order: %+v", order)
		if order.IsLong {
			if order.LimitPrice.Cmp(latestPrice) >= 0 {
				subgraph.OrderChan <- order
			}
		} else {
			if order.LimitPrice.Cmp(latestPrice) <= 0 {
				subgraph.OrderChan <- order
			}
		}
	}
}
