package main
import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)
var db *sql.DB
var err error
func main(){
	e:=gin.Default()
	e.Use(status())
	initdb()
	//登录
	e.POST("/login",login)
	//注册
	e.POST("/register",register)
	//修改密码
	e.POST("/updatepw",updatepw)
	//设置签名
	e.POST("/setsign",setsign)
	e.POST("/exit",exit)
	e.Run()
}
type Person struct{
	Id int `form:"id"`
	Name string `form:"name"`
	Password string `form:"password"`
	Sign string `form:"sign"`
}
//初始化数据库
func initdb()(err error){
	dsn:="root:123456@tcp(127.0.0.1)/database1"
	db,err=sql.Open("mysql",dsn)
	if err!=nil{
		fmt.Println("连接数据库失败")
		return err
	}
	err=db.Ping()
	if err!=nil{
		fmt.Println("连接数据库失败")
		return err
	} else{
		fmt.Println("连接数据库成功")
	}
	return nil
}
//登录
func login(c *gin.Context){
	_,err=c.Request.Cookie("user_cookie")
	if err==nil{
		c.Writer.Write([]byte("您已登录"))
		return
	}
	var person Person
	c.ShouldBind(&person)
	var mm string
	sqlstr:="select password from user where name=?;"
	err:=db.QueryRow(sqlstr,person.Name).Scan(&mm)
	if err==sql.ErrNoRows{
		c.Writer.Write([]byte("该用户未注册"))
		return
	}
	if mm==person.Password{
		fmt.Println(person.Name)
		c.SetCookie("user_cookie",person.Name,1000,
			"/", "localhost", false, true)
		c.Writer.Write([]byte("登陆成功！"))
	}else{
		c.Writer.Write([]byte("密码错误"))
	}

}
//注册
func register(c *gin.Context){
	var person Person
	c.ShouldBind(&person)
	sqlstr:="select *from user where name =?;"
	err=db.QueryRow(sqlstr,person.Name).Scan()
	if err!=sql.ErrNoRows{
		c.Writer.Write([]byte("账户已存在"))
		return
	}
	db.Exec("insert into user(name,password) values(?,?);",person.Name,person.Password)
	c.Writer.Write([]byte("注册成功"))

}
//改密码
func updatepw(c *gin.Context){
	cookie,err:=c.Request.Cookie("user_cookie")
	if err!=nil{
		c.Writer.Write([]byte("您未登录，请先登录"))
		return
	}
	uppassword:=c.PostForm("uppassword")
	sqlstr:="update user set password=? where name =?;"
	_,err=db.Exec(sqlstr,uppassword,cookie.Value)
	if err!=nil{
		c.Writer.Write([]byte("密码修改失败"))
		return
	}
	c.SetCookie("user_cookie","",-1,
		"/", "localhost", false, true)
	c.Writer.Write([]byte("密码修改成功"))
}
//设置个性签名
func setsign(c *gin.Context){
	cookie,err:=c.Request.Cookie("user_cookie")
	if err!=nil{
		c.Writer.Write([]byte("您未登录，请先登录"))
		return
	}
	newsign:=c.PostForm("sign")
	sqlstr:="update user set sign=? where name=?;"
	_,err=db.Exec(sqlstr,newsign,cookie.Value)
	if err!=nil{
		c.Writer.Write([]byte("签名更新失败"))
	}
	c.Writer.Write([]byte("签名更新成功！下次登陆生效"))
}
//退出登录
func exit(c *gin.Context){
	c.SetCookie("user_cookie","",-1,
		"/", "localhost", false, true)
	c.Writer.Write([]byte("已退出登录"))
}
//中间件验证是否为登录状态
func status()gin.HandlerFunc{
	return func(c *gin.Context){
		c.Next()
		cookie,_:=c.Request.Cookie("user_cookie")
		if cookie!=nil{
			sqlstr:="select sign from user where name=?;"
			var s string
			db.QueryRow(sqlstr,cookie.Value).Scan(&s)
			if len(s)>0{
				c.JSON(200,gin.H{
					"你曾说过：":s,
					"您好":cookie.Value+"!",
				})
			}else{
				fmt.Println(cookie.Value)
				c.JSON(200,gin.H{
					"您好":cookie.Value+"!",
				})
			}

		}else{
			c.JSON(200,gin.H{
				"您好":"游客!",
			})
		}

	}
}