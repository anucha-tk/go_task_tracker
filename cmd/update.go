package cmd

import (
	"errors"
	"strconv"

	"github.com/anucha-tk/task_tracker/internal/task"
	"github.com/anucha-tk/task_tracker/style"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task description with id",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New(style.ErrorStyle().Render("Task ID and Description is required!"))
		}

		idStr := args[0]
		id, err := strconv.ParseInt(idStr, 0, 0)
		if err != nil {
			return errors.New(style.ErrorStyle().Render("Task ID invalid type"))
		}
		description := args[1]
		return task.UpdateTask(id, description)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
