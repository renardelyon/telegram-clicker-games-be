package application

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Context    context.Context
	DBClient   *mongo.Client
	DBDatabase *mongo.Database
	Logger     *logrus.Logger
	// MigrationRunner *migration.Runner
	MigrationFlag string
	IsMigration   bool
	ServiceMode   string
	ServiceName   string
}
