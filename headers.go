package main

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func getCtx(fctx fiberCtx) context.Context {
	reqHeaders := fctx.GetReqHeaders()
	header := metadata.New(map[string]string{"authorization": reqHeaders["Authorization"][0]})
	ctx := metadata.NewOutgoingContext(fctx.Context(), header)
	return ctx
}
