package models

// QWenRequestBody 是发送给 Qwen API 的请求体结构
type QWenRequestBody struct {
	Model    string        `json:"model"`
	Messages []QWenMessage `json:"messages"`
	Stream   bool          `json:"stream"`
}

// QWenMessage 是对话消息的结构
type QWenMessage struct {
	Role    string `json:"role"`    // 角色：system, user, assistant
	Content string `json:"content"` // 消息内容
}
