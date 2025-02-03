package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var name string

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "prints hello world",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			fmt.Println("Hello world!")
		} else {
			fmt.Printf("Hello %s!\n", name)
		}
	},
}

func init() {
	helloCmd.Flags().StringVarP(&name, "name", "n", "", "your name")

	rootCmd.AddCommand(helloCmd)
}
