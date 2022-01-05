package cmd

import "github.com/spf13/cobra"

var studentCmd = &cobra.Command{
	Use: 	"student",
	Short: 	"学生服务",
	RunE: 	func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func init()  {
	rootCmd.AddCommand(studentCmd)
}
