package cmd

import (
	"github.com/anucha-tk/task_tracker/internal/task"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",

	RunE: func(cmd *cobra.Command, args []string) error {
		return task.ListTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
