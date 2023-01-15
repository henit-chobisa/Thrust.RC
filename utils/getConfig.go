package utils

import (
	"RCTestSetup/enums"

	"github.com/spf13/viper"
)

func getConfig(param enums.StartOption) interface{} {
	return viper.GetViper().Get(param.String())
}
