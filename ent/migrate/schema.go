// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UploadsColumns holds the columns for the "uploads" table.
	UploadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 70},
		{Name: "uid", Type: field.TypeUUID},
		{Name: "mime_type", Type: field.TypeString},
		{Name: "size", Type: field.TypeInt},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
	}
	// UploadsTable holds the schema information for the "uploads" table.
	UploadsTable = &schema.Table{
		Name:       "uploads",
		Columns:    UploadsColumns,
		PrimaryKey: []*schema.Column{UploadsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "uploads_users_uploads",
				Columns:    []*schema.Column{UploadsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 70},
		{Name: "channel_name", Type: field.TypeString, Unique: true, Size: 70},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "specialist", Type: field.TypeString, Nullable: true},
		{Name: "age", Type: field.TypeInt, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "language", Type: field.TypeString, Nullable: true},
		{Name: "country", Type: field.TypeString},
		{Name: "shor_bio", Type: field.TypeString, Size: 270},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UploadsTable,
		UsersTable,
	}
)

func init() {
	UploadsTable.ForeignKeys[0].RefTable = UsersTable
}
