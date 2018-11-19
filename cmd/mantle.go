package cmd

import (
	"mantle/pkg/init"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "mantle",
	Short:         "easier kube objects",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(_ *cobra.Command, args []string) error {
		return init.MantleInit()
	},
}
