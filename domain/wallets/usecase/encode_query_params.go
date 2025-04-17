package usecase

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"telegram-clicker-game-be/pkg/error_utils"

	"github.com/sirupsen/logrus"
)

func (u *usecase) EncodQueryParams(ctx context.Context, params any) (encodedqp string, err error) {
	u.logger.WithFields(logrus.Fields{
		"request_id": ctx.Value("request_id"),
		"upgradeId":  fmt.Sprintf("%+v", params),
	}).Info("Usecase: EncodQueryParams")

	var errTrace error
	defer error_utils.HandleErrorLog(&errTrace, u.logger)

	// Convert struct to JSON
	jsonData, err := json.Marshal(params)
	if err != nil {
		errTrace = error_utils.HandleError(err)
		return
	}

	base64Encoded := base64.StdEncoding.EncodeToString(jsonData)

	encodedqp = strings.TrimRight(
		strings.ReplaceAll(strings.ReplaceAll(base64Encoded, "+", "-"), "/", "_"), "=",
	)

	return
}
