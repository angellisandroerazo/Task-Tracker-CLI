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

// markDoneCmd represents the markDone command
var markDoneCmd = &cobra.Command{
	Use:   "mark-done [string]",
	Short: "Cambiar a tarea completada.",
	Run:   RunDone,
}

func init() {
	rootCmd.AddCommand(markDoneCmd)
}

func RunDone(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		fmt.Println(internal.ErrorStyle.Render("Task ID is required."))
		return
	}

	tasks := internal.OpenFile()

	taskFound := false
	for i, task := range tasks {
		if strconv.Itoa(task.Id) == args[0] {
			tasks[i].Status = "done"
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
		fmt.Println(internal.SuccessStyle.Render("Task changed to DONE."))
		return
	}

}
