package  models

import (
	"awesomeProject/dao"
	"database/sql"
	"fmt"
)
type User struct{
	Name string
	Password string
}
//需要创建user数据表 包含name,password
func Register(name string,password string)int{
	sqlstr:="select *from user where name =?;"
	_,err:=dao.DB.Query(sqlstr,name)
	if err!=sql.ErrNoRows{
		//账号已存在
		return 0
	}
	stmt,err:=dao.DB.Prepare("insert into user (name,password) value(?,?);")
	if err!=nil{
		fmt.Printf("mysql prepare failes:%v",err)
		return 1
	}
	defer stmt.Close()
	_,err=stmt.Exec(name,password)
	if err!=nil{
		fmt.Printf("insert failed:%v",err)
		return 1
	}
	//注册成功
	return 2
}
func Login(name string,password string)int{
	sqlstr:="select *from user where name=?;"
	row,err:=dao.DB.Query(sqlstr,name)
	if err==sql.ErrNoRows{
		//账号不存在
		return 0
	}
	defer row.Close()
	var user User
	err=row.Scan(&user.Name,&user.Password)
	if err!=nil{
		fmt.Printf("scan failed:%v",err)
		return 1
	}
	if user.Password!=password{
		//密码错误
		return 2
	}
	//登陆成功
	return 3
}