package  models
import (
	"awesomeProject/dao"
	"database/sql"
	"fmt"
)
type Message struct{
	//Id int
	Name string
	Text string
	//Count int
	//ParentId int//回复信息的id
}
//需要创建一个message数据表
//包含id,name,text,点赞数count(默认为0),父信息标记parentid(默认为-1)
func SaveMessage(name string,message string)int{
	stmt,err:=dao.DB.Prepare("insert into comment (name,text)value(?,?);")
	if err!=nil{
		fmt.Printf("mysql prepare failed:%v",err)
		return 1
	}
	defer stmt.Close()
	_,err=stmt.Exec(name,message)
	if err!=nil{
		fmt.Printf("insert message failed:%v",err)
		return 1
	}
	return 2
}
func ReplyMessage(name string,message string,pid int)int{
	stmt,err:=dao.DB.Prepare("insert into reply (name,text,parentid) value(?,?,?);")
	if err!=nil{
		fmt.Printf("mysql prepare failed:%v",err)
		return 1
	}
	defer stmt.Close()
	_,err=stmt.Exec(name,message,pid)
	if err!=nil{
		fmt.Printf("insert reply failed:%v",err)
		return 1
	}
	return 2
}
func ViewMessage(name string)[]Message{
	stmt,err:=dao.DB.Prepare("select  name,text,pid from message where name=?;")
	if err!=nil{
		return nil
	}
	rows,err:=stmt.Query(name)
	if err==sql.ErrNoRows{
		return nil
	}
	defer rows.Close()
	var mes []Message
	for rows.Next(){
		var message Message
		var pid int
		rows.Scan(&message.Name,&message.Text,&pid)
		mes=append(mes,message)
		
	}
	return mes
}
func Interesting(id int)int{
	stmt,err:=dao.DB.Query("select count from message where id=?;",id)
	if err==sql.ErrNoRows{
		fmt.Println("id不存在")
		return 1
	}
	defer stmt.Close()
	var newcount int
	err=stmt.Scan(&newcount)
	if err!=nil{
		fmt.Printf("scan failed:%v",err)
		return 1
	}
	newcount++
	_,err=dao.DB.Exec("update message set count =? where id=?;",newcount,id)
	if err!=nil{
		fmt.Printf("update failed:%v",err)
		return 1
	}
	return 2
}
