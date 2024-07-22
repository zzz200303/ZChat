package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type MessageInfo struct {
	Content  string `json:"content"`
	From     int64  `json:"from"`
	To       int64  `json:"to"`
	SendTime string `json:"send_time"`
}

func TestJson(t *testing.T) {
	var message = MessageInfo{
		From:     1,
		To:       2,
		Content:  "hello",
		SendTime: "1721619782",
	}
	jsonData, err := json.Marshal(message)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("jsonData:" + string(jsonData))
	var messageJson MessageInfo
	//key := fmt.Sprintf("%s::%d", constants.OFFLINE_MESSAGE, message.To) // 生成Redis键
	err = json.Unmarshal(jsonData, &messageJson) // 将JSON字符串解析为Message对象
	fmt.Println(messageJson)
	if err != nil {
		t.Error(err)
	}
}
