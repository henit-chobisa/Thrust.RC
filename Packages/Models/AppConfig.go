package models

import (
	"encoding/json"
	"io/ioutil"
)

type AppsConfig struct {
	Url          string   `json:"url"`
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	IgnoredFiles []string `json:"ignoredFiles,omitempty"`
}

func arrHasValue(v []string, find string) bool {
	for _, val := range v {
		if val == find {
			return true
		}
	}
	return false
}

func (config *AppsConfig) New(path string) (data *AppsConfig, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &data)
	hasAppJson := arrHasValue(data.IgnoredFiles, "**/app.json")

	if !hasAppJson {
		data.IgnoredFiles = append(data.IgnoredFiles, "**/app.json")
		rawData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(path, rawData, 0644)
		if err != nil {
			return nil, err
		}
	}

	return
}
