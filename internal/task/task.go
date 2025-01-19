package task

import (
	"fmt"
	"time"

	"github.com/anucha-tk/task_tracker/style"
)

type TaskStatus string

const (
	Task_Status_TODO        TaskStatus = "todo"
	Task_Status_IN_PROGRESS TaskStatus = "in-progress"
	Task_Status_DONE        TaskStatus = "done"
)

type Task struct {
	ID          int64      `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func NewTask(id int64, description string) Task {
	return Task{
		ID:          id,
		Description: description,
		Status:      Task_Status_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func AddTask(description string) error {
	tasks, err := ReadTasksFormFile()
	if err != nil {
		return err
	}

	var id int64
	if len(tasks) == 0 {
		id = 1
	} else {
		id = int64(len(tasks) + 1)
	}

	task := NewTask(id, description)
	tasks = append(tasks, task)
	err = WriteTaskToFile(tasks)
	if err != nil {
		errMsg := style.ErrorStyle().Render(fmt.Sprintf("Error can't add task: %v", err))
		fmt.Print(errMsg)
		return err
	}
	successMsg := style.SuccessStyle().Render(fmt.Sprintf("Add task successful, id: %d", id))
	fmt.Println(successMsg)

	return nil
}
