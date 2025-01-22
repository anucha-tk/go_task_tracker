package cmd

import (
	"github.com/anucha-tk/task_tracker/internal/task"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `
  Example:
  task_tracker list todo
  task_tracker list in-progress
  task_tracker list done
  `,

	RunE: func(cmd *cobra.Command, args []string) error {
		var status task.TaskStatus = "all"
		if len(args) != 0 {
			status = task.TaskStatus(args[0])
		}
		return task.ListTasks(status)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
