package main

import (
	"errors"
	"fmt"
	"os"
	"spotify/internal/back"
	"spotify/internal/login"
	"spotify/internal/next"
	"spotify/internal/p"
	"spotify/internal/pause"
	"spotify/internal/play"
	"spotify/internal/queue"
	"spotify/internal/repeat"
	"spotify/internal/save"
	"spotify/internal/shuffle"
	"spotify/internal/status"
	"spotify/internal/unsave"
	"spotify/internal/update"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	// TODO: https://github.com/spf13/viper/pull/1064
	viper.AddConfigPath("$HOME")
	viper.SetConfigName(".spotify-cli")
	viper.SetConfigType("json")
	viper.SafeWriteConfig()
	viper.ReadInConfig()

	root := &cobra.Command{
		Use:               "spotify",
		Short:             "Spotify for the terminal 🎵",
		Version:           "1.6.0",
		PersistentPreRunE: promptUpdate,
	}

	root.AddCommand(back.NewCommand())
	root.AddCommand(login.NewCommand())
	root.AddCommand(next.NewCommand())
	root.AddCommand(p.NewCommand())
	root.AddCommand(pause.NewCommand())
	root.AddCommand(play.NewCommand())
	root.AddCommand(queue.NewCommand())
	root.AddCommand(repeat.NewCommand())
	root.AddCommand(save.NewCommand())
	root.AddCommand(shuffle.NewCommand())
	root.AddCommand(status.NewCommand())
	root.AddCommand(unsave.NewCommand())
	root.AddCommand(update.NewCommand())
	root.AddCommand(completionCmd)

	// Hide help command
	root.SetHelpCommand(&cobra.Command{Hidden: true})

	// Rename default flag descriptions
	root.Flags().BoolP("help", "h", false, "Help for Spotify CLI.")
	root.Flags().BoolP("version", "v", false, "Version for Spotify CLI.")

	root.Execute()
}

func promptUpdate(cmd *cobra.Command, _ []string) error {
	if time.Now().Unix() < viper.GetInt64("prompt_update_timer") {
		return nil
	}

	isUpdated, err := update.IsUpdated(cmd)
	if err != nil {
		return err
	}
	if !isUpdated {
		cmd.Println("Use 'spotify update' to get the latest version.")
	}

	// Wait one day before the next prompt
	const day int64 = 24 * 60 * 60
	viper.Set("prompt_update_timer", time.Now().Unix()+day)

	return viper.WriteConfig()
}

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
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
			if err := cmd.Root().GenZshCompletion(os.Stdout); err != nil {
				return err
			}
			if _, err := fmt.Fprintf(os.Stdout, "compdef _spotify spotify"); err != nil {
				return err
			}
			return nil
		case "fish":
			return cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			return cmd.Root().GenPowerShellCompletion(os.Stdout)
		default:
			return errors.New("provided shell type is not supported")
		}
	},
}
