package usecase

import (
	"context"
	"telegram-clicker-game-be/pkg/error_utils"
	"time"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (u *usecase) ValidateAndBindUserInfo(ctx context.Context, telData string) (data initdata.InitData, err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"data":       telData,
	}).Info("Usecase: ValidateAndBindUserInfo")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, u.logger)

	err = initdata.Validate(telData, u.cfg.Telegram.BotToken, 10*time.Hour)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	data, err = initdata.Parse(telData)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}
