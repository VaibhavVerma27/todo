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

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [todo ID]",
	Short: "Mark a todo done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := loadTasks()

		if err != nil {
			log.Fatal(err)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Invalid todo Id")
		}

		found := false

		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
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

		fmt.Println("Todo marked as done")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
