package usecase

import (
	"context"

	"github.com/sirupsen/logrus"
)

type usecase struct {
	logger *logrus.Logger
}

type UsecaseInterface interface {
	EncodQueryParams(ctx context.Context, params any) (encodedqp string, err error)
}

func NewUsecase(logger *logrus.Logger) (UsecaseInterface, error) {
	return &usecase{
		logger: logger,
	}, nil
}
