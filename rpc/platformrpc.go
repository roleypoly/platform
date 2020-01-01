package platformrpc

import (
	ent "github.com/roleypoly/db/ent"
	"github.com/roleypoly/gripkit"
	discord "github.com/roleypoly/rpc/discord"
	platform "github.com/roleypoly/rpc/platform"
)

func RegisterRPCHandlers(gk *gripkit.Gripkit, platformSvc *PlatformService) {
	platform.RegisterPlatformServer(gk.Server, platformSvc)
}

type PlatformService struct {
	platform.UnimplementedPlatformServer
	db      *ent.Client
	discord *discord.DiscordClient
}

func NewPlatformService(db *ent.Client, discord *discord.DiscordClient) *PlatformService {
	return &PlatformService{
		db:      db,
		discord: discord,
	}
}
