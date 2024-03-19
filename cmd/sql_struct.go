package cmd

import (
	"clidev/pkg/sql"
	"log"

	"github.com/spf13/cobra"
)

var (
	dbType    string
	host      string
	userName  string
	password  string
	charset   string
	dbName    string
	tableName string
)

var sqlStructRunFunc = func(cmd *cobra.Command, args []string) {
	dbInfo := &sql.DBInfo{
		DBType:   dbType,
		Host:     host,
		UserName: userName,
		Password: password,
		Charset:  charset,
	}
	dbModel := sql.NewDBModel(dbInfo)
	err := dbModel.Connect()
	if err != nil {
		log.Fatalf("dbModel.Connect err: %v", err)
	}
	columns, err := dbModel.GetColumns(dbName, tableName)
	if err != nil {
		log.Fatalf("dbModel.GetColumns err: %v", err)
	}
	template := sql.NewStructTemplate()
	templateColumns := template.AssemblyColumns(columns)
	err = template.Generate(tableName, templateColumns)
	if err != nil {
		log.Fatalf("template.Gennerate err: %v", err)
	}
}

var sqlStructCmd = &cobra.Command{
	Use:   "struct",
	Short: "struct tool",
	Long:  "This is a sql struct command",
	Run:   sqlStructRunFunc,
}

func init() {
	sqlStructCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例的类型")
	sqlStructCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST")
	sqlStructCmd.Flags().StringVarP(&userName, "user", "", "", "请输入数据库的账号")
	sqlStructCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库账号的密码")
	sqlStructCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sqlStructCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库的名称")
	sqlStructCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}
