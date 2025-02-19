package main

import (
	"bufio"
	"fmt"
	"os"

	"github.huangli.gotools/tools/http_helper"
	"github.huangli.gotools/tools/log_helper"
)

var Dbg string

func main() {
	log_helper.InitLog(Dbg == "true")
	log_helper.Debug("泥嚎！")
	log_helper.Info("窝布嚎！")
	log_helper.Warn("嚎起来了！")
	log_helper.Error("油布嚎了！")

	log_helper.Debug("%s", log_helper.HiGreenText("嘿嘿嘿"))
	log_helper.Info("%s", log_helper.HiRedText("哈哈哈"))
	log_helper.Warn("%s", log_helper.HiWhiteText("鹅鹅鹅"))
	log_helper.Error("%s", log_helper.HiYellowText("嘎嘎嘎"))

	// 接收输入
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Chat: ")
	messages := []http_helper.Message{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" || text == "quit" || text == "bye" || text == "q" {
			break
		}
		messages = append(messages, http_helper.Message{
			Role:    "user",
			Content: text,
		})
		if len(messages) > 10 {
			messages = messages[1:]
		}
		response_msg, ok := http_helper.Chat(messages)
		if ok {
			messages = append(messages, response_msg)
			if len(messages) > 10 {
				messages = messages[1:]
			}
		}
		fmt.Println("--------")
		fmt.Printf("Chat: ")
	}
}
