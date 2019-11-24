package main

import (
	"context"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/roleypoly/platform/ent"
	platformrpc "github.com/roleypoly/platform/rpc"
	"go.uber.org/fx"
)

func NewDatabaseClient(lc fx.Lifecycle) (*ent.Client, error) {
	client, err := ent.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return client, nil
}

func main() {
	app := fx.New(
		fx.Provide(
			NewDiscordClient,
			NewAuthClient,
			NewGripkit,
			platformrpc.NewPlatformService,
			NewDatabaseClient,
		),
		fx.Invoke(platformrpc.RegisterRPCHandlers),
	)

	app.Run()
	<-app.Done()
}
