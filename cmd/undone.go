/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone [todo ID]",
	Short: "Mark a todo undone",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := loadTasks()
		if err != nil {
			log.Fatal(err)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Invalid todo ID")
		}

		found := false
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = false
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Todo not found")
			return
		}

		err = saveTasks(tasks)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Todo successfully marked as undone")
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)
}
