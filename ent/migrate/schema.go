// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// ClubsColumns holds the columns for the "clubs" table.
	ClubsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// ClubsTable holds the schema information for the "clubs" table.
	ClubsTable = &schema.Table{
		Name:        "clubs",
		Columns:     ClubsColumns,
		PrimaryKey:  []*schema.Column{ClubsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DevicesColumns holds the columns for the "devices" table.
	DevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"jukebox", "requester"}},
		{Name: "club_id", Type: field.TypeInt, Nullable: true},
	}
	// DevicesTable holds the schema information for the "devices" table.
	DevicesTable = &schema.Table{
		Name:       "devices",
		Columns:    DevicesColumns,
		PrimaryKey: []*schema.Column{DevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "devices_clubs_device",
				Columns: []*schema.Column{DevicesColumns[2]},

				RefColumns: []*schema.Column{ClubsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MusicsColumns holds the columns for the "musics" table.
	MusicsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "club_id", Type: field.TypeInt, Nullable: true},
	}
	// MusicsTable holds the schema information for the "musics" table.
	MusicsTable = &schema.Table{
		Name:       "musics",
		Columns:    MusicsColumns,
		PrimaryKey: []*schema.Column{MusicsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "musics_clubs_music",
				Columns: []*schema.Column{MusicsColumns[1]},

				RefColumns: []*schema.Column{ClubsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DeviceUserColumns holds the columns for the "device_user" table.
	DeviceUserColumns = []*schema.Column{
		{Name: "device_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// DeviceUserTable holds the schema information for the "device_user" table.
	DeviceUserTable = &schema.Table{
		Name:       "device_user",
		Columns:    DeviceUserColumns,
		PrimaryKey: []*schema.Column{DeviceUserColumns[0], DeviceUserColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "device_user_device_id",
				Columns: []*schema.Column{DeviceUserColumns[0]},

				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "device_user_user_id",
				Columns: []*schema.Column{DeviceUserColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClubsTable,
		DevicesTable,
		MusicsTable,
		UsersTable,
		DeviceUserTable,
	}
)

func init() {
	DevicesTable.ForeignKeys[0].RefTable = ClubsTable
	MusicsTable.ForeignKeys[0].RefTable = ClubsTable
	DeviceUserTable.ForeignKeys[0].RefTable = DevicesTable
	DeviceUserTable.ForeignKeys[1].RefTable = UsersTable
}
