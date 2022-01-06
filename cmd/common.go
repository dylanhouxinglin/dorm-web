package cmd

import (
	"dorm-web/common"
	"dorm-web/conf"
	"dorm-web/utils"
	"errors"
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/spf13/cobra"
	"github.com/micro/go-micro/registry/etcd"
	"log"
)

var commonCmd = &cobra.Command{
	Use: 	"common",
	Short: 	"通用服务",
	RunE: 	func(cmd *cobra.Command, args []string) error {
		host := utils.AcquireLocalHost()
		if host == "" || len(host) == 0 {
			fmt.Printf("获取本机host失败\n")
			return errors.New("获取本机host失败")
		}

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
		commonServer := common.NewGinServer(host+":8977", reg)
		return nil
	},

}

func init()  {
	rootCmd.AddCommand(commonCmd)
}
