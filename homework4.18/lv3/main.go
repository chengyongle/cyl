package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)
//在OPGG上查找某个位置的英雄强度排行
func main() {
	fmt.Println("请选择要查询的位置")
	fmt.Println("1.上路")
	fmt.Println("2.打野")
	fmt.Println("3.中路")
	fmt.Println("4.ADC")
	fmt.Println("5.辅助")
	var s int
	fmt.Scanln(&s)
	if s>5||s<1{
		fmt.Println("输入数字无效")
		return
	}
	switch s{
	case 1:query("TOP","上单")
	case 2:query("JUNGLE","打野")
	case 3:query("MID","中单")
	case 4:query("ADC","下路")
	case 5:query("SUPPORT","辅助")
	}
}
func query(p,q string){
	url:="https://www.op.gg/champion/statistics"
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
	//先找出单个位置部分的网页代码，因为有的AD可以打上单，在查询“下路”时就会出现在上路的胜率信息
	exp1:=`<tbody class="tabItem champion-trend-tier-`+p+`"[\s\S]+?<\/tbody>`
	re1:=regexp.MustCompile(exp1)
	pres:=re1.FindAllString(html,-1)
	//再从代码中找全部该位置英雄信息
	exp2:=`<.*table__name">(.*)<\/div>\s*<.*position">\s*`+q+`.*?[\s.]*<\/div>[^0-9]*(.*)%[^0-9]*(.*)%[^0-9]*(.*)\.`
	re2:=regexp.MustCompile(exp2)
	res:=re2.FindAllStringSubmatch(pres[0],-1)
	for i:=range res{
		fmt.Printf("%-20s 胜率：%s%%\t登场率：%s%%\t级别：T%s\n",res[i][1],res[i][2],res[i][3],res[i][4])
	}
}
