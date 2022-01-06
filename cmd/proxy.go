package cmd

import (
	"dorm-web/conf"
	"dorm-web/proxy"
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var proxyCmd = &cobra.Command{
	Use:   "gateway",
	Short: "代理网关",
	RunE: func(cmd *cobra.Command, args []string) error {
		config := &conf.Config{
			ConfPath: "./conf",
			ConfName: "server.yml",
		}
		conf, err := conf.ParseYamlConfig(config)
		if err != nil {
			fmt.Printf("解析配置失败\n")
			return err
		}
		reg := etcd.NewRegistry(registry.Addrs(conf.EtcdAddress))
		gateWayServer := proxy.NewGateWayServer(reg)

		go func() {
			_ = gateWayServer.Run()
		}()
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		return nil
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)
}
