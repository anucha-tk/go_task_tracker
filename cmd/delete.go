package cmd

import (
	"errors"
	"strconv"

	"github.com/anucha-tk/task_tracker/internal/task"
	"github.com/anucha-tk/task_tracker/style"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete task with id",

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New(style.ErrorStyle().Render("Task ID is required"))
		}
		idStr := args[0]
		id, err := strconv.ParseInt(idStr, 0, 0)
		if err != nil {
			return errors.New(style.ErrorStyle().Render("Task ID invalid type"))
		}
		return task.DeleteTask(id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
