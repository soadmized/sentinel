package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"

	"github.com/soadmized/sentinel/internal/build"
	"github.com/soadmized/sentinel/internal/config"
	"github.com/soadmized/sentinel/internal/queue"
)

const (
	httpServerCmd  = "http"
	queueServerCmd = "queue"
)

func Run(ctx context.Context, conf config.Config, args []string) error {
	if len(args) == 0 {
		return errors.Errorf("no command given. possible commands: \n%s\n%s", queueServerCmd, httpServerCmd)
	}

	cmd := args[0]

	err := processCommand(ctx, cmd, conf)
	if err != nil {
		return err
	}

	return nil
}

func processCommand(ctx context.Context, cmd string, conf config.Config) error {
	switch cmd {
	case httpServerCmd:
		return runHttp(ctx, conf)
	case queueServerCmd:
		return runQueue(ctx, conf)
	default:
		return errors.Errorf("wrong command given. possible commands: \n%s\n%s", httpServerCmd, queueServerCmd)
	}
}

func runHttp(ctx context.Context, conf config.Config) error {
	builder, err := build.New(conf)
	if err != nil {
		return fmt.Errorf("get builder: %w", err)
	}

	a, err := builder.API(ctx)
	if err != nil {
		return fmt.Errorf("get api from builder: %w", err)
	}

	a.Route()

	log.Fatal(a.Start(conf.AppPort))

	return nil
}

func runQueue(ctx context.Context, conf config.Config) error {
	b, err := build.New(conf)
	if err != nil {
		return fmt.Errorf("get builder: %w", err)
	}

	repo, err := b.Repo(ctx)
	if err != nil {
		return fmt.Errorf("get repo from builder: %w", err)
	}

	handler := queue.New(repo)

	log.Print("starting queue server")

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: conf.RedisHost},
		asynq.Config{
			Concurrency:      5,
			GroupAggregator:  asynq.GroupAggregatorFunc(queue.AggregateTasks),
			GroupMaxSize:     100,
			GroupMaxDelay:    time.Minute * 5,
			GroupGracePeriod: time.Minute * 2,
			Queues: map[string]int{
				queue.MainQueue:  7,
				queue.RetryQueue: 3,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.Handle(queue.TypeMultipleDatasetTask, &handler)

	if err := srv.Run(mux); err != nil {
		return err
	}

	return nil
}
