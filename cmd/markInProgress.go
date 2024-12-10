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

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress [string]",
	Short: "A brief description of your command",
	Run:   RunInProgress,
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markInProgressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markInProgressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func RunInProgress(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		fmt.Println(internal.ErrorStyle.Render("Task ID is required."))
		return
	}

	tasks := internal.OpenFile()

	taskFound := false
	for i, task := range tasks {
		if strconv.Itoa(task.Id) == args[0] {
			tasks[i].Status = "in-progress"
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
		fmt.Println(internal.SuccessStyle.Render("Task changed to IN-PROGRESS."))
		return
	}

}
