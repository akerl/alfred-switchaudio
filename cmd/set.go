package cmd

import (
	"github.com/akerl/alfred-switchaudio/utils"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set audio device",
	RunE:  setRunner,
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("direction", "d", "output", "input or output")
	setCmd.Flags().StringP("target", "t", "", "target device name")
}

func setRunner(cmd *cobra.Command, _ []string) error {
	flags := cmd.Flags()

	direction, err := flags.GetString("direction")
	if err != nil {
		return err
	}

	target, err := flags.GetString("target")
	if err != nil {
		return err
	}

	return utils.SetDevice(direction, target)
}
