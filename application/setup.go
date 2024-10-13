package application

import (
	"context"
	"errors"
	"telegram-clicker-game-be/config"
	lib_mongo "telegram-clicker-game-be/pkg/db/mongo"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Setup(cfg *config.Config, c *cli.Context) (*Application, error) {
	app := new(Application)
	if c.Bool("migrate-up") && c.Bool("migrate-down") {
		return app, errors.New("unexpected migration command should be one of that")
	}
	if c.Bool("migrate-up") || c.Bool("migrate-down") {
		app.IsMigration = true
		if err := runInit(
			initLogger(),
		// initMigration(cfg),
		)(app); err != nil {
			return app, err
		}
		if !c.Bool("migrate-up") {
			app.MigrationFlag = "down"
		}
		if !c.Bool("migrate-down") {
			app.MigrationFlag = "up"
		}
		return app, nil
	}
	if err := runInit(
		initLogger(),
		initDatabase(cfg),
	)(app); err != nil {
		return app, err
	}
	return app, nil
}

func runInit(appFuncs ...func(*Application) error) func(*Application) error {
	return func(app *Application) error {
		app.Context = context.Background()
		for _, fn := range appFuncs {
			if e := fn(app); e != nil {
				return e
			}
		}
		return nil
	}
}

// func initMigration(cfg *config.Config) func(*Application) error {
// 	return func(app *Application) error {
// 		sqlDB, err := db.NewMysqlDB(cfg)
// 		if err != nil {
// 			return err
// 		}
// 		runner, err := migration.NewRunner(
// 			migration.FileDriver("./migration_files"),
// 			migration.MysqlDriver(sqlDB),
// 			app.Logger,
// 		)
// 		if err != nil {
// 			return err
// 		}
// 		app.MigrationRunner = runner
// 		return nil
// 	}
// }

func initLogger() func(*Application) error {
	return func(app *Application) error {
		log := logrus.New()
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
		app.Logger = log
		return nil
	}
}

func initDatabase(cfg *config.Config) func(*Application) error {
	return func(app *Application) error {
		db, err := lib_mongo.NewMongoInstance(cfg)
		if err != nil {
			return err
		}
		app.DbClient = db
		return nil
	}
}
