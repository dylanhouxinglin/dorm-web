package common

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"golang.org/x/sync/errgroup"
	"time"
)

type CommonServer struct {
	*gin.Engine
	stopCh 			chan struct{}
	// 服务注册用
	Addr			string
	registry		registry.Registry
	ctx				context.Context
}

func NewGinServer(addr string, reg registry.Registry) *CommonServer {
	server := &CommonServer{
		Engine:   gin.New(),
		stopCh:   make(chan struct{}),
		Addr:     addr,
		registry: reg,
		ctx: 	  context.Background(),
	}
	return server
}

func (cs *CommonServer) RegisterUrl()  {

}

func (cs *CommonServer) Run() error {
	group, errCtx := errgroup.WithContext(cs.ctx)
	// 注册到etcd
	group.Go(func() error {
		ticker := time.NewTicker(2 * time.Second)
		for {
			select {
			case <- errCtx.Done():
				return nil
			case <- ticker.C:
				if err := cs.EtcdRegister(); err != nil {
					fmt.Printf("注册到etcd失败 %s\n", err)
					return err
				}
			}
		}
	})
	// 启动grpc服务

	return group.Wait()
}

func (cs *CommonServer) EtcdRegister() error {
	srv := &registry.Service{
		Name:      "common",
		Nodes:     []*registry.Node{
			{
				Id:			cs.Addr,
				Address:	cs.Addr,
				Metadata:	map[string]string{"type":"common"},
			},
		},
	}
	return cs.registry.Register(srv, registry.RegisterTTL(5*time.Second))
}