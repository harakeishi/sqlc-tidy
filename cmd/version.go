package cmd

import (
	"github.com/harakeishi/curver"
	"github.com/spf13/cobra"
)

// versionCmdは、アプリケーションのバージョンを表示するコマンドです。
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the application",
	Long:  `This is the application's version.`,
	Run: func(cmd *cobra.Command, args []string) {
		curver.EchoVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
