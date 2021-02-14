package dao
//初始化数据库
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB
func SqlInit() *sql.DB {
	dsn:="root:123456@tcp(127.0.0.1:3306)/database1"
	db,err:=sql.Open("mysql",dsn)
	if err!=nil{
		fmt.Println("连接数据库失败")
		return nil
	}
	err=db.Ping()
	if err!=nil{
		fmt.Println("连接数据库失败")
		return nil
	}
	DB=db
	return DB
}
