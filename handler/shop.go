package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	shop "shop/proto"
)

type Shop struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Shop) Call(ctx context.Context, req *shop.Request, rsp *shop.Response) error {
	log.Info("Received Shop.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Shop) Stream(ctx context.Context, req *shop.StreamingRequest, stream shop.Shop_StreamStream) error {
	log.Infof("Received Shop.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&shop.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Shop) PingPong(ctx context.Context, stream shop.Shop_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&shop.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
