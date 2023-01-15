package Utils

import (
	"RCTestSetup/enums"
	"github.com/spf13/viper"
)

func GetConfig(param enums.StartOption) interface{} {
	return viper.GetViper().Get(param.String())
}
