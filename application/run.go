package application

import (
	"fmt"
	"telegram-clicker-game-be/config"
	"telegram-clicker-game-be/middleware"
	route "telegram-clicker-game-be/routes"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (app *Application) Run(cfg *config.Config) error {
	if app.IsMigration {
		// return runMigration(app)
	}
	return runApp(cfg, app)
}

// func runMigration(app *Application) error {
// 	switch app.MigrationFlag {
// 	case "up":
// 		return app.MigrationRunner.Up()
// 	case "down":
// 		return app.MigrationRunner.Rollback()
// 	}
// 	return app.MigrationRunner.Up()
// }

func runApp(cfg *config.Config, app *Application) error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(CORSMiddleware())
	r.Use(RequestIDMiddleware(app))
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())

	if err := middleware.SetupAuthMiddleware(app.Logger, app.DBDatabase, r); err != nil {
		return err
	}

	apiRoute := r.Group("/api")

	if err := route.SetupAuthRoute(app.Logger, app.DBDatabase, app.DBClient, r, apiRoute); err != nil {
		return err
	}

	if err := route.SetupGameplayRoute(app.Logger, app.DBDatabase, app.DBClient, r, apiRoute); err != nil {
		return err
	}

	if err := route.SetupLeaderboardRoute(app.Logger, app.DBDatabase, r, apiRoute); err != nil {
		return err
	}

	if err := route.SetupTasksRoute(app.Logger, app.DBDatabase, app.DBClient, r, apiRoute); err != nil {
		return err
	}

	if err := route.SetupReferralRoute(app.Logger, app.DBDatabase, app.DBClient, r, apiRoute); err != nil {
		return err
	}

	app.Logger.Info("Starting server " + cfg.Application.ServerPort)
	if err := r.Run(fmt.Sprintf(":%s", cfg.Application.ServerPort)); err != nil {
		return err
	}
	return nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Authorization, X-SKIP-AUTH, X-App-Name")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func RequestIDMiddleware(app *Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Set("request_id", requestID)

		// Process the request
		c.Next()

		app.Logger.WithFields(logrus.Fields{
			"request_id": requestID,
			"path":       c.Request.URL.Path,
		}).
			Info("Request processed")
	}
}
