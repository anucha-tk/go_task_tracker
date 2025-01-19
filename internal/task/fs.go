package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/anucha-tk/task_tracker/style"
	"github.com/charmbracelet/lipgloss"
)

func tasksFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error get current work directory: %v\n", err)
		return ""
	}
	return path.Join(cwd, "tasks.json")
}

func ReadTasksFormFile() ([]Task, error) {
	filePath := tasksFilePath()

	_, err := os.Stat(filePath)

	// tasks.json is not exist
	if os.IsNotExist(err) {
		greenStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
		fmt.Println(greenStyle.Render("Creating task.json..."))

		_, err := os.Create(filePath)
		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())
		if err != nil {
			errMsg := style.ErrorStyle().Render(fmt.Sprintf("Error creating tasks.json: %v", err))
			fmt.Print(errMsg)
			return nil, err
		}

		return []Task{}, nil
	}

	// tasks.json is exist
	file, err := os.Open(filePath)
	if err != nil {
		errMsg := style.ErrorStyle().Render(fmt.Sprintf("Error can't open tasks.json: %v", err))
		fmt.Print(errMsg)
		return nil, err
	}

	tasks := []Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		errMsg := style.ErrorStyle().Render(fmt.Sprintf("Error can't decode tasks.json: %v", err))
		fmt.Print(errMsg)
		return nil, err
	}

	return tasks, nil
}

func WriteTaskToFile(tasks []Task) error {
	filePath := tasksFilePath()
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return err
	}

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		fmt.Printf("Error encoder file: %v\n", err)
		return err
	}

	return nil
}
