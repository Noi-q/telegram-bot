package tool

import (
	"bufio"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"os"
)

var Data = make([]map[string]interface{}, 10000)

func OpenFile() (string, *os.File) {
	file, err := os.Open("static/result.json")
	if err != nil {
		log.Println("读取失败: ", err)
	}
	var fileString string
	readerFile := bufio.NewReader(file)
	for {
		readString, err := readerFile.ReadString('\n')
		if err == io.EOF {
			break
		}
		fileString = fileString + readString
	}
	return fileString, file
}

// 解析Telegram数据
func ParseJSON() {
	fileJSON, file := OpenFile()
	value := gjson.Get(fileJSON, "messages")
	var i = 0
	value.ForEach(func(k, v gjson.Result) bool {
		title := v.Get("text.0").String()
		link := v.Get("text.1.href").String()
		des := v.Get("text.3").String()
		//log.Println(title)
		//log.Println(link)
		//log.Println(des)
		if title == "" {
			title = "title Null"
		}
		if link == "" {
			title = "link Null"
		}
		if des == "" {
			des = "des Null"
		}
		Data[i] = map[string]interface{}{
			"title": title,
			"link":  link,
			"des":   des,
		}
		i += 1
		return true
	})
	defer file.Close()
}
