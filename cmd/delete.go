/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"angellisandroerazo/Task-Tracker-CLI/internal"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [string]",
	Short: "Borrar una tarea.",
	Long:  `Borra una tarea de la lista de tareas.`,
	Run:   RunDelete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func RunDelete(cmd *cobra.Command, args []string) {

	if len(args) != 1 {
		fmt.Println(internal.ErrorStyle.Render("Task ID is required."))
		return
	}

	tasks := internal.OpenFile()

	updateTasks := []internal.Task{}
	taskFound := false
	for _, task := range tasks {
		if strconv.Itoa(task.Id) != args[0] {
			updateTasks = append(updateTasks, task)
			taskFound = true
			break
		}
	}

	if !taskFound {
		fmt.Println(internal.ErrorStyle.Render("Task not found with ID."))
		return
	}

	ok := internal.SaveFile(updateTasks)

	if ok {
		fmt.Println(internal.SuccessStyle.Render("Task delete successfully."))
	}
}
