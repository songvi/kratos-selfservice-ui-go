package cmd

import (
	"fmt"
	"os"

	"github.com/ory/viper"
	"github.com/spf13/cobra"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "kraui",
		Short: "serve open service",
		Long:  "Test cmd with cobra",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Main command")
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is  ./configs.json)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startCmd)
}

func er(msg interface{}) {
	fmt.Println("Error: ", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("configs")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())		
	} else {
		er(err)
	}
}
