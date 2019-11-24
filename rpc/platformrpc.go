package platformrpc

import (
	"github.com/roleypoly/gripkit"
	ent "github.com/roleypoly/platform/ent"
	platform "github.com/roleypoly/rpc/platform"
)

func RegisterRPCHandlers(gk *gripkit.Gripkit, platformSvc *PlatformService) {
	platform.RegisterPlatformServer(gk.Server, platformSvc)
}

type PlatformService struct {
	platform.UnimplementedPlatformServer
	db *ent.Client
}

func NewPlatformService(db *ent.Client) *PlatformService {
	return &PlatformService{
		db: db,
	}
}
