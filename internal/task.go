package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

func OpenFile() []Task {
	file, err := os.ReadFile("tasks.json")
	if err != nil {

		fmt.Println(ErrorStyle.Render("tasks.json file does not exist."))

		file, err := os.Create("tasks.json")
		if err != nil {
			fmt.Println("Error creating file: ", err)
			return nil
		}

		fmt.Println(SuccessStyle.Render("File created successfully."))

		defer file.Close()

	}

	tasks := []Task{}

	json.Unmarshal(file, &tasks)

	return tasks
}

func SaveFile(tasks []Task) bool {
	newData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(ErrorStyle.Render("Error converting data to JSON: ", err.Error()))
		return false
	}

	err = os.WriteFile("tasks.json", newData, 0644)
	if err != nil {
		fmt.Println(ErrorStyle.Render("Error writing file: ", err.Error()))
		return false
	}

	return true
}
