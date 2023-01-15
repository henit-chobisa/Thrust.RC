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
var WatcherMode string
var composeFilePath string

var start = &cobra.Command{
	Use:                   "start [flags] <App-Directory>",
	Short:                 "This initiates Rocket.Chat container and injects your app inside.",
	Long:                  "Initiates a pre-setup Rocket.Chat development environment for using Rocket.Chat apps using docker, installs your app inside and launches a browser window for you.\n" + constants.Red + "Docker must be running to use this." + constants.White,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		checkConfig(args[0])
		fmt.Println(viper.GetString("config"))
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
	start.Flags().BoolVarP(&Watcher, "watcher.watcher", "w", true, constants.Blue+"Specify whether you want use hot-reloading"+constants.White)
	start.Flags().StringVarP(&Username, "admin.username", "a", "user0", constants.Blue+"Admin Username for your rocket.chat server"+constants.White)
	start.Flags().StringVarP(&Email, "admin.email", "e", "a@b.com", constants.Blue+"Admin Email for your rocket.chat server"+constants.White)
	start.Flags().StringVarP(&Password, "admin.password", "p", "123456", constants.Blue+"Admin Username for your rocket.chat server"+constants.White)
	start.Flags().StringVarP(&Name, "admin.name", "n", "user", constants.Blue+"Admin Name for your rocket.chat server"+constants.White)
	start.Flags().BoolVarP(&Virtual, "virtual", "v", true, constants.Blue+"Mounts your app directory to a companion container where all the app dependencies are present like node, npm & rc-apps cli, false uses the local environment for the dependencies."+constants.White)
	start.Flags().BoolVarP(&Deps, "deps", "d", true, constants.Blue+"Installs your app dependencies at the beginning, disable if you don't wanna run `npm i` on every run."+constants.White)

	start.Flags().StringVar(&composeFilePath, "composefilepath", "./", constants.Blue+"docker-compose file path, if you want any additional containers to start along with the environment"+constants.White)

	bindWithFlags()
}

func checkConfig(appDir string) {
	viper.GetViper().AddConfigPath(ConfigPath)
	viper.GetViper().SetConfigFile(".rc.yaml")
	err := viper.GetViper().ReadInConfig()

	if err != nil {
		fmt.Println(err.Error())
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println(constants.Yellow + "Config File Not Provided, generating one with default configuration for you in the given directory" + constants.Yellow)
			generateDefaultConfig(appDir)

		} else {
			fmt.Println(constants.Red + "Woops, Something went wrong :(" + constants.White)
		}
	}
}

func bindWithFlags() {
	viper.BindPFlags(start.Flags())
}

func generateDefaultConfig(appDir string) {
	viper.WriteConfigAs(".rc.yaml")
}
