package main

import (
	"context"
	"os"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/roleypoly/gripkit"
	"go.uber.org/fx"
)

func NewGripkit(lc fx.Lifecycle) *gripkit.Gripkit {
	gk := gripkit.Create(
		gripkit.WithHTTPOptions(gripkit.HTTPOptions{
			Addr:        os.Getenv("PLATFORM_SVC_PORT"),
			TLSCertPath: os.Getenv("TLS_CERT_PATH"),
			TLSKeyPath:  os.Getenv("TLS_KEY_PATH"),
		}),
		gripkit.WithGrpcWeb(
			grpcweb.WithOriginFunc(func(o string) bool { return true }),
		),
		gripkit.WithDebug(),
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go gk.Serve()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			gk.Server.Stop()
			return nil
		},
	})

	return gk
}
