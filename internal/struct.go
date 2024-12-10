package internal

import "github.com/charmbracelet/lipgloss"

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

var ErrorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#bb2124")).
	Padding(1).
	Align(lipgloss.Center)

var SuccessStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#22bb33")).
	Padding(1).
	Align(lipgloss.Center)

var TextStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#f0ad4e")).
	Padding(1).
	Align(lipgloss.Center)
