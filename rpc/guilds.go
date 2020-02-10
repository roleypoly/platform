package platformrpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	dbGuild "github.com/roleypoly/db/ent/guild"
	"github.com/roleypoly/rpc/platform"
	"github.com/roleypoly/rpc/shared"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog"
)

func getSession(ctx context.Context) (*shared.RoleypolySession, error) {
	session, ok := ctx.Value(contextSession).(*shared.RoleypolySession)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Not logged in.")
	}

	return session, nil
}

func (p *PlatformService) EnumerateMyGuilds(ctx context.Context, req *empty.Empty) (*platform.GuildEnumeration, error) {
	session, err := getSession(ctx)
	if err != nil {
		return nil, err
	}

	query := &shared.IDQuery{
		MemberID: session.User.DiscordUser.ID,
	}

	guilds, err := p.discord.GetGuildsByMember(context.Background(), query)
	if err != nil {
		return nil, err
	}

	presentableGuilds := make([]*platform.PresentableGuild, len(guilds.Guilds))

	for idx, guild := range guilds.Guilds {
		pGuild, err := p.makePresentableGuild(guild, session.User.DiscordUser.ID)
		if err != nil {
			klog.Error("guild ", guild.ID, " failed to be presentable: ", err)
			continue
		}

		presentableGuilds[idx] = pGuild
	}

	return &platform.GuildEnumeration{
		Guilds: presentableGuilds,
	}, nil
}

func (p *PlatformService) makePresentableGuild(guild *shared.Guild, memberID string) (*platform.PresentableGuild, error) {
	memberQuery := &shared.IDQuery{
		MemberID: memberID,
		GuildID:  guild.ID,
	}

	member, err := p.discord.GetMember(context.Background(), memberQuery)
	if err != nil {
		klog.Errorf("Query of Member: %s in Guild: %s errored: %v", guild.ID, err)
		return nil, status.Error(codes.Internal, "Member request failed.")
	}

	roles, err := p.discord.GetGuildRoles(context.Background(), memberQuery)
	if err != nil {
		klog.Errorf("Query of roles in Guild: %s errored: %v", guild.ID, err)
		return nil, status.Error(codes.Internal, "Role request failed.")
	}

	// TODO: catch errors and ensure this data exists.
	guildData, err := p.db.Guild.Query().Where(dbGuild.Snowflake(guild.ID)).First(context.Background())
	if err != nil {
		klog.Errorf("Query of data in Guild: %s errored: %v", guild.ID, err)
		return nil, status.Error(codes.Internal, "Data request failed.")
	}

	pGuild := &platform.PresentableGuild{}

	pGuild.Guild = guild
	pGuild.ID = guild.ID
	pGuild.Member = member
	pGuild.Roles = roles
	pGuild.Data = guildDataToProto(guildData)

	return pGuild, nil
}

func (p *UnauthenticatedPlatformService) GetGuildSlug(ctx context.Context, req *shared.IDQuery) (*shared.Guild, error) {
	// TODO: safety
	return p.discord.GetGuild(context.Background(), req)
}

func (p *PlatformService) GetGuild(ctx context.Context, req *shared.IDQuery) (*platform.PresentableGuild, error) {
	session, err := getSession(ctx)
	if err != nil {
		return nil, err
	}

	memberID := session.User.DiscordUser.ID

	guild, err := p.discord.GetGuild(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return p.makePresentableGuild(guild, memberID)
}

func (p *PlatformService) UpdateMyRoles(ctx context.Context, req *platform.UpdateRoles) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented.")
}

func (p *PlatformService) UpdateGuildData(ctx context.Context, req *platform.GuildData) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented.")
}
