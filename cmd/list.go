/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"sort"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := loadTasks()

		if err != nil {
			log.Fatal(err)
		}

		if len(tasks) == 0 {
			fmt.Println("No Todos available.")
			return
		}

		sort.Slice(tasks, func(i, j int) bool { return tasks[i].Priority > tasks[j].Priority })

		for _, task := range tasks {
			status := "❌"
			if task.Done {
				status = "✅"
			}

			fmt.Printf("%s %d: %s (Priority: %d)\n", status, task.ID, task.Name, task.Priority)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
