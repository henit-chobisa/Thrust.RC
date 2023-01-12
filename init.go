package main

import (
	"AppsCompanion/Packages/ConfigReader"
	constants "AppsCompanion/Packages/Constants"
	watcher "AppsCompanion/Packages/FileWatcher"
	"fmt"
	"os/exec"
)

func main() {
	data := ConfigReader.ReadConfig("config.json")
	appDir := constants.AppDir_default

	if data["appDir"] != nil {
		appDir = fmt.Sprintf("%v", data["appDir"])
	}

	InitiatePhase1(data, appDir)
	InitiatePhase2(data)
	InitiatePhase3(data, appDir)

	exec.Command("gp", "preview", "http://localhost:3000", "--external").Output()

	if data["watcher"] == nil || data["watcher"] == true || data["watcher"] == "true" {
		watcher.Watch(appDir, fmt.Sprintf("%v", data["watcherMode"]))
	}
}
