package cmd

import (
	"fmt"

   	"github.com/spf13/cobra"
	"github.com/airlanggatirta/pawon-warga/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
        Use:   "version",
        Short: "Print the version number of pawon-warga",
        Long:  `All software has versions. This is pawon-warga`,
        Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Build Date:", version.BuildDate)
			fmt.Println("Git Commit:", version.GitCommit)
			fmt.Println("Version:", version.Version)
			fmt.Println("Environment:", version.Environment)
			fmt.Println("Go Version:", version.GoVersion)
			fmt.Println("OS / Arch:", version.OsArch)
        },
}

func init() {
        rootCmd.AddCommand(versionCmd)
}
