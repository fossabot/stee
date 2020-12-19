package command

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCommand = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate autocompletion scripts for shells",
	Long: `To load completions:

	Bash:
	
	$ source <(stee completion bash)
	
	# To load completions for each session, execute once:
	Linux:
	  $ stee completion bash > /etc/bash_completion.d/stee
	MacOS:
	  $ stee completion bash > /usr/local/etc/bash_completion.d/stee
	
	Zsh:
	
	# If shell completion is not already enabled in your environment you will need
	# to enable it.  You can execute the following once:
	
	$ echo "autoload -U compinit; compinit" >> ~/.zshrc
	
	# To load completions for each session, execute once:
	$ stee completion zsh > "${fpath[1]}/_stee"
	
	# You will need to start a new shell for this setup to take effect.
	
	Fish:
	
	$ stee completion fish | source
	
	# To load completions for each session, execute once:
	$ stee completion fish > ~/.config/fish/completions/stee.fish
	`,
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
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}

func init() {
	rootCommand.AddCommand(completionCommand)
}
