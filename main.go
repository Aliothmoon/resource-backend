package main

import (
	"context"
	"fmt"
	"github.com/MirrorChyan/resource-backend/internal/cache"
	"github.com/MirrorChyan/resource-backend/internal/config"
	"github.com/MirrorChyan/resource-backend/internal/db"
	"github.com/MirrorChyan/resource-backend/internal/ent"
	"github.com/MirrorChyan/resource-backend/internal/logger"
	"github.com/MirrorChyan/resource-backend/internal/pkg/stg"
	"github.com/MirrorChyan/resource-backend/internal/wire"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	_ "github.com/MirrorChyan/resource-backend/internal/banner"
	_ "net/http/pprof"
)

const BodyLimit = 50 * 1024 * 1024

func main() {

	setUpConfigAndLog()

	mysql, err := db.NewDataSource()

	if err != nil {
		zap.L().Fatal("failed to connect to database",
			zap.Error(err),
		)
	}

	defer func(m *ent.Client) {
		if err := m.Close(); err != nil {
			zap.L().Fatal("failed to close database")
		}
	}(mysql)

	if err := mysql.Schema.Create(context.Background()); err != nil {
		zap.L().Fatal("failed creating schema resources",
			zap.Error(err),
		)
	}

	// deps
	var (
		redis   = db.NewRedis()
		redSync = db.NewRedSync(redis)
		storage = stg.New()
		group   = cache.NewVersionCacheGroup()
		app     = fiber.New(fiber.Config{
			BodyLimit:   BodyLimit,
			ProxyHeader: fiber.HeaderXForwardedFor,
		})
	)

	handlerSet := wire.NewHandlerSet(zap.L(), mysql, redis, redSync, storage, group)

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: zap.L(),
	}))

	initRoute(app, handlerSet)

	addr := fmt.Sprintf(":%d", config.CFG.Server.Port)

	if err := app.Listen(addr); err != nil {
		zap.L().Fatal("failed to start server",
			zap.Error(err),
		)
	}

}

func setUpConfigAndLog() {
	config.CFG = config.New()
	zap.ReplaceGlobals(logger.New())
}

func initRoute(app *fiber.App, handlerSet *wire.HandlerSet) {
	r := app.Group("/")

	handlerSet.ResourceHandler.Register(r)

	handlerSet.VersionHandler.Register(r)

	handlerSet.MetricsHandler.Register(r)
}
