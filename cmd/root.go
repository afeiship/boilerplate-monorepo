/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"ytbdown/misc"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ytbdown",
	Short: "Download mp3/mp4 by youtube-dl.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		srcUrl := args[0]
		isMp3, _ := cmd.Flags().GetBool("mp3")
		isMp4, _ := cmd.Flags().GetBool("mp4")

		if isMp3 {
			misc.YtbMp3(srcUrl)
		}

		if isMp4 {
			misc.YtbMp4(srcUrl)
		}
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ytbdown.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("mp3", "3", false, "Download mp3 music.")
	rootCmd.Flags().BoolP("mp4", "4", false, "Download mp4 video.")
}
