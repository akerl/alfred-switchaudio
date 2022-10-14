package cmd

import (
	"fmt"
	"strings"

	"github.com/akerl/alfred-switchaudio/utils"

	"github.com/spf13/cobra"
)

var xml = "<item arg=\"%s\" uid=\"%s\"><title>%s</title><subtitle/><icon>icon.png</icon></item>\n"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List audio devices",
	RunE:  listRunner,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("direction", "d", "output", "input or output")
}

func listRunner(cmd *cobra.Command, _ []string) error {
	direction, err := cmd.Flags().GetString("direction")
	if err != nil {
		return err
	}

	devices, err := utils.AllOtherDevices(direction)
	if err != nil {
		return err
	}

	fmt.Println("<?xml version=\"1.0\"?><items>")
	for _, device := range devices {
		uid := strings.ReplaceAll(device, " ", "_")
		fmt.Printf(xml, device, uid, device)
	}
	fmt.Println("</items>")

	return nil
}
