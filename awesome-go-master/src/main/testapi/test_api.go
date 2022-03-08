package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://api-vehicle.smartisan.com/v1/ugc/digg/like/?debug-mode=vehicle-debug&device_id=69948601533&vehicle_type=1234&app_name=byte_autoservice_lite&app_version=1.0.0&iid=6844346744600905485&os_version=10&device_platform=android&model=MRX-W09&aid=3292"
	method := "POST"

	payload := strings.NewReader("category_type=toutiao&item_type=article&item_id=6858875127090807303")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Cookie", "odin_tt=9391e8409ffaa838177c19e42f89b52b1be937cb3db86b7b4b6229c680f158c4b615a9a7d06fa0c437fb5c43348a465a2219b5fe3af9d310f72b55dbbf804245; d_ticket=279c32ea76a10cef485b48b0bec33a16f29e1; sid_guard=d9ebd9fd487967e295152da948bf5d35%7C1596779656%7C5184000%7CTue%2C+06-Oct-2020+05%3A54%3A16+GMT; uid_tt=32f2fa8f04557c7c81f774cc4f2ce408; sid_tt=d9ebd9fd487967e295152da948bf5d35; sessionid=d9ebd9fd487967e295152da948bf5d35")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
