// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// GuildsColumns holds the columns for the "guilds" table.
	GuildsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "message", Type: field.TypeString, Size: 2147483647},
		{Name: "categories", Type: field.TypeJSON},
	}
	// GuildsTable holds the schema information for the "guilds" table.
	GuildsTable = &schema.Table{
		Name:        "guilds",
		Columns:     GuildsColumns,
		PrimaryKey:  []*schema.Column{GuildsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GuildsTable,
	}
)

func init() {
}
