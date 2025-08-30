package ai

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func QueryDeepSeek() {

	url := os.Getenv("AI_DeepSeek_URL")
	method := "POST"

	payload := strings.NewReader(`{
	  "messages": [
		{
		  "content": "You are a helpful assistant",
		  "role": "system"
		},
		{
		  "content": "Hi",
		  "role": "user"
		}
	  ],
	  "model": "deepseek-chat",
	  "frequency_penalty": 0,
	  "max_tokens": 2048,
	  "presence_penalty": 0,
	  "response_format": {
		"type": "text"
	  },
	  "stop": null,
	  "stream": false,
	  "stream_options": null,
	  "temperature": 1,
	  "top_p": 1,
	  "tools": null,
	  "tool_choice": "none",
	  "logprobs": false,
	  "top_logprobs": null
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer <TOKEN>")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Body关闭失败:", err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
