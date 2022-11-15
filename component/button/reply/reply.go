package reply

import (
	"bot/common/constant"
	"bot/common/request"
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func Help(bot *tele.Bot) {
	bot.Handle(&constant.ReplyButtonHelp, func(c tele.Context) error {
		return c.Send("123")
	})
}

func User(bot *tele.Bot) {
	bot.Handle(&constant.ReplyButtonUser, func(c tele.Context) error {
		request.Get("http://baidu.com", map[string]interface{}{
			"admin": "123",
		})
		user := fmt.Sprintf("用户: %d\n索引次数: %d\n邀请人数: %d\n会员: %s",
			c.Message().Chat.ID,
			20,
			0,
			"未开通")
		return c.Send(user, constant.InlineButtonMenu)
	})
}
