package command

import (
	"bot/common/constant"
	"bot/common/request"
	"bot/tool"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"strconv"
)

// 开始指令
func Start(bot *tele.Bot) {
	bot.Handle("/start", func(context tele.Context) error {
		sID := strconv.Itoa(int(context.Message().Chat.ID)) // 用户id
		fName := context.Message().Chat.FirstName           // 用户昵称
		response := request.Post("/api/user/insert", "user_id="+sID+"&first_name="+fName)
		parseData := request.ParseData(response)
		fmt.Println(parseData)
		if parseData["code"] == float64(200) {
			return context.Send("欢迎使用本机器人\n你当前账号登录状态: 登录成功!", constant.ReplyMenu, constant.InlineButtonMenu)
		}
		defer response.Body.Close()
		return context.Send("登录失败!", constant.ReplyMenu, constant.InlineButtonMenu)
	})
}

// 索引指令
func Repertory(bot *tele.Bot) {
	bot.Handle("/index", func(context tele.Context) error {
		payloadContent := context.Message().Payload
		if payloadContent == "" {
			return context.Send("索引内容不能为空!")
		}
		tool.ParseJSON()
		return context.Send(payloadContent)
	})
}
