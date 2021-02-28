package cmd

import (
	"NanjoFan/Tour/internal/sql2struct"
	"github.com/spf13/cobra"
	"log"
)

var (
	username string		// 用户名
	password string		// 密码
	host string			// IP
	charset string		// 字符集
	dbType string		// 数据库类型
	dbName string		// 数据库名
	tableName string	// 表名
)

var sqlpCmd = &cobra.Command{
	Use:   "sql",
	Short: "SQL 转换和处理",
	Long:  "SQL 转换和处理",
	Run: func(cmd *cobra.Command, args []string) { },
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "SQL 转换",
	Long:  "SQL 转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbMode := sql2struct.NewDBMode(dbInfo)
		// 连接数据库
		err := dbMode.Connect()
		if err != nil {
			log.Fatalln("DBModel.Connect Err:", err)
		}

		columns, err := dbMode.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalln("dbMode.GetColumns Err:", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalln("template.Generate Err:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sqlpCmd)
	sqlpCmd.AddCommand(sql2structCmd)

	sql2structCmd.Flags().StringVarP(&username, "username", "","", "请输入数据库账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "","", "请输入数据库密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "","127.0.0.1:3306", "请输入数据库的 HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "","utf8mb4", "请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "","mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "","", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "","", "请输入表名称")
}


