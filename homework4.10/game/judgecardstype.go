package game

import (
	"strconv"
)

//牌型枚举类型
type CardType int

const (
	cardType = iota
	Single          //单张
	Double          //对子
	Triple         	//三不带
	Tripleone       //三带一
	Tripletwo       //三带二
	Bomb            //炸弹
	Bombchain		//连炸
	Fourt         	//四带二
	Fourdouble		//四带一对
	Fourtdouble		//四带两对
	Fourtchain		//连四带二
	Fourdoublechain	//连四带一对
	Fourtbchain		//连四带两对
	Plane           //无翅飞机
	Planeone        //三带一飞机
	Planetwo	    //三带二飞机
	Splane			//特殊飞机
	Dragon    		//顺子
	DoubleDragon    //连对
	Kingbomb        //王炸
	Falsetype       //不合规牌型
	Overwhelm  		//要不起
)
//判断牌张数，最大同值牌数，同值牌最小值,并返回上述三值和牌值-数量map
func judge(cards []string)(n,maxn,nmin int,m map[int]int){
	n=len(cards)
	maxn=0
	nmin=18
	m=make(map[int]int)
	var a int
	for _,value:=range cards{
		a,_=strconv.Atoi(value[1:])
		m[a]++
	}
	for _,value:=range m{
		if value>=maxn{
			maxn=value
		}
	}
	for key,value:=range m{
		if value==maxn&&key<=nmin{
			nmin=key
		}
	}
	return n,maxn,nmin,m
}
//抽象出牌结构体
type Cards struct{
	Value 	[]string//牌面
	Len 	int//牌张数
	Maxn	int//最大同值牌数
	Nmin	int//同值牌最小值，用于比较大小
	CT		CardType//牌型
}
//判断牌型
func Judge(cards []string)Cards{
	Sort(cards)
	n,maxn,nmin,m:=judge(cards)//牌张数n，最大同值牌数maxn，同值牌最小值nmin
	if n==0{
		return Cards{cards,0,0,0,Overwhelm}
	}
	switch maxn{
	case 1://单张，王炸，顺子
		if n==1{//单张
			return Cards{cards,n,maxn,nmin,Single}
		}else if m[16]==1&&m[17]==1&&n==2{
			//王炸
			return Cards{cards,999,999,999,Kingbomb}
		}else if continuity(maxn,5,m){
			//顺子
			return Cards{cards,n,maxn,nmin,Dragon}
		}else{
			//错误类型
			return Cards{CT: Falsetype}
		}
	case 2://对子，连对
		if n==2{
			return Cards{cards,n,maxn,nmin,Double}
		}else if pure(1,m)&&continuity(maxn,3,m){
			return Cards{cards,n,maxn,nmin,DoubleDragon}
		}else{
			//错误类型
			return Cards{CT: Falsetype}
		}
	case 3://3不带，3带一，3带二，无翅飞机，3带一飞机，3带二飞机，特殊飞机（33344455）
		if n==3{//3不带
			return Cards{cards,n,maxn,nmin,Triple}
		}else if n==4{//3带1
			return Cards{cards,n,maxn,nmin,Tripleone}
		}else if n==5&&pure(1,m){//3带2
			return Cards{cards,n,maxn,nmin,Tripletwo}
		} else if continuity(maxn,2,m)&&pure(1,m)&&pure(2,m){
			//无翅飞机
			return Cards{cards,n,maxn,nmin,Plane}
		}else if continuity(maxn,2,m)&&plane(1,m){
			//3带一飞机
			return Cards{cards,n,maxn,nmin,Planeone}
		}else if continuity(maxn,2,m)&&plane(2,m){
			//3带二飞机
			return Cards{cards,n,maxn,nmin,Planetwo}
		}else if continuity(maxn,2,m)&&splane(m){//特殊飞机
			return Cards{cards,n,maxn,nmin,Splane}
		}else{
			//错误类型
			return Cards{CT: Falsetype}
		}
	case 4://炸弹，连炸，四带2，四带一对，四带两对，连四带2，连四带一对，连四带两对
		if n==4{//炸弹
			return Cards{cards,n,maxn,nmin,Bomb}
		}else if continuity(maxn,2,m)&&pure(1,m)&&pure(2,m)&&pure(3,m){
			//连炸
			return Cards{cards,n,maxn,nmin,Bombchain}
		}else if n==6&&pure(2,m){
			//四带二
			return Cards{cards,n,maxn,nmin,Fourt}
		}else if n==6&&pure(1,m){
			//四带一对
			return Cards{cards,n,maxn,nmin,Fourdouble}
		}else if n==8&&pure(1,m) {
			//四带两对
			return Cards{cards, n, maxn, nmin, Fourtdouble}
		}else if continuity(maxn,2,m)&&four(2,0,m){
			//连四带二
			return Cards{cards, n, maxn, nmin, Fourtchain}
		}else if continuity(maxn,2,m)&&four(0,1,m){
			//连四带一对
			return Cards{cards, n, maxn, nmin, Fourdoublechain}
		}else if continuity(maxn,2,m)&&four(0,2,m){
			//连四带两对
			return Cards{cards, n, maxn, nmin, Fourtbchain}
		}else{
			//错误类型
			return Cards{CT: Falsetype}
		}
	}
	return Cards{CT: Falsetype}
}

