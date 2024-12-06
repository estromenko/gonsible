/*
Copyright © 2024 Eduard Stromenko estromenko@mail.ru
*/
package cmd

import (
	"github.com/estromenko/gonsible/internal/inventory"
	"github.com/estromenko/gonsible/internal/pipeline"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <my-pipeline.toml>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		invent, err := inventory.New(inventoryPath)
		if err != nil {
			panic(err)
		}

		pipe, err := pipeline.New(args[0])
		if err != nil {
			panic(err)
		}

		if err := pipe.Execute(invent); err != nil {
			panic(err)
		}
	},
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
