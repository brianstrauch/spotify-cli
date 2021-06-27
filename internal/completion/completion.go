package completion

import (
	"github.com/spf13/cobra"
	"os"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script.",
		Long: `To load completions:

Bash:

$ source <(spotify completion bash)

# To load completions for each session, execute once:
Linux:
  $ spotify completion bash > /etc/bash_completion.d/spotify
MacOS:
  $ spotify completion bash > /usr/local/etc/bash_completion.d/spotify

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ spotify completion zsh > "${fpath[1]}/_spotify"

# You will need to start a new shell for this setup to take effect.

Fish:

$ spotify completion fish | source

# To load completions for each session, execute once:
$ spotify completion fish > ~/.config/fish/completions/spotify.fish
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				return cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				return cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				return cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				return cmd.Root().GenPowerShellCompletion(os.Stdout)
			}
			return nil
		},
	}
}
