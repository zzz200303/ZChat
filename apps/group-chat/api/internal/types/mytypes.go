package types

type MessageJson struct {
	Content  string `json:"content"`
	From     string `json:"from"`
	To       string `json:"to"`
	Type     string `json:"type"` //群消息还是用户消息
	SendTime string `json:"send_time"`
}
