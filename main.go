package main

import "github.huangli.gotools/tools/log_helper"

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
}
