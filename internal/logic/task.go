package logic

import (
	"context"
	"github.com/MirrorChyan/resource-backend/internal/config"
	"github.com/MirrorChyan/resource-backend/internal/model"
	"github.com/bytedance/sonic"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"time"
)

type logger struct {
	lg *zap.SugaredLogger
}

func (l logger) Debug(args ...any) {
	l.lg.Debug(append([]any{"asynq: "}, args...)...)
}

func (l logger) Info(args ...any) {
	l.lg.Info(append([]any{"asynq: "}, args...)...)
}

func (l logger) Warn(args ...any) {
	l.lg.Warn(append([]any{"asynq: "}, args...)...)
}

func (l logger) Error(args ...any) {
	l.lg.Error(append([]any{"asynq: "}, args...)...)
}

func (l logger) Fatal(args ...any) {
	l.lg.Fatal(append([]any{"asynq: "}, args...)...)
}

func InitAsynqServer(l *zap.Logger, v *VersionLogic) *asynq.Server {
	var (
		conf = config.GConfig
	)
	server := asynq.NewServer(asynq.RedisClientOpt{
		Addr: conf.Redis.Addr,
		DB:   conf.Redis.AsynqDB,
	}, asynq.Config{
		Logger:      logger{l.Sugar()},
		Concurrency: 100,
	})
	mux := asynq.NewServeMux()
	mux.HandleFunc(DiffTask, func(ctx context.Context, task *asynq.Task) error {
		var payload model.PatchTaskPayload
		if err := sonic.Unmarshal(task.Payload(), &payload); err != nil {
			return err
		}
		l.Sugar().Info("generate incremental update package task: ", string(task.Payload()))
		var (
			start = time.Now()

			target     = payload.TargetVersionId
			current    = payload.CurrentVersionId
			system     = payload.OS
			arch       = payload.Arch
			resourceId = payload.ResourceId
		)
		err := v.GenerateIncrementalPackage(ctx, resourceId, target, current, system, arch)
		if err != nil {
			l.Sugar().Error("generate incremental update package task failed: ", string(task.Payload()))
			return err
		}
		l.Info("generate incremental update package task success",
			zap.Int64("cost time", time.Since(start).Milliseconds()),
			zap.String("resource id", resourceId),
			zap.Int("current", current),
			zap.Int("target", target),
			zap.String("os", system),
			zap.String("arch", arch),
		)
		return nil
	})

	if err := server.Start(mux); err != nil {
		panic(err)
	}
	return server
}
