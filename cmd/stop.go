package cmd

import (
	"fmt"
	constants "thrust/Packages/Constants"

	"github.com/spf13/cobra"
)

var rma bool
var rmv bool
var rmc bool

var stopCommand = &cobra.Command{
	Use:                   "stop [Flags] <directory>",
	Short:                 "This stops all the containers responsible for running the environment",
	Long:                  "This command stops and optionally removes containers and volumes such as Rocket.Chat mongodb etc responsible for execution of the companion, if removed the companion has to download them again.",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Function Ran")
	},
}

func init() {
	// rootCmd.AddCommand(stopCommand)
	initStopFlags()
}

func initStopFlags() {
	stopCommand.Flags().BoolVarP(&rmv, "remove-images", "v", false, constants.Blue+"This removes all volumes that were generated while starting the companion"+constants.White)
	stopCommand.Flags().BoolVarP(&rmc, "remove-containers", "c", true, constants.Blue+"This removes all the images that were generated while starting the companion"+constants.White)
	stopCommand.Flags().BoolVarP(&rma, "remove-all", "a", false, constants.Blue+"This removes both the images and volumes that were generated while starting the companion"+constants.White)
}
