// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AcceptanceStatusColumns holds the columns for the "acceptance_status" table.
	AcceptanceStatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true, Comment: "删除时间"},
		{Name: "name", Type: field.TypeString, Nullable: true, Comment: "应用名称"},
		{Name: "status", Type: field.TypeString, Nullable: true, Comment: "应用状态"},
		{Name: "app_id", Type: field.TypeString, Nullable: true, Comment: "应用ID"},
		{Name: "app_key", Type: field.TypeString, Nullable: true, Comment: "应用密钥"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注"},
		{Name: "creator_id", Type: field.TypeUint32, Nullable: true, Comment: "创建者ID"},
	}
	// AcceptanceStatusTable holds the schema information for the "acceptance_status" table.
	AcceptanceStatusTable = &schema.Table{
		Name:       "acceptance_status",
		Columns:    AcceptanceStatusColumns,
		PrimaryKey: []*schema.Column{AcceptanceStatusColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "acceptancestatus_id",
				Unique:  false,
				Columns: []*schema.Column{AcceptanceStatusColumns[0]},
			},
		},
	}
	// RealtimeWarehousingColumns holds the columns for the "realtime_warehousing" table.
	RealtimeWarehousingColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true, Comment: "删除时间"},
		{Name: "name", Type: field.TypeString, Nullable: true, Comment: "应用名称"},
		{Name: "status", Type: field.TypeString, Nullable: true, Comment: "应用状态"},
		{Name: "app_id", Type: field.TypeString, Nullable: true, Comment: "应用ID"},
		{Name: "app_key", Type: field.TypeString, Nullable: true, Comment: "应用密钥"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注"},
		{Name: "creator_id", Type: field.TypeUint32, Nullable: true, Comment: "创建者ID"},
	}
	// RealtimeWarehousingTable holds the schema information for the "realtime_warehousing" table.
	RealtimeWarehousingTable = &schema.Table{
		Name:       "realtime_warehousing",
		Columns:    RealtimeWarehousingColumns,
		PrimaryKey: []*schema.Column{RealtimeWarehousingColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "realtimewarehousing_id",
				Unique:  false,
				Columns: []*schema.Column{RealtimeWarehousingColumns[0]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AcceptanceStatusTable,
		RealtimeWarehousingTable,
	}
)

func init() {
	AcceptanceStatusTable.Annotation = &entsql.Annotation{
		Table:     "acceptance_status",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	RealtimeWarehousingTable.Annotation = &entsql.Annotation{
		Table:     "realtime_warehousing",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
}
