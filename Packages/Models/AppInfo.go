package models

import (
	"encoding/json"
	"io/ioutil"
)

type AppInfo struct {
	Id                 string
	Version            string
	RequiredApiVersion string
	Name               string
	Description        string
	Author             struct {
		Name     string
		Homepage string
	}
}

func (info *AppInfo) New(path string) (data *AppInfo, err error) {
	fileJson, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(fileJson, &data)
	return
}
