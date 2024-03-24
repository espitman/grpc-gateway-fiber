package main

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

func getCtx(fctx fiberCtx) context.Context {
	reqHeaders := fctx.GetReqHeaders()

	headers := make(map[string]string)
	headers["authorization"] = reqHeaders["Authorization"][0]
	headers["trace-id"] = uuid.NewString()

	header := metadata.New(headers)
	ctx := metadata.NewOutgoingContext(fctx.Context(), header)
	return ctx
}
