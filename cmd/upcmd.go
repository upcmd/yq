package cmd

import (
	"gopkg.in/op/go-logging.v1"
	"os"

	"github.com/spf13/cobra"
)

func UpCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "yq",
		Short: "yq is a lightweight and portable command-line YAML processor.",
		Long:  `yq is a lightweight and portable command-line YAML processor. It aims to be the jq or sed of yaml files.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var format = logging.MustStringFormatter(
				`%{color}%{time:15:04:05} %{shortfunc} [%{level:.4s}]%{color:reset} %{message}`,
			)
			var backend = logging.AddModuleLevel(
				logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), format))

			if verbose {
				backend.SetLevel(logging.DEBUG, "")
			} else {
				backend.SetLevel(logging.ERROR, "")
			}

			logging.SetBackend(backend)
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose mode")
	rootCmd.PersistentFlags().BoolVarP(&outputToJSON, "tojson", "j", false, "output as json. By default it prints a json document in one line, use the prettyPrint flag to print a formatted doc.")
	rootCmd.PersistentFlags().BoolVarP(&prettyPrint, "prettyPrint", "P", false, "pretty print")
	rootCmd.PersistentFlags().IntVarP(&indent, "indent", "I", 2, "sets indent level for output")
	rootCmd.Flags().BoolVarP(&version, "version", "V", false, "Print version information and quit")
	rootCmd.PersistentFlags().BoolVarP(&colorsEnabled, "colors", "C", false, "print with colors")

	//p("verbose:", rootCmd.Flag("verbose").Value)
	rootCmd.AddCommand(
		CreateUpReadCmd(),
		createUpValidateCmd(),
		createUpDeleteCmd(),
	)
	rootCmd.SetOutput(os.Stdout)

	return rootCmd
}
