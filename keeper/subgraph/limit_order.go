package subgraph

import (
	"log"
	"time"

	"github.com/machinebox/graphql"
	"github.com/shopspring/decimal"
)

type LimitOrder struct {
	Id         string
	Owner      string
	Account    string
	Keeper     string
	IsLong     bool
	IsCanceled bool
	OrderId    uint
	ExpireAt   uint
	CreatedAt  uint
	UpdatedAt  uint
	OpenSize   decimal.Decimal
	LimitPrice decimal.Decimal
	DealPrice  decimal.Decimal
}

type QueryLimitOrdersResp struct {
	LimitOrders []LimitOrder
}

func QueryLimitOrders(keeper string) []LimitOrder {
	currentTime := time.Now().Unix()
	req := graphql.NewRequest(`
	query ($keeper: String!, $currentTime: BigInt) {
		limitOrders(first: 1000, orderBy: expireAt,
			where: {expireAt_gt: $currentTime, keeper: $keeper, isCanceled: false, dealPrice: 0}) {
			id
			owner
			account
			keeper
			isLong
			isCanceled
			orderId
			expireAt
			openSize
			limitPrice
			dealPrice
			createdAt
			updatedAt
		}
	}
	`)
	req.Var("keeper", keeper)
	req.Var("currentTime", currentTime)

	var respData QueryLimitOrdersResp
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Print(err)
	}
	limitOrders := respData.LimitOrders
	return limitOrders
}