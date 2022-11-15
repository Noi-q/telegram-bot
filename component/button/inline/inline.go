package inline

import (
	"bot/common/constant"
	"bot/common/request"
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func Sign(bot *tele.Bot) {

	bot.Handle(&constant.InlineButtonSign, func(context tele.Context) error {
		fmt.Println("----", context.Message().Chat.ID)
		sign := fmt.Sprintf(
			"【签到状态】：%s\n"+
				"【用户】       ：%d",
			"成功",
			context.Message().Chat.ID,
		)
		return context.Send(sign, constant.InlineButtonMenu)
	})

	bot.Handle(&constant.InlineButtonUSer, func(context tele.Context) error {
		request.Get("http://baidu.com", map[string]interface{}{
			"admin": "123",
		})
		user := fmt.Sprintf("【用户】: %d\n【索引次数】: %d\n【邀请人数】: %d\n【会员】: %s",
			context.Message().Chat.ID,
			20,
			0,
			"未开通")
		return context.Send(user, constant.InlineButtonMenu)

	})
}
