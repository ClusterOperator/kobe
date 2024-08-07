package root

import (
	"github.com/ClusterOperator/kobe/cmd/client/adhoc"
	"github.com/ClusterOperator/kobe/cmd/client/playbook"
	"github.com/ClusterOperator/kobe/cmd/client/project"
	"github.com/ClusterOperator/kobe/cmd/client/task"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "kobe",
	Short: "A kobe client cli",
}

func init() {
	Cmd.AddCommand(project.Cmd)
	Cmd.AddCommand(playbook.Cmd)
	Cmd.AddCommand(task.Cmd)
	Cmd.AddCommand(adhoc.Cmd)
}
