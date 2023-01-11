package main

import (
	"RCTestSetup/Packages/Colors"
	"RCTestSetup/Packages/Figure"
	initiateadmin "RCTestSetup/Packages/InitiateAdmin"
	"fmt"
)

func InitiatePhase2(data map[string]interface{}) {
	fmt.Println(Colors.Blue() + "Phase 2 : Configuring Rocket.Chat App, installing admin\n" + Figure.Line())
	initiateadmin.Initiate(data)
}
