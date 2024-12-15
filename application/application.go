package application

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Context    context.Context
	DBClient   *mongo.Client
	DBDatabase *mongo.Database
	Logger     *logrus.Logger
	HttpClient *resty.Client
	// MigrationRunner *migration.Runner
	MigrationFlag string
	IsMigration   bool
	ServiceMode   string
	ServiceName   string
}
