package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)
type User struct {
	Name string `form:"username"`
	Password string `form:"password"`

}
type Userfile struct {
	Username string `form:"username"`
	Filename string `form:"filename"`
	Fileintroduction string `form:"fileintroduction"`
	Filepath string `form:"filepath"`
}
func login(c *gin.Context){
	var m=make(map[string]User)
	f,err:=os.OpenFile("./users.json",os.O_CREATE|os.O_RDWR,0644)
	if err!=nil{
		fmt.Printf("open file err:%v",err)
		return
	}
	defer f.Close()
	var bytes []byte
	bytes,err=ioutil.ReadAll(f)
	if err!=nil{
		fmt.Printf("err%v",err)
		return
	}
	json.Unmarshal(bytes,&m)
	var user User
	c.ShouldBind(&user)
	for k,v:=range m{
		if k==user.Name{
			if v.Password==user.Password{
				c.SetCookie("usercookie",user.Name,1000,"/",
					"localhost",false,true)
				c.Writer.Write([]byte("登录成功"))
				return
			}else{
				c.Writer.Write([]byte("密码错误"))
				return
			}
		}
	}
	c.Writer.Write([]byte("该账号未注册"))
}
func register(c *gin.Context){
	var m=make(map[string]User)
	f,err:=os.OpenFile("./users.json",os.O_CREATE|os.O_RDWR,0644)
	if err!=nil{
		fmt.Printf("open file err:%v",err)
		return
	}
	defer f.Close()
	var bytes []byte
	bytes,err=ioutil.ReadAll(f)
	if err!=nil{
		fmt.Printf("err:%v",err)
		return
	}
	json.Unmarshal(bytes,&m)
	var user User
	c.ShouldBind(&user)
	for _,v:=range m{
		if v.Name==user.Name{
			c.Writer.Write([]byte("账号已存在，请登录"))
			return
		}
	}
	m[user.Name]=user
	bytes,err=json.Marshal(m)
	if err!=nil{
		fmt.Printf("err:%v",err)
		return
	}
	os.Truncate("./users.json",0)
	f.Seek(0,0)
	_,err=f.WriteString(string(bytes))
	if err!=nil{
		fmt.Printf("err:%v",err)
		return
	}
	c.SetCookie("usercookie",user.Name,1000,"/","localhost",false,true)
	c.Writer.Write([]byte("注册成功"))
}
func Middleare() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Next()
		cookie,_:=c.Request.Cookie("usercookie")
		if cookie==nil{
			c.JSON(200,gin.H{
				"message":"游客你好！",
			})
		}else{
			c.JSON(200,gin.H{
				"message":cookie.Value+"你好！",
			})
		}

	}
}
func upload(c *gin.Context){
	cookie,_:=c.Request.Cookie("usercookie")
	if cookie==nil{
		c.Writer.Write([]byte("请先登录"))
		return
	}
	var m=make(map[string]Userfile)
	f,err:=os.OpenFile("./userfile.json",os.O_CREATE|os.O_RDWR,0644)
	if err!=nil{
		fmt.Printf("open file err:%v",err)
		return
	}
	defer f.Close()
	var bytes []byte
	bytes,err=ioutil.ReadAll(f)
	if err!=nil{
		fmt.Printf("err:%v",err)
		return
	}
	json.Unmarshal(bytes,&m)
	var uf Userfile
	file, err := c.FormFile("file")
	if err != nil {
		c.Writer.Write([]byte("获取文件失败"))
		return
	}
	dst:="./"+file.Filename
	c.SaveUploadedFile(file,dst)
	uf.Username=c.PostForm("username")
	uf.Filename=file.Filename
	uf.Fileintroduction=c.PostForm("fileintroduction")
	uf.Filepath=dst
	m[uf.Username]=uf
	bytes,err=json.Marshal(m)
	if err!=nil{
		fmt.Printf("err:%v",err)
		return
	}
	os.Truncate("./users.json",0)
	f.Seek(0,0)
	_,err=f.WriteString(string(bytes))
	if err!=nil{
		fmt.Printf("err:%v",err)
		return
	}
	c.Writer.Write([]byte("文件上传成功"))
}
func draw(c *gin.Context){
	cookie,_:=c.Request.Cookie("usercookie")
	if cookie==nil{
		c.Writer.Write([]byte("请先登录"))
		return
	}
	var m=make(map[string]Userfile)
	f,err:=os.OpenFile("./userfile.json",os.O_CREATE|os.O_RDWR,0644)
	if err!=nil{
		fmt.Printf("open file err:%v",err)
		return
	}
	defer f.Close()
	var bytes []byte
	bytes,err=ioutil.ReadAll(f)
	if err!=nil{
		fmt.Printf("err:%v",err)
		return
	}
	json.Unmarshal(bytes,&m)
	//生成随机种子
	rand.Seed(time.Now().Unix())
	s:=rand.Intn(len(m))
	for k,_:=range m{
		if s==0{
			str:="恭喜"+k+"这个比中奖"
			c.Writer.Write([]byte(str))
			c.JSON(200,gin.H{
				"作品名称":m[k].Filename,
				"作品简介":m[k].Fileintroduction,
				"作品路径":m[k].Filepath,
			})
		}
		s--
	}
}