package cmd

import (
	"errors"
	"strconv"

	"github.com/anucha-tk/task_tracker/internal/task"
	"github.com/anucha-tk/task_tracker/style"
	"github.com/spf13/cobra"
)

var markInprogressCmd = &cobra.Command{
	Use:   "mark-inprogress",
	Short: "Make task status to in-progress",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New(style.ErrorStyle().Render("Task ID is required."))
		}
		strId := args[0]
		id, err := strconv.ParseInt(strId, 10, 64)
		if err != nil {
			return errors.New(style.ErrorStyle().Render("Invalid id type"))
		}
		return task.UpdateStatus(task.Task_Status_IN_PROGRESS, id)
	},
}

func init() {
	rootCmd.AddCommand(markInprogressCmd)
}
