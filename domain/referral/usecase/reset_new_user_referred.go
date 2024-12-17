package usecase

import (
	"context"

	"telegram-clicker-game-be/domain/referral/model"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (u *usecase) ResetNewUserReferred(ctx context.Context) (res []model.User, err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
	}).Info("Usecase: ResetNewUserReferred")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, u.logger)

	userInfo := ctx.Value("user_info").(*initdata.InitData).User

	err = u.referralRepo.ResetNewUserReferred(ctx, int(userInfo.ID))
	if err != nil {
		return
	}

	return
}
