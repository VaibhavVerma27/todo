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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [todo ID]",
	Short: "Delete a todo",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := loadTasks()
		if err != nil {
			log.Fatal(err)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("invalid todo ID")
		}

		newTasks := []Task{}
		for _, task := range tasks {
			if task.ID != id {
				newTasks = append(newTasks, task)
			}
		}

		err = saveTasks(newTasks)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Todo deleted successfully: %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
