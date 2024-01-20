package cmd

import (
	"fmt"
	"os"

	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".shortenurl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".shortenurl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	err := configx.LoadWithViper(viper.GetViper())
	cobra.CheckErr(err)

	err = logging.InitWithConfig(configx.C.Log)
	cobra.CheckErr(err)
}
