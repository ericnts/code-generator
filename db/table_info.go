package db

import (
	"errors"
	"github.com/ericnts/code-generator/config"
	"github.com/ericnts/code-generator/constant"
)

type TableInfo struct {
	Name    string
	Comment string
}

type ColumnInfo struct {
	Name         string
	DefaultValue string
	Nullable     string
	Type         constant.ColumnType
	MaxLength    string
	ColumnKey    string
	Comment      string
}

func (p *ColumnInfo) GetType() constant.FieldType {
	switch p.Type {
	case constant.CTChar, constant.CTVarchar, constant.CTTinyblob, constant.CTTinytext, constant.CTBlob,
		constant.CTText, constant.CTMediumblob, constant.CTMediumtext, constant.CTLongblob, constant.CTLongtext:
		return constant.FTString
	case constant.CTTinyint, constant.CTSmallint, constant.CTMediumint, constant.CTInt, constant.CTInteger, constant.CTDecimal:
		return constant.FTInt
	case constant.CTBigint:
		return constant.FTInt64
	case constant.CTFloat, constant.CTDouble:
		return constant.FTFloat
	case constant.CTDate, constant.CTTime, constant.CTYear, constant.CTDatetime, constant.CTTimestamp:
		return constant.FTTime
	default:
		return constant.FTString
	}
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
	err = DB.Raw(sqlStr, config.DBDatabase).Scan(&result).Error
	return
}

func GetTableComment(tableName string) (result string, err error) {
	sqlStr := `SELECT
					table_comment 
				FROM
					information_schema.TABLES 
				WHERE
					table_schema = ? AND table_name = ?
				ORDER BY
					table_name`
	tx := DB.Raw(sqlStr, config.DBDatabase, tableName).Scan(&result)
	if tx.RowsAffected != 1 {
		return "", errors.New("表不存在")
	}
	err = tx.Error
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
					column_comment as comment 
				FROM
					information_schema.COLUMNS 
				WHERE
					table_name = ? 
					AND table_schema = ?
				ORDER BY
					ordinal_position`

	err = DB.Raw(sqlStr, tableName, config.DBDatabase).Scan(&result).Error
	return
}
