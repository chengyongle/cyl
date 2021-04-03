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
	v:=release("大司马","起飞")
	fmt.Println(v)

}
func (v *video)like(){
	v.thumbs_up++
}
func(v *video)collect(){
	v.collection++
}
func(v *video)put(){
	v.coin++
}
func(v *video)triple(){
	v.thumbs_up++
	v.collection++
	v.coin++
}
func release(aname,vname string)video{
	a:=Author{Name: aname}
	return video{title: aname,author: a}
}