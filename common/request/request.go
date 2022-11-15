package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "http://127.0.0.1"

func ParseData(response *http.Response) map[string]interface{} {
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var mapData map[string]interface{}
	json.Unmarshal(all, &mapData)
	return mapData
}
