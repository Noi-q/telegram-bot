package request

import (
	"fmt"
	"net/http"
	"net/url"
)

func Get(path string, param map[string]interface{}) {
	// param
	data := url.Values{}
	for key, value := range param {
		data.Set(key, value.(string))
	}
	uri, err := url.ParseRequestURI(URL + path)
	if err != nil {
		fmt.Printf("Parse url requestUrl failed, err: %v\n", err)
	}
	uri.RawQuery = data.Encode()
	fmt.Println(uri.String())
	resp, err := http.Get(uri.String())
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp)
}
