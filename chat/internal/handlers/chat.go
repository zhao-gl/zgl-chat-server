package handlers

import (
	"bufio"
	"chat/internal/models"
	"chat/internal/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
)

// HandleCreateSession 处理创建新会话
func HandleCreateSession(c *gin.Context) {
	session, err := services.CreateSession()
	if err != nil {
		fmt.Println(err)
		SendResponse(c, 500, nil, "创建失败")
		return
	}
	// 返回会话ID
	SendResponse(c, 200, session, "创建成功")
}

// HandleQuerySessionById 处理查询指定ID的会话
func HandleQuerySessionById(c *gin.Context) {
	id := c.Query("id")
	// sessionId 空校验
	if id == "" {
		SendResponse(c, 500, nil, "参数错误")
		return
	}
	session, err := services.QuerySessionById(id)
	if err != nil {
		SendResponse(c, 500, nil, "获取失败")
		return
	}
	SendResponse(c, 200, session, "获取成功")
}

// HandleAllSessions 查询全部历史会话
func HandleAllSessions(c *gin.Context) {
	history, err := services.QueryAllSessions()
	if err != nil {
		SendResponse(c, 500, nil, "获取失败")
		return
	}
	SendResponse(c, 200, history, "获取成功")
}

// HandleDelSession 删除会话
func HandleDelSession(c *gin.Context) {
	id := c.Query("id")
	res, err := services.DelSession(id)
	if err != nil {
		SendResponse(c, 500, nil, "删除失败")
		return
	}
	SendResponse(c, 200, res, "删除成功")
}

// HandleAsk 处理聊天请求
func HandleAsk(c *gin.Context) {
	var req models.RequestContent
	// 绑定 JSON 数据到结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req.Type, req.Role, req.Content, req.Model)

	// 设置 SSE 头部
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	bodyTxt := services.Ask(models.RequestContent{
		Type:    req.Type,
		Role:    req.Role,
		Content: req.Content,
		Model:   req.Model,
	})

	// 逐行读取并处理流数据
	reader := bufio.NewReader(strings.NewReader(string(bodyTxt)))
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			c.SSEvent("error", err.Error())
			return
		}

		// 处理 SSE 数据行
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "data:") {
			data := strings.TrimSpace(strings.TrimPrefix(line, "data:"))

			// 检查是否结束
			if data == "[DONE]" {
				c.SSEvent("end", "stream finished")
				break
			}

			// 发送消息内容
			if data != "" {
				//fmt.Println("发送数据:", data)
				response, err := convertChatResponse(data)
				if err != nil {
					return
				}
				c.SSEvent("message", response)
				c.Writer.Flush()
			}
		}
	}
}

// 转换函数
func convertChatResponse(data string) (models.SSEResponse, error) {
	var response models.QWenCompletionChunk
	if err := json.Unmarshal([]byte(data), &response); err != nil {
		errRes := models.SSEResponse{
			ID: "error",
		}
		return errRes, err
	}

	simpleResp := models.SSEResponse{
		ID: "111",
	}

	// 提取所有 choices 中的 content
	var contentBuilder strings.Builder
	for _, choice := range response.Choices {
		contentBuilder.WriteString(choice.Delta.Content)
	}
	simpleResp.Content = contentBuilder.String()
	fmt.Println("发送数据:", simpleResp)
	return simpleResp, nil
}
