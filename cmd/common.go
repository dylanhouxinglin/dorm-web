package cmd

import "github.com/spf13/cobra"

var commonCmd = &cobra.Command{
	Use: 	"common",
	Short: 	"通用服务",
	RunE: 	func(cmd *cobra.Command, args []string) error {
		return nil
	},

}

func init()  {
	rootCmd.AddCommand(commonCmd)
}
