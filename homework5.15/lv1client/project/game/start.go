package game

import (
	"math/rand"
	"strconv"
	"time"
)

//生成一副牌
/*
牌值映射
黑桃A红桃B梅花C方块D 大小王X
3 4 5 6 7 8 9 10 J  Q  K  A  2 小王 大王
3 4 5 6 7 8 9 10 11 12 13 14 15 16 17
 */
//牌值映射
var m map[int]string=map[int]string{11:"J",12:"Q",13:"K",14:"A",15:"2",16:"小王",17:"大王"}
func CreatePoker()[]string{
	pokers:=make([]string,54)
	for i:=3;i<=16;i++{
		if i==16{
			pokers[(i-3)*4]="X16"
			pokers[(i-3)*4+1]="X17"
			break
		}
		pokers[(i-3)*4]="A"+strconv.Itoa(i)
		pokers[(i-3)*4+1]="B"+strconv.Itoa(i)
		pokers[(i-3)*4+2]="C"+strconv.Itoa(i)
		pokers[(i-3)*4+3]="D"+strconv.Itoa(i)
	}
	return pokers
}
//洗牌
func Shuffle(cards []string){
	rand.Seed(time.Now().Unix())
	for len(cards)>0{
		n:=len(cards)
		r:=rand.Intn(n)
		cards[r],cards[n-1]=cards[n-1],cards[r]
		cards=cards[:n-1]
	}
}
//发牌
//playerid 1-3为每个玩家初始牌，0为底牌，由地主所得
func Deal(playerid int,cards []string)[]string{
	if playerid<0||playerid>3{
		return nil
	}
	var playercards []string
	if playerid==0{
		playercards=cards[51:54]
		return playercards
	}else{
		n:=(playerid-1)*17
		playercards=cards[n:n+17]
	}
	return playercards
}
//牌字符转化为图案
func Transform(handcard []string)[]string{
	var result []string
	for _,str:=range handcard{
		result=append(result, transform(str))
	}
	return result
}
//转化单个字符串
func transform(card string)string{
	var res string
	switch card[0]{
	case 'A':res="♠"
	case 'B':res="♥"
	case 'C':res="♣"
	case 'D':res="♦"
	case 'X':
		nu,_:=strconv.Atoi(card[1:])
		res=m[nu]
		return res
	}
	n,_:=strconv.Atoi(card[1:])
	if n<=10{
		res=res+strconv.Itoa(n)
	}else{
		res=res+m[n]
	}
	return res
}
