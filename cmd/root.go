package cmd

import "github.com/spf13/cobra"
import "fmt"
import "os"

var rootCmd = &cobra.Command {
	Short: "Simple archiver",
}

func Execute() {
	if err:= rootCmd.Execute(); err!=nil {
		_,_ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func handleErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}