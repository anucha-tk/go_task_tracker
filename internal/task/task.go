package task

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/anucha-tk/task_tracker/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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
	successMsg := style.SuccessStyle().Render(fmt.Sprintf("✏️  Add task successful, id: %d", id))
	fmt.Println(successMsg)

	return nil
}

func UpdateTask(id int64, description string) error {
	tasks, err := ReadTasksFormFile()
	if err != nil {
		errMsg := style.ErrorStyle().Render(fmt.Sprintf("Error can't read task from file: %v\n", err))
		return fmt.Errorf(errMsg)
	}

	taskExist := false
	for i := range tasks {
		if tasks[i].ID == id {
			taskExist = true
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			break
		}
	}
	if !taskExist {
		errMsg := style.ErrorStyle().Render(fmt.Sprintf("Task id %d not found", id))
		return fmt.Errorf(errMsg)
	}

	successMsg := style.SuccessStyle().Render(fmt.Sprintf("🔄 Update task successful, id: %d, %s", id, description))
	fmt.Println(successMsg)
	return WriteTaskToFile(tasks)
}

func DeleteTask(id int64) error {
	tasks, err := ReadTasksFormFile()
	if err != nil {
		errMsg := style.ErrorStyle().Render(fmt.Sprintf("Error can't read task from file: %v\n", err))
		return fmt.Errorf(errMsg)
	}

	taskExist := false
	for i, task := range tasks {
		if task.ID == id {
			taskExist = true
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}

	if !taskExist {
		errMsg := style.ErrorStyle().Render(fmt.Sprintf("Task id %d not found", id))
		return fmt.Errorf(errMsg)
	}

	successMsg := style.SuccessStyle().Render(fmt.Sprintf("🗑️ Delete task id:%d successful", id))
	fmt.Println(successMsg)
	return WriteTaskToFile(tasks)
}

func ListTasks(status TaskStatus) error {
	tasks, err := ReadTasksFormFile()
	if err != nil {
		errMsg := style.ErrorStyle().Render(fmt.Sprintf("Error can't read task from file: %v\n", err))
		return fmt.Errorf(errMsg)
	}

	// filter
	filterTasks := []Task{}
	for _, task := range tasks {
		switch status {
		case "all":
			filterTasks = tasks
		case Task_Status_TODO, Task_Status_IN_PROGRESS, Task_Status_DONE:
			if task.Status == status {
				filterTasks = append(filterTasks, task)
			}
		}
		if status == "all" {
			break
		}
	}

	// convert []Task to [][]string
	rows := make([][]string, len(filterTasks))
	for i, task := range filterTasks {
		rows[i] = []string{
			fmt.Sprintf("%d", task.ID),
			task.Description,
			string(task.Status),
			task.CreatedAt.Format("2006-01-02 15:04:05"),
			task.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	headers := []string{"ID", "Description", "Status", "CreatedAt", "UpdatedAt"}
	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Padding(0, 1)
	headerStyle := baseStyle.Foreground(lipgloss.Color("252")).Bold(true)
	CapitalizeHeaders := func(data []string) []string {
		for i := range data {
			data[i] = strings.ToUpper(data[i])
		}
		return data
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).
		Headers(CapitalizeHeaders(headers)...).
		Width(80).
		Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}
			return baseStyle.Foreground(lipgloss.Color("252"))
		})

	fmt.Println(t)

	return nil
}
