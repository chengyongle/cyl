package dao
//初始化数据库
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type User struct {
	ID int `gorm:"auto_increment"`
	Username string  `gorm:"column:username"`
	Password string  `gorm:"column:password"`
}
var DB *gorm.DB
func SqlInit() *gorm.DB {
	dsn:="root:123456@tcp(127.0.0.1:3306)/database1?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",dsn)
	if err!=nil{
		fmt.Println("连接数据库失败")
		return nil
	}
	fmt.Println("连接数据库成功")
	DB=db
	DB.AutoMigrate(&User{})
	return DB
}
