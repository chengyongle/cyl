package  models

import (
	"fightlandlords/dao"
	"fmt"
	"github.com/jinzhu/gorm"
)
func Register(name string,password string)int{
	var user dao.User
	err:=dao.DB.Where("username=?",name).Find(&user).Error
	if err==nil{
		//账号已存在
		return 0
	}else if err==gorm.ErrRecordNotFound{
		//账号不存在
		user.Username=name
		user.Password=password
		dao.DB.Create(&user)
		return 1
	}else{
		//其他错误
		fmt.Printf("err:%v",err)
		return 2
	}

}
func Login(name string,password string)int{
	var user dao.User
	err:=dao.DB.Where("username=?",name).Find(&user).Error
	if err==nil{
		//账号存在
		if user.Password==password{
			//登陆成功
			return 3
		}else{
			//密码错误
			return 2
		}
	}else if err==gorm.ErrRecordNotFound{
		//账号不存在
		return 0
	}else{
		//查询异常
		fmt.Printf("err:%v",err)
		return 1
	}

}
