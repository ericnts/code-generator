package main

import "github.com/ericnts/code-generator/common/orm"

type TableInfo struct {
	Name    string
	Comment string
}

type ColumnInfo struct {
	Name         string
	DefaultValue string
	Nullable     string
	Type         string
	MaxLength    string
	ColumnKey    string
	Comment      string
}

func FindTable(schemaName string) (result []TableInfo, err error) {
	sqlStr := `SELECT
					table_name as name,
					table_comment as comment
				FROM
					information_schema.TABLES 
				WHERE
					table_schema = ? 
				ORDER BY
					table_name`
	err = orm.DB.Raw(sqlStr, schemaName).Scan(&result).Error
	return
}

func FindColumn(tableName string) (result []ColumnInfo, err error) {
	sqlStr := `SELECT
					column_name as name,
					column_default as default_value,
					is_nullable as nullable,
					data_type as type,
					character_maximum_length as max_length,
					column_key as column_key,
					column_comment as common 
				FROM
					information_schema.COLUMNS 
				WHERE
					table_name = ? 
				ORDER BY
					ordinal_position`

	err = orm.DB.Raw(sqlStr, tableName).Scan(&result).Error
	return
}
