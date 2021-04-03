package main

import "fmt"
type Author struct {
	Name string             //名字
	VIP bool                //是否是高贵的带会员
	Icon string             //头像
	Signature string        //签名
	Focus int               //关注人数
}
type video struct {
	title string
	author Author
	comment []string
 	barrage_list []string
	label []string
	thumbs_up int
	coin int
	collection int
	forward int
	Related_recommendations []video
}
func main(){
	a:=Author{
		Name:      "艺能人金广发",
		VIP:       false,
		Icon:      "nil",
		Signature: "提供《私家观察》《温情惧场》《都市传说》等系列、油管：艺能人金广发",
		Focus:     21,
	}
	v1:=video{title: "《都市传说》冷雨夜我不想归家",author:
	a,label:[]string{"生活","搞笑"},thumbs_up: 7434,coin: 8455,collection: 545,forward:54}
	v2:=video{title: "《都市传说》一场由空调引起的倾谈",author: a,
		label:[]string{"生活","搞笑"},thumbs_up: 8434,
		coin: 1007,collection: 634,forward:600,Related_recommendations: []video{v1}}
	fmt.Println(v2)

}
