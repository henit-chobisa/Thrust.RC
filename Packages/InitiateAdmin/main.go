package initiateadmin

import (
	"RCTestSetup/Packages/Colors"
	"RCTestSetup/Packages/Figure"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func setUser(user0 string, email string, pass string, name string) (bool, bool) {
	url := "http://localhost:3000/api/v1/users.register"
	jsonData := map[string]string{`username`: user0, `email`: email, `pass`: pass, `name`: name}
	data, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println(err)
		return false, false
	}
	resp, error := http.Post(url, "application/json", bytes.NewBuffer(data))
	if error != nil {
		return false, false
	}
	data, _ = ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(data, &response)
	if response["error"] == "Username is already in use" {
		return false, true
	}
	boolValue, err := strconv.ParseBool(fmt.Sprintf("%v", response["success"]))
	if err != nil {
		return false, false
	}
	return boolValue, false
}

func Initiate(data map[string]interface{}) {
	spinner := Figure.Spinner(" Creating Admin User, required for App Installation", Colors.Green(), "")
	spinner.Start()
	user := data["admin"].(map[string]interface{})
	iterations := 0
	status := false
	breakLoop := false
	time.Sleep(5 * time.Second)
	for iterations < 20 {
		status, breakLoop = setUser(fmt.Sprintf("%v", user["username"]), fmt.Sprintf("%v", user["email"]), fmt.Sprintf("%v", user["pass"]), fmt.Sprintf("%v", user["name"]))

		if status || breakLoop {
			break
		}
		time.Sleep(20 * time.Second)
		iterations++
	}
	spinner.Stop()
	if breakLoop {
		fmt.Println("\n" + Colors.Green() + "⭕ Admin User Already Present, Gracefully Aborting Operation ...\n")
		return
	}
	fmt.Println("\n" + Colors.Green() + "✅ Successfully created admin user for Rocket.Chat\n")

}
