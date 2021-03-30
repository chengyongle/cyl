package dao

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)
type User struct{
	Name string
	Password string
}
func Register(c *gin.Context){
	_,err:=c.Request.Cookie("user_cookie")
	if err==nil{
		c.JSON(200,gin.H{
			"code":20000,
			"message":"您已登录",
		})
		return
	}
	name:=c.PostForm("username")
	password:=c.PostForm("password")
	fmt.Println(name,password)
	sqlstr:="select *from user where name =?;"
	err=DB.QueryRow(sqlstr,name).Scan()
	if err!=sql.ErrNoRows{
		//账号已存在
		c.JSON(200,gin.H{
			"code":20000,
			"message":"账号已存在，请登录",
		})
		return
	}
	stmt,err:=DB.Prepare("insert into user (name,password) value(?,?);")
	if err!=nil{
		fmt.Printf("mysql prepare failes:%v",err)
		c.String(200,"注册失败")
		return
	}
	defer stmt.Close()
	_,err=stmt.Exec(name,password)
	if err!=nil{
		fmt.Printf("insert failed:%v",err)
		c.String(200,"注册失败")
		fmt.Println("注册失败")
		return
	}
	c.String(200,"注册成功")
	return
}
func Login(c *gin.Context){
	_,err:=c.Request.Cookie("user_cookie")
	if err==nil{
		c.String(200,"你已登录")
		return
	}
	name:=c.PostForm("username")
	password:=c.PostForm("password")
	sqlstr:="select name,password from user where name=?;"
	var user User
	err=DB.QueryRow(sqlstr,name).Scan(&user.Name,&user.Password)
	if err==sql.ErrNoRows{
		c.Writer.Write([]byte("账号不存在，请先注册"))
		return
	}
	if user.Password!=password{
		//密码错误
		c.Writer.Write([]byte("密码错误"))
		return
	}
	//登陆成功
	c.SetCookie("user_cookie", user.Name, 1000, "/", "127.0.0.1", false, true)
	c.Writer.Write([]byte("登陆成功"))
	return
}