// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/MirrorChyan/resource-backend/internal/cache"
	"github.com/MirrorChyan/resource-backend/internal/ent"
	"github.com/MirrorChyan/resource-backend/internal/handler"
	"github.com/MirrorChyan/resource-backend/internal/logic"
	"github.com/MirrorChyan/resource-backend/internal/logic/dispense"
	"github.com/MirrorChyan/resource-backend/internal/repo"
	"github.com/MirrorChyan/resource-backend/internal/vercomp"
	"github.com/go-redsync/redsync/v4"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewHandlerSet(logger *zap.Logger, db *ent.Client, rdb *redis.Client, redsync2 *redsync.Redsync, cg *cache.VersionCacheGroup, verComparator *vercomp.VersionComparator) *HandlerSet {
	resource := repo.NewResource(db)
	resourceLogic := logic.NewResourceLogic(logger, resource)
	resourceHandler := handler.NewResourceHandler(logger, resourceLogic)
	repoRepo := repo.NewRepo(db)
	version := repo.NewVersion(db)
	storage := repo.NewStorage(db)
	latestVersion := repo.NewLatestVersion(db)
	latestVersionLogic := logic.NewLatestVersionLogic(logger, latestVersion, verComparator)
	storageLogic := logic.NewStorageLogic(logger, storage)
	distributeLogic := dispense.NewDistributeLogic(logger, rdb)
	versionLogic := logic.NewVersionLogic(logger, repoRepo, version, storage, latestVersionLogic, storageLogic, rdb, redsync2, cg, distributeLogic)
	versionHandler := handler.NewVersionHandler(logger, resourceLogic, versionLogic, verComparator)
	metricsHandler := handler.NewMetricsHandler()
	heathCheckHandler := handler.NewHeathCheckHandlerHandler()
	handlerSet := &HandlerSet{
		ResourceHandler:   resourceHandler,
		VersionHandler:    versionHandler,
		MetricsHandler:    metricsHandler,
		HeathCheckHandler: heathCheckHandler,
	}
	return handlerSet
}

// wire.go:

type HandlerSet struct {
	ResourceHandler   *handler.ResourceHandler
	VersionHandler    *handler.VersionHandler
	MetricsHandler    *handler.MetricsHandler
	HeathCheckHandler *handler.HeathCheckHandler
}

var handlerSetStruct = wire.Struct(new(HandlerSet), "*")
