package cmd

import (
	"mantle/pkg/initialize"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "pulsar",
	Short:         "deploys and manages apache pulsar",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(_ *cobra.Command, args []string) error {
		return initialize.MantleInit()
	},
}
