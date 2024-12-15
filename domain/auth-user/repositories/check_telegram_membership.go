package repo

import (
	"context"
	"strconv"
	"strings"
	"telegram-clicker-game-be/domain/auth-user/response"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
)

func (r *repo) CheckMembershipTelegram(ctx context.Context, telegramId int) (res response.TelegramMembershipResponse, err error) {
	r.logger.WithContext(ctx).
		WithFields(logrus.Fields{
			"telegram_id": telegramId,
			"request_id":  ctx.Value("request_id"),
		}).
		Info("Repo: CheckMembershipTelegram")

	var errorTrace error
	defer error_utils.HandleErrorLog(&errorTrace, r.logger)
	cfg := r.cfg.Telegram

	formattedChatId := cfg.ChannelUsername
	if !strings.HasPrefix(cfg.ChannelUsername, "@") && !strings.HasPrefix(cfg.ChannelUsername, "-100") {
		formattedChatId = "@" + cfg.ChannelUsername
	}

	telegramUrl := "https://api.telegram.org/bot{botToken}/getChatMember"

	var errRes map[string]interface{}
	resp, err := r.httpClient.R().SetPathParams(map[string]string{
		"botToken": cfg.BotToken,
	}).SetQueryParams(map[string]string{
		"chat_id": formattedChatId,
		"user_id": strconv.Itoa(telegramId),
	}).
		SetResult(&res).
		SetError(&errRes).
		Get(telegramUrl)

	r.logger.WithContext(ctx).
		WithFields(logrus.Fields{
			"body": resp.String(),
		}).
		Info("http func has been executed")
	if err != nil {
		r.logger.WithContext(ctx).
			WithFields(logrus.Fields{
				"status_code": resp.StatusCode(),
				"status":      resp.Status(),
				"errRes":      errRes,
			}).
			Error("http check telegram membership error")
		errorTrace = error_utils.HandleError(err)
		return
	}

	return
}
