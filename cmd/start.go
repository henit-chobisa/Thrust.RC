package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	constants "thrust/Packages/Constants"
	models "thrust/Packages/Models"

	"thrust/Packages/Handlers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var start = &cobra.Command{
	Use:                   "start [flags] <App-Directory>",
	Short:                 "This initiates Rocket.Chat container and injects your app inside.",
	Long:                  "Initiates a pre-setup Rocket.Chat development environment for using Rocket.Chat apps using docker, installs your app inside and launches a browser window for you.\n" + constants.Red + "Docker must be running to use this." + constants.White,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		path, err := filepath.Abs(args[0])

		if err != nil {
			return err
		}

		checkConfig(args[0])

		appInfo, err := getAppInfo(path)

		if err != nil {
			fmt.Println(constants.Red + "`app.json` file not found in the given directory " + path + "\n Please consider rechecking and try again.")
			return nil
		}

		if err != nil {
			return err
		}

		err = Handlers.HandleDependencyCheck()
		if err != nil {
			return err
		}

		imagesToPull, err := Handlers.HandlePullingImages()

		if err != nil {
			return err
		}

		if len(imagesToPull) != 0 {
			err = Handlers.PullImages(imagesToPull)
			if err != nil {
				return err
			}

			verifyimagesToPull, err := Handlers.HandlePullingImages()

			if err != nil {
				return err
			}
			if len(verifyimagesToPull) != 0 {
				fmt.Println(constants.Red + "\nLooks like there is some issue in pulling the images, not all the images are pulled, Verification Failed\n" + constants.White)
				return nil
			}
		}

		containersToStart, startCompanion, companionID, err := Handlers.CheckRequiredContainers(appInfo)

		if err != nil {
			return err
		}

		_, startRocketChat := containersToStart[constants.RocketChatImage]
		_, startMongoDb := containersToStart[constants.MongoDBImage]

		if startMongoDb || startRocketChat {
			err = Handlers.StartContainersWithDefaultNetwork(containersToStart)
		}

		if err != nil {
			return err
		}

		Handlers.CreateAdminUser()

		fmt.Println(startCompanion)

		if startCompanion {
			err := Handlers.StartCompanionContainer(path, appInfo)
			if err != nil {
				return err
			}
		} else {
			Handlers.ShowLogs(companionID)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(start)
	initStartFlags()
}

func initStartFlags() {
	// start.Flags().StringP("config", "c", "./", constants.Blue+"Path of the configuration file for companion, if not found will use the default values."+constants.White)
	// start.Flags().BoolP("watcher.watcher", "w", true, constants.Blue+"Specify whether you want use hot-reloading"+constants.White)
	// start.Flags().StringP("admin.username", "a", "user0", constants.Blue+"Admin Username for your rocket.chat server"+constants.White)
	// start.Flags().StringP("admin.email", "e", "a@b.com", constants.Blue+"Admin Email for your rocket.chat server"+constants.White)
	// start.Flags().StringP("admin.password", "p", "123456", constants.Blue+"Admin Username for your rocket.chat server"+constants.White)
	// start.Flags().StringP("admin.name", "n", "user", constants.Blue+"Admin Name for your rocket.chat server"+constants.White)
	// start.Flags().BoolP("virtual", "v", true, constants.Blue+"Mounts your app directory to a companion container where all the app dependencies are present like node, npm & rc-apps cli, false uses the local environment for the dependencies."+constants.White)
	// start.Flags().BoolP("deps", "d", true, constants.Blue+"Installs your app dependencies at the beginning, disable if you don't wanna run `npm i` on every run."+constants.White)

	// start.Flags().String("composefilepath", "./", constants.Blue+"docker-compose file path, if you want any additional containers to start along with the environment"+constants.White)

	// TODO : Take Rocket.Chat version too
	// TODO : have a flag of new Rocket.Chat Server
	// TODO : Allow sharing of context in server, what you have to do is load you app into an ec2 instance and share that instance to the other user, there would be a new command something like thrust view <id of the instance> and we will open this container as a remote container to the workspace

	bindWithFlags()
}

func getAppInfo(path string) (appInfo *models.AppInfo, err error) {
	if _, err := os.Stat(path + "/app.json"); err == nil {
		appInfo, err = appInfo.New(path + "/app.json")
		return appInfo, err
	} else {
		return nil, err
	}
}

func checkConfig(appDir string) {
	viper.GetViper().AddConfigPath(viper.GetString("config"))
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
