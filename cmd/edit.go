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

var newPriority int

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [todo ID] [new name]",
	Short: "Edit a todo",
	Args:  cobra.ExactArgs(2),
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
				tasks[i].Name = args[1]
				if newPriority > 0 {
					tasks[i].Priority = newPriority
				}
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

		fmt.Println("Todo edited successfully")
	},
}

func init() {
	editCmd.Flags().IntVarP(&newPriority, "priority", "p", 0, "New priority")
	rootCmd.AddCommand(editCmd)
}
