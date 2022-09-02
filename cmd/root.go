package cmd

import (
	"fmt"
	"github.com/afeiship/nx"
	"os"
	"ytbdown/misc"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ytbdown",
	Short: "Download mp3/mp4 by youtube-dl/yt-dlp.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var mpx any
		srcUrl := args[0]
		isMp3, _ := cmd.Flags().GetBool("mp3")
		isMp4, _ := cmd.Flags().GetBool("mp4")
		isKeep, _ := cmd.Flags().GetBool("keep")
		name, _ := cmd.Flags().GetString("name")

		if isMp3 {
			mpx = nx.If(isMp3, "mp3", nil)
		}

		if isMp4 {
			mpx = nx.If(isMp4, "mp4", nil)
		}

		mps := fmt.Sprintf("%v", mpx)
		misc.YtbMpx(mps, srcUrl, isKeep, name)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("name", "n", "%(title)s", "Download filename.")
	rootCmd.Flags().BoolP("mp3", "3", false, "Download mp3 music.")
	rootCmd.Flags().BoolP("mp4", "4", false, "Download mp4 video.")
	rootCmd.Flags().BoolP("keep", "k", false, "Download mpx keep original file.")
}
