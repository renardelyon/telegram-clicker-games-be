package usecase

import (
	"context"
	"slices"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (u *usecase) CheckMembershipTelegram(ctx context.Context) (result bool, err error) {
	// TODO Caching
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
	}).Info("Usecase: CheckMembershipTelegram")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, u.logger)

	userInfo := ctx.Value("user_info").(*initdata.InitData).User

	res, err := u.authRepo.CheckMembershipTelegram(ctx, int(userInfo.ID))
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	memberStatus := []string{"creator", "administrator", "member"}
	result = slices.Contains(memberStatus, res.Result.Status)

	return
}
