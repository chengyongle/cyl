package main

import (
	"fightlandlords/game"
	"fmt"
)
//const (
//	cardType = iota
//	Single          //单张
//	Double          //对子
//	Triple         	//三不带
//	Tripleone       //三带一   4
//	Tripletwo       //三带二   5
//	Bomb            //炸弹
//	Bombchain		//连炸
//	Fourt         	//四带二
//	Fourdouble		//四带一对
//	Fourtdouble		//四带两对
//	Fourtchain		//连四带二
//	Fourdoublechain	//连四带一对
//	Fourtbchain		//连四带两对
//	Plane           //无翅飞机  14
//	Planeone        //三带一飞机  15
//	Planetwo	    //三带二飞机  16
//	Splane			//特殊飞机  17
//	Dragon    		//顺子   18
//	DoubleDragon    //连对   19
//	Kingbomb        //王炸  20
//	Falsetype       //不合规牌型  21
//  Overwhelm  		//要不起
//)
func main() {
	a:=[]string{"A12","D12","B12","C12","A11","C11","B11","D11","D10","C10","A10","B10"}
	b:=[]string{"B12","A12","D11","D10","D12"}
	fmt.Println(game.Reduce(b,a))
}
