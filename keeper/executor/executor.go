package executor

import (
	"limit/contract"
	"limit/subgraph"

	"log"
)

func Run() {
	log.Println("executor running")
	for {
		order := <-subgraph.OrderChan
		log.Printf("receive order: %+v", order)
		tx, err := contract.ExecuteLimitOrder(order.Account, order.OrderId)
		if err == nil {
			log.Printf("ExecuteLimitOrder tx: %s", tx)
		}
	}
}
