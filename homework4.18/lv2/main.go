package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	fmt.Println("请输入学号")
	var id string
	fmt.Scanln(&id)
	if len(id)!=10{
		fmt.Println("学号格式错误")
		return
	}
	schedule(id)
}
func schedule(id string){
	url:="http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=2019211098"
	r,err:=http.NewRequest("GET",url,nil)
	if err!=nil{
		fmt.Printf("err:%v",err)
		return
	}
	c:=&http.Client{}
	r.Header.Set("accept-language","zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	r.Header.Set("user-agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36 Edg/90.0.818.41")
	response,err:=c.Do(r)
	if err!=nil{
		fmt.Printf("err:%v",err)
		return
	}
	body,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		fmt.Printf("err:%v",err)
		return
	}
	html:=string(body)
	exp:=`<br>[\w\s]+-(.*)<br>(.*)[\s\S]+?<br>(.*)<fo.*?>(.*)<\/font.*'>(.*)<\/span>`
	re:=regexp.MustCompile(exp)
	res:=re.FindAllStringSubmatch(html,-1)
	//res[][2]最后两个字符为\r，会消去前面的内容，故用切片去掉末尾两个字符（大坑！！！！）
	for i:=range res{
		fmt.Printf("课程名：%s    %s 周次：%s 备注：%s 任课老师：%s\n",
			res[i][1],res[i][2][:len(res[i][2])-2],res[i][3],res[i][4],res[i][5])
	}
}