/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [todo]",
	Short: "Add a todo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := loadTasks()
		if err != nil {
			log.Fatal(err)
		}

		newTask := Task{
			ID:       genNextId(tasks),
			Priority: priority,
			Name:     args[0],
			Done:     false,
		}

		tasks = append(tasks, newTask)
		err = saveTasks(tasks)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Added todo %s\n", newTask.Name)
	},
}

func init() {
	addCmd.Flags().IntVarP(&priority, "priority", "p", 1, "Priority of task")
	rootCmd.AddCommand(addCmd)
}
