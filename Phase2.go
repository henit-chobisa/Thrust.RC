package main

import (
	"AppsCompanion/Packages/Colors"
	"AppsCompanion/Packages/Figure"
	initiateadmin "AppsCompanion/Packages/InitiateAdmin"
	"fmt"
)

func InitiatePhase2(data map[string]interface{}) {
	fmt.Println(Colors.Blue() + "Phase 2 : Configuring Rocket.Chat App, installing admin\n" + Figure.Line())
	initiateadmin.Initiate(data)
}
