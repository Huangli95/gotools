package log_helper

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	is_debug     bool
	logger       *log.Logger
	elogger      *log.Logger
	GreenText    func(a ...interface{}) string
	HiGreenText  func(a ...interface{}) string
	WhiteText    func(a ...interface{}) string
	HiWhiteText  func(a ...interface{}) string
	YellowText   func(a ...interface{}) string
	HiYellowText func(a ...interface{}) string
	RedText      func(a ...interface{}) string
	HiRedText    func(a ...interface{}) string
)

func InitLog(debug_mod bool) {
	is_debug = debug_mod
	GreenText = color.New(color.FgGreen).SprintFunc()
	HiGreenText = color.New(color.FgHiGreen).SprintFunc()
	WhiteText = color.New(color.FgWhite).SprintFunc()
	HiWhiteText = color.New(color.FgHiWhite).SprintFunc()
	YellowText = color.New(color.FgYellow).SprintFunc()
	HiYellowText = color.New(color.FgHiYellow).SprintFunc()
	RedText = color.New(color.FgRed).SprintFunc()
	HiRedText = color.New(color.FgHiRed).SprintFunc()

	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	elogger = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}

func Debug(format string, args ...interface{}) {
	if is_debug {
		logger.Print("[D] ", GreenText(fmt.Sprintf(format, args...)), "\n")
	}
}

func Info(format string, args ...interface{}) {
	logger.Print("[I] ", WhiteText(fmt.Sprintf(format, args...)), "\n")
}

func Warn(format string, args ...interface{}) {
	logger.Print("[W] ", YellowText(fmt.Sprintf(format, args...)), "\n")
}

func Error(format string, args ...interface{}) {
	elogger.Print("[E] ", RedText(fmt.Sprintf(format, args...)), "\n")
}
