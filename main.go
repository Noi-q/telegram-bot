package main

import (
	"bot/common/constant"
	"bot/common/request"
	"bot/component/button/inline"
	"bot/component/button/reply"
	"bot/component/command"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"time"
)

func main() {
	token := GetToken()
	pref := tele.Settings{
		Token: token,
		//Token:  "XXX:xxxxxxxxxx",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalln(err)
		return
	}
	bot.Use(middleware.Logger())
	bot.Use(middleware.AutoRespond())

	// 使用回复键盘
	constant.InlineButtonMenu.Inline(
		constant.InlineButtonMenu.Row(constant.InlineButtonSign, constant.InlineButtonCount),
		constant.InlineButtonMenu.Row(constant.InlineButtonHelp, constant.InlineButtonUSer),
	)
	// 使用内联键盘
	constant.ReplyMenu.Reply(
		constant.ReplyMenu.Row(constant.ReplyButtonHelp),
		constant.ReplyMenu.Row(constant.ReplyButtonUser),
	)
	// command-开始指令
	command.Start(bot)
	command.Repertory(bot)
	// ---- inline ----
	// 签到按钮
	inline.Sign(bot)
	// ---- reply ----
	// 帮助按钮
	reply.Help(bot)
	// 个人中心按钮
	reply.User(bot)
	bot.Start()
}

func GetToken() string {
	response := request.Post("/api/auth/token", "")
	parseData := request.ParseData(response)
	if parseData["code"] == float64(200) {
		fmt.Println("连接成功")
	}
	stringToken := parseData["token"]
	return stringToken.(string)
}
