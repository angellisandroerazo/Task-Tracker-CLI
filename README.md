# Task Traker CLI

Task Tracker is a command-line tool designed to help you efficiently manage your tasks. Whether you're keeping track of personal to-dos or organizing work-related projects, Task Tracker provides a simple and intuitive way to create, update, filter, and delete tasks directly from the terminal.

## Requirements
- Go programming language installed on your machine.
- Git for version control.

## Getting Started
1. Clone the repository:
   ```bash
   git clone https://github.com/angellisandroerazo/task-tracker-cli.git
   ```

2. Navigate to the project directory:
   ```bash
   cd task-tracker-cli
   ```

3. Build the CLI application:
   ```bash
   go build -o task-cli
   ```
   ## Usage
To use the task tracker, you can execute commands like:
- Adding a new task:
  ```bash
  ./task-cli add "Buy groceries"
  ```

- Listing all tasks:
  ```bash
  ./task-cli list
  ```

- Listing tasks by status:
  ```bash
  ./task-cli list [todo, in-progress, done]
  ```

- Updating a task:
  ```bash
  ./task-cli update 1 "New task description"
  ```

- Deleting a task:
  ```bash
  ./task-cli delete 1
  ```

- Marking a task as done:
  ```bash
  ./task-cli mark-done 1
  ```

- Marking a task as in progress:
  ```bash
  ./task-cli mark-in-progress 1
  ```

## Problem Statement
This project addresses a task management problem inspired by the challenges outlined in the [Task Tracker roadmap](https://roadmap.sh/projects/task-tracker).