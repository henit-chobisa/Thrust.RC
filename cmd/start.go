package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
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
	Args: func(cmd *cobra.Command, args []string) error {
		if viper.GetBool("appMode") {
			if len(args) != 1 {
				return errors.New("Sorry, as you are using Thrust in app mode, you must provide path of the app")
			}
		} else {
			if len(args) != 0 {
				return errors.New("Thrust is not going to process any app as you're not in app mode, there is no need to provide any argument")
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		isAppMode := viper.GetBool("appMode")

		var appInfo *models.AppInfo
		var path string

		if isAppMode {
			path, err = filepath.Abs(args[0])

			if err != nil {
				return err
			}

			checkConfig(args[0])

			appInfo, err = getAppInfo(path)

			if err != nil {
				return err
			}

			_, err = getRCConfig(path)
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			Handlers.Cleanup(appInfo)
			os.Exit(1)
		}()

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

		containersToStart, startCompanion, containerIDs, err := Handlers.CheckRequiredContainers(appInfo)

		if err != nil {
			return err
		}

		_, startRocketChat := containersToStart[constants.RocketChatImage]
		_, startMongoDb := containersToStart[constants.MongoDBImage]

		if startMongoDb || startRocketChat {
			_, err = Handlers.StartContainersWithDefaultNetwork(containersToStart)
		}

		if err != nil {
			return err
		}

		Handlers.CreateAdminUser()

		if startCompanion && isAppMode {
			err := Handlers.StartCompanionContainer(path, appInfo)
			if err != nil {
				return err
			}
		} else {
			if isAppMode {
				Handlers.ShowLogs(containerIDs[constants.CompanionImage])
			} else {
				fmt.Printf(constants.Blue + "\n🐳 Showing logs of Already Running Rocket.Chat Instance\n\n" + constants.White)
				Handlers.ShowLogs(containerIDs[constants.RocketChatImage])
			}
		}

		return nil
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("RunE has been completed and is terminating")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(start)
	initStartFlags()
}

func initStartFlags() {

	start.Flags().StringP("admin.username", "a", "user0", constants.Blue+"Admin Username for your rocket.chat server"+constants.White)
	start.Flags().StringP("admin.email", "e", "a@b.com", constants.Blue+"Admin Email for your rocket.chat server"+constants.White)
	start.Flags().StringP("admin.password", "p", "123456", constants.Blue+"Admin Username for your rocket.chat server"+constants.White)
	start.Flags().StringP("admin.name", "n", "user", constants.Blue+"Admin Name for your rocket.chat server"+constants.White)
	start.Flags().Bool("appMode", true, constants.Blue+"Using Thrust's Rocket.Chat Instance for other dependent softwares"+constants.White)

	// start.Flags().BoolP("virtual", "v", true, constants.Blue+"Mounts your app directory to a companion container where all the app dependencies are present like node, npm & rc-apps cli, false uses the local environment for the dependencies."+constants.White)

	// start.Flags().BoolP("deps", "d", true, constants.Blue+"Installs your app dependencies at the beginning, disable if you don't wanna run `npm i` on every run."+constants.White)

	// start.Flags().String("composefilepath", "./", constants.Blue+"docker-compose file path, if you want any additional containers to start along with the environment"+constants.White)

	// TODO : Take Rocket.Chat version too

	// TODO : have a flag of new Rocket.Chat Server

	// TODO : Have a configuration option that is also associated with the flags and turn off app mode, only has the rocket.chat mode and show you the rocket.chat logs

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

func getRCConfig(path string) (appConfig *models.AppsConfig, err error) {
	if _, err = os.Stat(path + "/.rcappsconfig"); err == nil {
		appConfig, err = appConfig.New(path + "/.rcappsconfig")
	}
	return
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
