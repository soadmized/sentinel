package build

import (
	"context"
	"crypto/subtle"
	"fmt"
	"html/template"

	"github.com/hibiken/asynq"
	influx "github.com/influxdata/influxdb-client-go/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/soadmized/sentinel/internal/api"
	"github.com/soadmized/sentinel/internal/config"
	"github.com/soadmized/sentinel/internal/producer"
	"github.com/soadmized/sentinel/internal/queue/qclient"
	"github.com/soadmized/sentinel/internal/repo"
	"github.com/soadmized/sentinel/internal/service"
)

type Builder struct {
	conf config.Config
}

func New(conf config.Config) (*Builder, error) {
	b := Builder{conf: conf}

	return &b, nil
}

func (b *Builder) API(ctx context.Context) (*api.API, error) {
	srv, err := b.service(ctx)
	if err != nil {
		return nil, err
	}

	server := echo.New()
	t := &api.Template{
		Templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	server.Renderer = t
	server.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte(b.conf.AppUser)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(b.conf.AppPass)) == 1 {
			return true, nil
		}

		return false, nil
	}))

	a := api.API{
		Service: srv,
		Server:  server,
	}

	return &a, nil
}

func (b *Builder) service(ctx context.Context) (*service.Service, error) {
	r, err := b.Repo(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get repo")
	}

	queueClient := b.queueClient()

	srv := service.Service{
		Repo:        r,
		QueueClient: queueClient,
	}

	return &srv, nil
}

func (b *Builder) Repo(ctx context.Context) (*repo.Repo, error) {
	url := fmt.Sprintf("http://localhost:%d", b.conf.Influx.Port)
	client := influx.NewClient(url, b.conf.Influx.Token)

	if ok, err := client.Ping(ctx); !ok {
		return nil, fmt.Errorf("cant connect to influx: %w", err)
	}

	writer := client.WriteAPI(b.conf.Influx.Org, b.conf.Influx.Bucket)
	reader := client.QueryAPI(b.conf.Influx.Org)

	r, err := repo.New(writer, reader)
	if err != nil {
		return nil, errors.Wrap(err, "building repo is failed")
	}

	return &r, nil
}

func (b *Builder) queueClient() *qclient.Client {
	redisOpts := asynq.RedisClientOpt{
		Addr: b.conf.RedisHost,
	}

	asynqClient := asynq.NewClient(redisOpts)

	return qclient.New(asynqClient)
}

func (b *Builder) kafkaClient(topic string) (*producer.Client, error) {
	if b.conf.Kafka.Brokers == nil {
		return nil, nil
	}

	var topicOpt kgo.ConsumerOpt

	if topic != "" {
		topicOpt = kgo.ConsumeTopics(topic)
	}

	client, err := kgo.NewClient(kgo.SeedBrokers(b.conf.Kafka.Brokers...), kgo.ConsumerGroup("sentinel-group"), topicOpt)
	if err != nil {
		return nil, errors.Wrap(err, "get kafka client")
	}

	if client == nil {
		return nil, nil
	}

	return producer.New(client, b.conf.Kafka.DatasetTopic), nil
}
