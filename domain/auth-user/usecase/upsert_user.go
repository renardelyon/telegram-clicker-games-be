package usecase

import (
	"context"
	"telegram-clicker-game-be/constant"
	"telegram-clicker-game-be/domain/auth-user/model"
	"telegram-clicker-game-be/pkg/error_utils"
	"time"

	"github.com/sirupsen/logrus"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u *usecase) UpsertUser(ctx context.Context) (err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
	}).Info("Usecase: UpsertUser")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, u.logger)

	userInfo := ctx.Value("user_info").(*initdata.InitData).User

	// get user document
	res, err := u.authRepo.FindDocumentByTelegrarmId(ctx, userInfo.ID)
	if err != nil && err != mongo.ErrNoDocuments {
		return
	}

	upgrades := make([]model.Upgrade, 0)
	tasks := make([]model.Task, 0)

	if len(res.Upgrades) <= 0 && len(res.Tasks) <= 0 {
		taskMasters, err := u.authRepo.GetAllTasks(ctx)
		if err != nil {
			errTrace = error_utils.HandleError(err)
			return err
		}

		upgradeMasters, err := u.authRepo.GetAllUpgrades(ctx)
		if err != nil {
			errTrace = error_utils.HandleError(err)
			return err
		}

		for _, tsk := range taskMasters {
			newTsk := model.Task{
				TaskId:      tsk.Id,
				Status:      constant.INCOMPLETE,
				LastUpdated: time.Now(),
			}

			tasks = append(tasks, newTsk)
		}

		for _, up := range upgradeMasters {
			newUp := model.Upgrade{
				UpgradeId: up.Id,
				NextCost:  up.BaseCost,
				Level:     1,
			}

			upgrades = append(upgrades, newUp)
		}
	} else {
		tasks = res.Tasks
		upgrades = res.Upgrades
	}

	res.TelegramId = userInfo.ID
	res.FirstName = userInfo.FirstName
	res.LastName = userInfo.LastName
	res.UserName = userInfo.Username
	res.LangCode = userInfo.LanguageCode
	res.IsPremium = userInfo.IsPremium
	res.Upgrades = upgrades
	res.Tasks = tasks
	res.Referral.Referrals = func() []int {
		if len(res.Referral.Referrals) <= 0 {
			return []int{}
		}

		return res.Referral.Referrals
	}()
	res.CreatedAt = func() time.Time {
		if res.CreatedAt.IsZero() {
			return time.Now()
		}

		return res.CreatedAt
	}()

	err = u.authRepo.UpsertUserData(ctx, &res)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return
}
