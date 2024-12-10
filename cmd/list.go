/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"angellisandroerazo/Task-Tracker-CLI/internal"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar las tareas.",
	Long:  `Se listan todas las tareas.`,
	Run:   RunList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func RunList(cmd *cobra.Command, args []string) {

	tasks := internal.OpenFile()

	var filteredTasks []internal.Task

	status := "all"

	if len(args) > 0 {
		status = strings.ToLower(args[0])
	}

	switch status {
	case "all":
		filteredTasks = tasks
	case "todo":
		for _, task := range tasks {
			if task.Status == "todo" {
				filteredTasks = append(filteredTasks, task)
			}
		}

	case "in-progress":
		for _, task := range tasks {
			if task.Status == "in-progress" {
				filteredTasks = append(filteredTasks, task)
			}
		}

	case "done":
		for _, task := range tasks {
			if task.Status == "done" {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}
	rows := [][]string{}
	for _, task := range filteredTasks {
		rows = append(rows, []string{
			strconv.Itoa(task.Id),
			task.Description,
			task.Status,
			task.CreatedAt,
			task.UpdatedAt,
		})
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#5bc0de"))).
		Headers("ID", "Task", "Status", "Created At", "Updated At").
		Rows(rows...)

	fmt.Println(internal.TextStyle.Render("Tasks: ", status))
	fmt.Println(t)
}
