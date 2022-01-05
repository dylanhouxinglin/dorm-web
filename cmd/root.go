package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: 	"root",
}

func Exe()  {
	_ = rootCmd.Execute()
}