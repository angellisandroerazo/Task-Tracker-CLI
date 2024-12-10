/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"angellisandroerazo/Task-Tracker-CLI/internal"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [string] [string]",
	Short: "Actualizar una tarea.",
	Long:  `Cambia el nombre de una tarea.`,
	Run:   RunUpdate,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func RunUpdate(cmd *cobra.Command, args []string) {

	if len(args) != 2 {
		fmt.Println(internal.ErrorStyle.Render("Task ID and new description required."))
		return
	}

	tasks := internal.OpenFile()

	taskFound := false
	for i, task := range tasks {
		if strconv.Itoa(task.Id) == args[0] {
			tasks[i].Description = args[1]
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			taskFound = true
			break
		}
	}

	if !taskFound {
		fmt.Println(internal.ErrorStyle.Render("Task not found with ID."))
		return
	}

	ok := internal.SaveFile(tasks)

	if ok {
		fmt.Println(internal.SuccessStyle.Render("Task updated successfully"))
		return
	}

}
