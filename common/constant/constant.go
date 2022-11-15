package constant

import tele "gopkg.in/telebot.v3"

// 内联按钮
var (
	ReplyMenu       = &tele.ReplyMarkup{ResizeKeyboard: true}
	ReplyButtonHelp = ReplyMenu.Text("帮助")
	ReplyButtonUser = ReplyMenu.Text("个人中心")
)

// 回复按钮
var (
	InlineButtonMenu = &tele.ReplyMarkup{
		ResizeKeyboard:  false,
		Selective:       true,
		OneTimeKeyboard: true,
	}
	InlineButtonSign  = InlineButtonMenu.Data("签到", "sign", "456")
	InlineButtonCount = InlineButtonMenu.Data("统计", "count", "456")
	InlineButtonHelp  = InlineButtonMenu.Data("帮助", "help", "456")
	InlineButtonUSer  = InlineButtonMenu.Data("个人中心", "user", "456")
)
