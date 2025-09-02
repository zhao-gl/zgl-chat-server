package models

type RequestContent struct {
	Type    string `json:"type"`
	Role    string `json:"role"`
	Content string `json:"content"`
	Model   string `json:"model"`
}

// Message 消息结构体
type Message struct {
	Id         string `json:"id"`
	SessionId  string `json:"sessionId"`
	Role       string `json:"role"`       // "user" 或 "ai"
	Content    string `json:"content"`    // 消息内容
	CreateTime string `json:"createTime"` // 时间戳
	UpdateTime string `json:"updateTime"`
}

type Session struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	CreateTime string    `json:"createTime"`
	UpdateTime string    `json:"updateTime"`
	MsgList    []Message `json:"msgList"`
}

// QWenCompletionChunk 原始响应结构体定义
type QWenCompletionChunk struct {
	ID      string `json:"id"`
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"delta"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason interface{} `json:"finish_reason"`
	} `json:"choices"`
	Object            string      `json:"object"`
	Usage             interface{} `json:"usage"`
	Created           int64       `json:"created"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
	Model             string      `json:"model"`
}

type SSEResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}
