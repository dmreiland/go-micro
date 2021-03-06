package handler

import (
	"code.google.com/p/go.net/context"

	log "github.com/golang/glog"
	"github.com/myodc/go-micro/server"
	example "github.com/myodc/go-micro/template/proto/example"
)

type Example struct{}

func (e *Example) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
	log.Info("Received Example.Call request")

	rsp.Msg = server.Id + ": Hello " + req.Name

	return nil
}
