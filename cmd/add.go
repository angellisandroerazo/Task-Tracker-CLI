/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"angellisandroerazo/Task-Tracker-CLI/internal"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a task.",
	Run:   RunAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func RunAdd(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		fmt.Println(internal.ErrorStyle.Render("Task description is required."))
		return
	}

	tasks := internal.OpenFile()

	task := internal.Task{
		Id:          len(tasks) + 1,
		Description: strings.Join(args, " "),
		Status:      "todo",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	tasks = append(tasks, task)

	ok := internal.SaveFile(tasks)

	if ok {
		fmt.Println(internal.SuccessStyle.Render("Task added successfully (ID: ", strconv.Itoa(len(tasks)), ")."))
		return
	}

}
