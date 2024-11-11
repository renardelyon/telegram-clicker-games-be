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

func (u *usecase) ValidateAndBindUserInfo(ctx context.Context, telData string) (data initdata.InitData, err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"data":       telData,
	}).Info("Usecase: ValidateAndBindUserInfo")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, u.logger)
	// TODO validate
	// err := initdata.Validate("query_id=AAGjwyFlAgAAAKPDIWUYz3sP&user=%7B%22id%22%3A5991678883%2C%22first_name%22%3A%22Putu%22%2C%22last_name%22%3A%22Naga%22%2C%22username%22%3A%22Singarajal%22%2C%22language_code%22%3A%22en%22%2C%22allows_write_to_pm%22%3Atrue%7D&auth_date=1729341198&hash=7670c35759f1a1611bb99ede28c4a121a488b0375be1e9110a83567570ca3e06", "627618978:amnnncjocxKJf", 24*time.Hour)

	data, err = initdata.Parse(telData)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	// if there are document with that id stop here
	_, err = u.authRepo.FindDocumentByTelegrarmId(ctx, data.User.ID)
	if err == nil {
		return
	}
	if err != mongo.ErrNoDocuments {
		errTrace = error_utils.HandleError(err)
		return
	}

	taskMasters, err := u.authRepo.GetAllTasks(ctx)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	upgradeMasters, err := u.authRepo.GetAllUpgrades(ctx)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	upgrades := make([]model.Upgrade, 0)
	tasks := make([]model.Task, 0)

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

	gameState := model.GameState{
		ClickCount:   0,
		Balance:      0,
		TotalBalance: 0,
		Energy:       0,
		MaxEnergy:    0,
		MiningRate:   0,
	}

	user := data.User
	newUser := model.Users{
		TelegramId: user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		UserName:   user.Username,
		CreatedAt:  time.Now(),
		LangCode:   user.LanguageCode,
		IsPremium:  user.IsPremium,
		Upgrades:   upgrades,
		Tasks:      tasks,
		GameStates: gameState,
		Referral: model.Referral{
			Referrals: []int{},
		},
	}

	// TODO upsert
	err = u.authRepo.InserUserData(ctx, &newUser)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	return

}
