// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/roleypoly/platform/ent/schema"
)

// Guild is the model entity for the Guild schema.
type Guild struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Categories holds the value of the "categories" field.
	Categories []schema.Category `json:"categories,omitempty"`
}

// FromRows scans the sql response data into Guild.
func (gu *Guild) FromRows(rows *sql.Rows) error {
	var scangu struct {
		ID         int
		Message    sql.NullString
		Categories []byte
	}
	// the order here should be the same as in the `guild.Columns`.
	if err := rows.Scan(
		&scangu.ID,
		&scangu.Message,
		&scangu.Categories,
	); err != nil {
		return err
	}
	gu.ID = strconv.Itoa(scangu.ID)
	gu.Message = scangu.Message.String
	if value := scangu.Categories; len(value) > 0 {
		if err := json.Unmarshal(value, &gu.Categories); err != nil {
			return fmt.Errorf("unmarshal field categories: %v", err)
		}
	}
	return nil
}

// Update returns a builder for updating this Guild.
// Note that, you need to call Guild.Unwrap() before calling this method, if this Guild
// was returned from a transaction, and the transaction was committed or rolled back.
func (gu *Guild) Update() *GuildUpdateOne {
	return (&GuildClient{gu.config}).UpdateOne(gu)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gu *Guild) Unwrap() *Guild {
	tx, ok := gu.config.driver.(*txDriver)
	if !ok {
		panic("ent: Guild is not a transactional entity")
	}
	gu.config.driver = tx.drv
	return gu
}

// String implements the fmt.Stringer.
func (gu *Guild) String() string {
	var builder strings.Builder
	builder.WriteString("Guild(")
	builder.WriteString(fmt.Sprintf("id=%v", gu.ID))
	builder.WriteString(", message=")
	builder.WriteString(gu.Message)
	builder.WriteString(", categories=")
	builder.WriteString(fmt.Sprintf("%v", gu.Categories))
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (gu *Guild) id() int {
	id, _ := strconv.Atoi(gu.ID)
	return id
}

// Guilds is a parsable slice of Guild.
type Guilds []*Guild

// FromRows scans the sql response data into Guilds.
func (gu *Guilds) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scangu := &Guild{}
		if err := scangu.FromRows(rows); err != nil {
			return err
		}
		*gu = append(*gu, scangu)
	}
	return nil
}

func (gu Guilds) config(cfg config) {
	for _i := range gu {
		gu[_i].config = cfg
	}
}