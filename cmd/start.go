package cmd

import (
	constants "RCTestSetup/Packages/Constants"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigPath string
var Watcher bool
var Username string
var Password string
var Name string
var Email string
var Virtual bool
var Deps bool

var start = &cobra.Command{
	Use:                   "start [flags] <App-Directory>",
	Short:                 "This initiates Rocket.Chat container and injects your app inside.",
	Long:                  "Initiates a pre-setup Rocket.Chat development environment for using Rocket.Chat apps using docker, installs your app inside and launches a browser window for you.\n" + constants.Red + "Docker must be running to use this." + constants.White,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		checkConfig(args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(start)
	initStartFlags()
	// start.SetContext()
}

func initStartFlags() {
	start.Flags().StringVarP(&ConfigPath, "config", "c", "./", constants.Blue+"Path of the configuration file for companion, if not found will use the default values."+constants.White)
	start.Flags().BoolVarP(&Watcher, "watcher", "w", true, constants.Blue+"Specify whether you want use hot-reloading"+constants.White)
	start.Flags().StringVarP(&Username, "username", "a", "user0", constants.Blue+"Admin Username for your rocket.chat server"+constants.White)
	start.Flags().StringVarP(&Email, "email", "e", "a@b.com", constants.Blue+"Admin Email for your rocket.chat server"+constants.White)
	start.Flags().StringVarP(&Password, "password", "p", "123456", constants.Blue+"Admin Username for your rocket.chat server"+constants.White)
	start.Flags().StringVarP(&Name, "name", "n", "user", constants.Blue+"Admin Name for your rocket.chat server"+constants.White)
	start.Flags().BoolVarP(&Virtual, "Virtual", "v", true, constants.Blue+"Mounts your app directory to a companion container where all the app dependencies are present like node, npm & rc-apps cli, false uses the local environment for the dependencies."+constants.White)
	start.Flags().BoolVarP(&Deps, "Deps", "d", true, constants.Blue+"Installs your app dependencies at the beginning, disable if you don't wanna run `npm i` on every run."+constants.White)
}

func checkConfig(appDir string) {
	viper.GetViper().AddConfigPath(ConfigPath)
	viper.GetViper().SetConfigFile(".rc.yml")
	err := viper.GetViper().ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println(constants.Yellow + "Config File Not Provided, generating one with default configuration for you in the given directory" + constants.Yellow)
			generateDefaultConfig(appDir)

		} else {
			fmt.Println(constants.Red + "Woops, Something went wrong :(" + constants.White)
		}
	}
}

func generateDefaultConfig(appDir string) {

	fmt.Println(appDir)

	viper.SetDefault("configPath", "./")
	viper.SetDefault("appDir", appDir)
	viper.SetDefault("admin.username", Username)
	viper.SetDefault("admin.email", Email)
	viper.SetDefault("admin.password", Password)
	viper.SetDefault("admin.name", Name)
	viper.SetDefault("watcher.watcher", Watcher)
	viper.SetDefault("watcher.mode", "deep")
	viper.SetDefault("virtual", true)
	viper.SetDefault("composeFilePath", "./")
	viper.SetDefault("installDependencies", Deps)

	viper.WriteConfigAs(".rc.yml")
}
