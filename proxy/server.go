package proxy

import (
	"context"
	"dorm-web/utils"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"golang.org/x/sync/errgroup"
	"sync"
)

type GinServer struct {
	*gin.Engine
	sync.RWMutex
	stopCh 				chan struct{}
	serviceName			[]string
	services			map[string]*utils.Engine
	pendingEngine		map[string][]*utils.Engine
	registry			registry.Registry
}

func NewGateWayServer(reg registry.Registry) *GinServer {
	server := &GinServer{
		Engine: 	 gin.New(),
		stopCh: 	 make(chan struct{}),
		serviceName: []string{"common"},
		registry: 	 reg,
	}
	server.RegisterUrl()
	return server
}

func (gs *GinServer) RegisterUrl() {
	gs.GET("/ready", func(ctx *gin.Context) {
		ctx.String(200, "success")
	})
}

func (gs *GinServer) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<- gs.stopCh
		cancel()
	}()

	group, _ := errgroup.WithContext(ctx)
	// web服务
	group.Go(func() error {
		err := endless.ListenAndServe(":8976", gs)
		return err
	})
	// 监听服务
	group.Go(func() error {

	})

	return group.Wait()
}

func (gs *GinServer) EtcdRegister() error {
	return nil
}

func (gs *GinServer) WatchDiscovery(targetService string, ctx context.Context) error {
	services, err := gs.registry.GetService(targetService)
	if err != nil && err != registry.ErrNotFound {
		fmt.Printf("监听服务 %s 失败\n", targetService)
		return err
	}
	fmt.Printf("监听 %s 中...\n", targetService)
	for _, sv := range services {
		for _, engine := range sv.Nodes {
			gs.AddEngine(targetService, engine)
		}
	}
}

func (gs *GinServer) AddEngine(service string, engine *registry.Node)  {
	gs.RLock()
	if _, ok := gs.services[service]; !ok {
		if _, ok := gs.pendingEngine[service]; ok {
			//
		} else {

		}
	} else {

	}
	gs.RUnlock()


}