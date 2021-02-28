package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接的核心对象
type DBModel struct {
	DBEngine *sql.DB
	DBInfo *DBInfo
}

// 存储MySql链接的基本信息
type DBInfo struct {
	DBType string
	Host string
	UserName string
	Password string
	Charset string
}

// 需要的一些表中的字段
type TableColumn struct {
	ColumnName string	// 列的名称
	DataType  string	// 数据类型。 仅包含类型信息
	IsNullable string	// 是否允许为NULL
	ColumnKey string	// 是否被索引
	ColumnType string	// 数据类型，主要是精度长度或者有无符号
	ColumnComment string	// 注释信息
}

// TODO
var DBTypeToStructType = map[string]string{
	"int": "int32",
	"tinyint": "int8",
	"smallint":"int",
	"mediumint":"int64",
	"bit":"int",
	"bool":"bool",
	"enum":"string",
	"set":"string",
	"varchar(255)":"string",
}

func NewDBMode(info *DBInfo) *DBModel {
	return &DBModel{
		DBInfo: info,
	}
}

// 连接数据库
func (m *DBModel)Connect() error {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local",
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
		)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil{
		return err
	}
	return nil
}

// 获取列数据
func (m *DBModel)GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT " +
		"COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?"
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.DataType,
			&column.ColumnComment)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}
	return columns, nil
}
