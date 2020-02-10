package platformrpc

import (
	"github.com/roleypoly/db/ent"
	"github.com/roleypoly/db/ent/schema"
	"github.com/roleypoly/rpc/platform"
)

func categoryTypeToProto(catType string) platform.Category_CategoryType {
	switch catType {
	case "multi":
		return platform.Category_multi
	case "single":
		return platform.Category_single
	default:
		return platform.Category_multi
	}
}

func categoryTypeToDB(catType platform.Category_CategoryType) string {
	switch catType {
	case platform.Category_multi:
		return "multi"
	case platform.Category_single:
		return "single"
	default:
		return "multi"
	}
}

func categoryToProto(category schema.Category) *platform.Category {
	return &platform.Category{
		ID:       category.ID,
		Name:     category.Name,
		Roles:    category.Roles,
		Hidden:   category.Hidden,
		Position: int32(category.Position),
		Type:     categoryTypeToProto(category.Type),
	}
}

func categoryToDB(category *platform.Category) schema.Category {
	return schema.Category{
		ID:       category.ID,
		Name:     category.Name,
		Roles:    category.Roles,
		Hidden:   category.Hidden,
		Position: int(category.Position),
		Type:     categoryTypeToDB(category.Type),
	}
}

func categoriesToProto(categories []schema.Category) []*platform.Category {
	protos := make([]*platform.Category, len(categories))

	for idx, category := range categories {
		protos[idx] = categoryToProto(category)
	}

	return protos
}

func categoriesToDB(categories []*platform.Category) []schema.Category {
	entries := make([]schema.Category, len(categories))

	for idx, category := range categories {
		entries[idx] = categoryToDB(category)
	}

	return entries
}

func guildDataToProto(guildData *ent.Guild) *platform.GuildData {
	return &platform.GuildData{
		ID:           guildData.Snowflake,
		Message:      guildData.Message,
		Categories:   categoriesToProto(guildData.Categories),
		Entitlements: guildData.Entitlements,
	}
}

func guildDataToDB(guildData *platform.GuildData) *ent.Guild {
	return &ent.Guild{
		Snowflake:    guildData.ID,
		Message:      guildData.Message,
		Categories:   categoriesToDB(guildData.Categories),
		Entitlements: guildData.Entitlements,
	}
}
