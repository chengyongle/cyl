package game

import (
	"sort"
	"strconv"
)

//判断连续性(参数p表示牌型最小连续张数，如顺子为5，连对为3，飞机为2)
func continuity(maxn ,p int,m map[int]int)bool{
	var a []int
	for key,value:=range m{
		if value==maxn{
			a=append(a,key)
		}
	}
	sort.Ints(a)
	n:=len(a)
	l:=a[0]
	r:=a[n-1]
	if r-l==n-1&&r-l>=p-1&&r<=14{
		return true
	}else {
		return false
	}
}
//检验是否纯净(连对不能带单张，飞机不能只带一张单张牌,4带2不能乱带3，不能带俩王)，参数p为不能带的同值牌张数
func pure(p int,m map[int]int)bool{
	nw:=0//大小王数量
	for key,value:=range m{
		if value==p{
			return false
		}
		if key==16||key==17{
			nw++
		}
	}
	return nw!=2
}
//检验飞机合法性，一组三张同值牌带一个单张或者对，x表示单张或对
func plane(x int,m map[int]int)bool{
	nx:=0//x数量
	nt:=0//刻子数量
	nw:=0//大小王数量
	for key,value:=range m{
		if key==16||key==17{
			nw++
		}
		if (x==1&&value==2)||(x==2&&value==1){//3带2不能有1  3带1不能有2
			return false
		}
		if value==x{
			nx++
		}
		if value==3{
			nt++
		}
	}
	return (nx==nt)&&(nw!=2)//不能带俩王
}
//特殊飞机验证33344466
func splane(m map[int]int)bool{
	nt:=0//刻子数量
	nd:=0//对子数量
	for _,value:=range m{
		if value==1{//不能有单张
			return false
		}
		if value==2{
			nd++
		}
		if value==3{
			nt++
		}
	}
	return (2*nd)==nt
}
//验证四带多合法性 x表示带单张数量,y表示带对数量,xy必有一个为0
func four(x,y int,m map[int]int)bool{
	n:=0//带牌数量
	nw:=0//大小王数量
	nf:=0//炸弹数量
	for key,value:=range m{
		if key==16||key==17{
			nw++
		}
		if value==4{
			nf++
		}
		if (x==2&&value==2)||(value==1&&y!=0){
			return false
		}
		if value==1||value==2{
			n++
		}
	}
	if nw==2{
		return false
	}
	return (x==2&&(2*nf)==n)||(y==1&&nf==n)||(y==2&&(2*nf)==n)
}
//手牌排序
func Sort(handcard []string){
	var a,b int
	n:=len(handcard)
	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			a,_=strconv.Atoi(handcard[i][1:])
			b,_=strconv.Atoi(handcard[j][1:])
			if a>b{
				handcard[i],handcard[j]=handcard[j],handcard[i]
			}
		}
	}
}