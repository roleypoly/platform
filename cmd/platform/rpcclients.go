package main

import (
	authBackend "github.com/roleypoly/rpc/auth/backend"
	"github.com/roleypoly/rpc/discord"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
)

func newRpcClient(dial string) *grpc.ClientConn {
	conn, err := grpc.Dial(dial, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}

func NewDiscordClient() *discord.DiscordClient {
	conn := newRpcClient(os.Getenv("DISCORD_SVC_DIAL"))
	d := discord.NewDiscordClient(conn)
	return &d
}

func NewAuthClient() *authBackend.AuthBackendClient {
	conn := newRpcClient(os.Getenv("AUTH_SVC_DIAL"))
	ac := authBackend.NewAuthBackendClient(conn)
	return &ac
}
