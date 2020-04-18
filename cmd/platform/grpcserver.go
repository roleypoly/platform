package main

import (
	"context"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/roleypoly/gripkit"
	"go.uber.org/fx"
	"google.golang.org/grpc"
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
		gripkit.WithOptions(
			grpc.UnaryInterceptor(
				grpc_middleware.ChainUnaryServer(
					grpc_auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) { return ctx, nil }),
				),
			),
		),
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
