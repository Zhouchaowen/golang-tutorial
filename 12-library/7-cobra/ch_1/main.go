package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:  "say",
	Long: "Root command",
}

var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello!!!")
	},
}

var ByeCmd = &cobra.Command{
	Use:   "bye",
	Short: "Say goodbye",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bye!!!")
	},
}

func init() {
	RootCmd.AddCommand(HelloCmd, ByeCmd, CompletionCmd)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var CompletionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate completion script",
	Long:                  "To load completions",
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}
