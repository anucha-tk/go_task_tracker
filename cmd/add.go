package cmd

import (
	"errors"

	"github.com/anucha-tk/task_tracker/internal/task"
	"github.com/anucha-tk/task_tracker/style"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New(style.ErrorStyle().Render("Task Description is required!"))
		}

		description := args[0]
		return task.AddTask(description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
