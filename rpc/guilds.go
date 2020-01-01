package platformrpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	roleypoly "github.com/roleypoly/rpc/shared"
)

func (p *PlatformService) GetMyGuilds(ctx context.Context, req *empty.Empty) (*roleypoly.GuildList, error) {
	session := ctx.Value("session").(*roleypoly.RoleypolySession)
	query := roleypoly.IDQuery{
		MemberID: session.Id,
	}

	p.discord
}
