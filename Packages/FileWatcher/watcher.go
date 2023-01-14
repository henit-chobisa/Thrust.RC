package watcher

import (
	constants "RCTestSetup/Packages/Constants"
	"RCTestSetup/Packages/InstallApp"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

func getAllDirectories(path string) []string {
	searchDir := path

	fileList := []string{}

	filepath.Walk(searchDir, func(filePath string, f os.FileInfo, err error) error {
		if f.Name() == ".git" || f.Name() == "node_modules" {
			return filepath.SkipDir
		}
		if f.IsDir() {
			fileList = append(fileList, "./"+filePath)
		}
		return nil
	})

	return fileList
}

func triggerReload(path string) {
	_ = InstallApp.Install(path, "http://localhost:3000", "user0", "123456", true)
}

func initTimer(timer *time.Timer, path string) {
	for range timer.C {
		triggerReload(path)
		timer.Stop()
	}
}

func Watch(path string, mode string) error {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	dirs := getAllDirectories(path)

	if mode == "nil" {
		mode = constants.AppDir_mode_deep
	}

	fmt.Println(constants.Purple + "🔫 Started hot reloading with below configuration\n" + "Watcher Mode    :    " + mode + "\n" + "Path            :    " + path + "\n\n")

	fmt.Println(constants.Purple + "Waiting for file changes on path " + path)

	var timer *time.Timer

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) || event.Has(fsnotify.Remove) || event.Has(fsnotify.Create) {
					if timer == nil {
						timer = time.NewTimer(10 * time.Second)
						go initTimer(timer, path)
					}
					timer.Stop()
					timer.Reset(10 * time.Second)
				}
			case err, ok := <-watcher.Errors:
				fmt.Println("error", err)
				if !ok {
					return
				}
			}
		}
	}()

	for _, path := range dirs {
		error := watcher.Add(path)
		if error != nil {
			fmt.Println("Error watching directory", path)
		}
	}

	<-make(chan struct{})

	return nil
}
