package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var dbname = flag.String("db", "cms", "the database name")
var tblname = flag.String("tbl", "cms_menu", "the table name to export")
var savepath = flag.String("path", "./", "the path to save file")

/**
生成数据库实体字段映射
*/
func main() {
	flag.Parse()
	fmt.Println("table name -->", *tblname)

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "root", "127.0.0.1:3306", "information_schema")

	db := sqlx.MustConnect("mysql", dns)

	var fs []FieldInfo
	err := db.Select(&fs, "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_COMMENT, IS_NULLABLE FROM COLUMNS WHERE TABLE_NAME=? and table_schema=?", *tblname, *dbname)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if len(fs) > 0 {
		var buffer bytes.Buffer
		buffer.WriteString("package models\n")
		buffer.WriteString("type " + fmtFieldDefine(*tblname) + " struct {\n")
		for _, v := range fs {
			buffer.WriteString("" + fmtFieldDefine(v.ColName) + " ")
			switch v.DataType {
			case "int", "tinyint", "smallint":
				if v.IsNullable == "YES" {
					buffer.WriteString("sql.NullInt64 ")
				} else {
					buffer.WriteString("int ")
				}
			case "bigint":
				if v.IsNullable == "YES" {
					buffer.WriteString("sql.NullInt64 ")
				} else {
					buffer.WriteString("int64 ")
				}
			case "char", "varchar", "longtext", "text", "tinytext":
				if v.IsNullable == "YES" {
					buffer.WriteString("sql.NullString ")
				} else {
					buffer.WriteString("string ")
				}
			case "date", "datetime", "timestamp":
				buffer.WriteString("time.Time ")
			case "double", "float":
				if v.IsNullable == "YES" {
					buffer.WriteString("sql.NullFloat64 ")
				} else {
					buffer.WriteString("float64 ")
				}
			default:
				// 其他类型当成string处理
				if v.IsNullable == "YES" {
					buffer.WriteString("sql.NullString ")
				} else {
					buffer.WriteString("string ")
				}
			}

			buffer.WriteString(fmt.Sprintf("`db:\"%s\" json:\"%s form:\"%s\"`\n", v.ColName, v.ColName, v.ColName))

		}
		buffer.WriteString(`}`)

		fmt.Println(buffer.String())

		filename := *savepath + "\\" + *tblname + ".go"
		f, _ := os.Create(filename)
		f.Write([]byte(buffer.String()))
		f.Close()

		cmd := exec.Command("goimports", "-w", filename)
		cmd.Run()
	} else {
		fmt.Println("查询不到数据")
	}
}

type FieldInfo struct {
	ColName    string `db:"COLUMN_NAME"`
	DataType   string `db:"DATA_TYPE"`
	ColComment string `db:"COLUMN_COMMENT"`
	IsNullable string `db:"IS_NULLABLE"`
}

func fmtFieldDefine(src string) string {
	temp := strings.Split(src, "_") // 有下划线的，需要拆分
	var str string
	for i := 0; i < len(temp); i++ {
		b := []rune(temp[i])
		for j := 0; j < len(b); j++ {
			if j == 0 {
				// 首字母大写转换
				b[j] -= 32
				str += string(b[j])
			} else {
				str += string(b[j])
			}
		}
	}

	return str
}
