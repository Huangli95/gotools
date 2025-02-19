package http_helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var client http.Client

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	Model     string  `json:"model"`
	CreatedAt string  `json:"created_at"`
	Message   Message `json:"message"`
	Done      bool    `json:"done"`
}

func Get(url string, timeout time.Duration) (string, bool) {
	client.Timeout = timeout
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", false
	}
	response, err := client.Do(request)
	if err != nil {
		return "", false
	}
	defer response.Body.Close()
	// 读取响应
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", false
	}
	return string(body), true
}

func postStreamer(url string, body string) (Message, bool) {
	client.Timeout = time.Minute * 5
	request, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return Message{}, false
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return Message{}, false
	}
	buf := make([]byte, 1024)
	data := ""
	resposne_role := ""
	for {
		n, err := response.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading response body: ", err)
			return Message{}, false
		}
		resp := ChatResponse{}
		if err := json.Unmarshal(buf[:n], &resp); err != nil {
			fmt.Println("Error unmarshalling response body: ", err)
			return Message{}, false
		}
		if resposne_role == "" {
			resposne_role = resp.Message.Role
		}
		data += resp.Message.Content
		fmt.Print(resp.Message.Content)
	}
	fmt.Println()
	return Message{Role: resposne_role, Content: data}, true
}

func Chat(msgs []Message) (Message, bool) {
	request := ChatRequest{
		Model:    "deepseek-r1:14b",
		Messages: msgs,
	}
	data, _ := json.Marshal(request)
	return postStreamer("http://127.0.0.1:11434/api/chat", string(data))
}
