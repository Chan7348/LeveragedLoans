package subgraph

import (
	"limit/conf"

	"context"

	"github.com/machinebox/graphql"
)

var (
	ctx       context.Context
	client    *graphql.Client
	OrderChan chan LimitOrder
)

func Init() {
	ctx = context.Background()
	client = graphql.NewClient(conf.Config.Subgraph)
	OrderChan = make(chan LimitOrder, 1000)
}
