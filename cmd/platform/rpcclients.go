package main

import (
	"context"
	"log"
	"os"

	authBackend "github.com/roleypoly/rpc/auth/backend"
	"github.com/roleypoly/rpc/discord"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type sharedSecretCredential struct {
	Authorization string
}

func (s sharedSecretCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"Authorization": s.Authorization}, nil
}

func (sharedSecretCredential) RequireTransportSecurity() bool {
	return false
}

func newRpcClient(dial string, authorization string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		dial,
		grpc.WithTransportCredentials(credentials.NewTLS(nil)),
		grpc.WithPerRPCCredentials(sharedSecretCredential{Authorization: authorization}))
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}

func NewDiscordClient() discord.DiscordClient {
	conn := newRpcClient(os.Getenv("DISCORD_SVC_DIAL"), "Shared "+os.Getenv("SHARED_SECRET"))
	return discord.NewDiscordClient(conn)

}

func NewAuthClient() authBackend.AuthBackendClient {
	conn := newRpcClient(os.Getenv("AUTH_SVC_DIAL"), "Shared "+os.Getenv("SHARED_SECRET"))
	return authBackend.NewAuthBackendClient(conn)
}
