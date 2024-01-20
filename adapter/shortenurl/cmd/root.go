package cmd

import (
	"os"

	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shortenurl",
	Short: "a shorten url application",
	Long: `
███████╗██╗  ██╗ ██████╗ ██████╗ ████████╗███████╗███╗   ██╗    ██╗   ██╗██████╗ ██╗     
██╔════╝██║  ██║██╔═══██╗██╔══██╗╚══██╔══╝██╔════╝████╗  ██║    ██║   ██║██╔══██╗██║     
███████╗███████║██║   ██║██████╔╝   ██║   █████╗  ██╔██╗ ██║    ██║   ██║██████╔╝██║     
╚════██║██╔══██║██║   ██║██╔══██╗   ██║   ██╔══╝  ██║╚██╗██║    ██║   ██║██╔══██╗██║     
███████║██║  ██║╚██████╔╝██║  ██║   ██║   ███████╗██║ ╚████║    ╚██████╔╝██║  ██║███████╗
╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝╚══════╝
                                                                                         
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.shortenurl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	err := configx.Load(cfgFile, "shortenurl")
	cobra.CheckErr(err)

	err = logging.InitWithConfig(configx.C.Log)
	cobra.CheckErr(err)
}
