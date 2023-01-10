package ConfigReader

import (
	"encoding/json"
	"io/ioutil"
)

func ReadConfig(path string) map[string]interface{} {
	config, err := ioutil.ReadFile(path)
	var response map[string]interface{}
	if err == nil {
		json.Unmarshal(config, &response)
	}
	return response
}
