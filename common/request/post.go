package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Post(path string, param string) *http.Response {
	contentType := "application/x-www-form-urlencoded"
	uri := URL + path
	resp, err := http.Post(uri, contentType, strings.NewReader(param))
	if err != nil {
		fmt.Printf("post failed, err: %v\n", err)
		return nil
	}
	return resp
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	marshal, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal.Err: %s", err)
	}
	w.Write(marshal)
}
