package main

import (
	pricepb "git.alibaba.ir/taraaz/salvation2/monorepo/pkg/protos/protogen/price_service"
	"google.golang.org/grpc"
	"log"
)

func newPriceServiceClient() pricepb.PriceServiceClient {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	return pricepb.NewPriceServiceClient(cc)
}
