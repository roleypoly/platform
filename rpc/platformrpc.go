package platformrpc

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	ent "github.com/roleypoly/db/ent"
	"github.com/roleypoly/gripkit"
	authBackend "github.com/roleypoly/rpc/auth/backend"
	discord "github.com/roleypoly/rpc/discord"
	platform "github.com/roleypoly/rpc/platform"
)

func RegisterRPCHandlers(gk *gripkit.Gripkit, platformSvc *PlatformService, unauthSvc *UnauthenticatedPlatformService) {
	platform.RegisterPlatformServer(gk.Server, platformSvc)
	platform.RegisterPlatformServer(gk.Server, unauthSvc)
}

type PlatformService struct {
	platform.UnimplementedPlatformServer
	db      *ent.Client
	discord discord.DiscordClient
	auth    authBackend.AuthBackendClient
}

type UnauthenticatedPlatformService struct {
	platform.UnimplementedPlatformServer
	db      *ent.Client
	discord discord.DiscordClient
}

func NewPlatformService(db *ent.Client, discord discord.DiscordClient, auth authBackend.AuthBackendClient) *PlatformService {
	return &PlatformService{
		db:      db,
		discord: discord,
		auth:    auth,
	}
}

type contextSessionType struct{}

var contextSession contextSessionType = contextSessionType{}

func (p *PlatformService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return ctx, err
	}

	sessionQuery := &authBackend.SessionQuery{
		Token: token,
	}
	session, err := p.auth.GetSession(context.Background(), sessionQuery)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, contextSession, session), nil
}

func NewUnauthenticatedPlatformService(db *ent.Client, discord discord.DiscordClient) *UnauthenticatedPlatformService {
	return &UnauthenticatedPlatformService{
		db:      db,
		discord: discord,
	}
}
