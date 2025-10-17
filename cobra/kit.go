package boxcobra

import "github.com/spf13/cobra"

type CmdFunc func() *cobra.Command

type CmdFuncList []CmdFunc

func (l CmdFuncList) RegisterTo(root *cobra.Command) {
	for _, f := range l {
		root.AddCommand(f())
	}
}
