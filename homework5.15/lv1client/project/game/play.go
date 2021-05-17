package game

import "fmt"

//出牌
//(参数为上家牌 要出牌 手牌)
func Play(shangjia Cards ,playcard ,handcard []string)(Cards,[]string){
	res:=Judge(playcard)
	//上家不出
	if shangjia.CT==Overwhelm&&res.CT!=Falsetype{
		handcard=Reduce(playcard ,handcard)
		return res,handcard
	}
	//牌型合法，牌型相同，牌数相同，牌比上家大
	if res.CT!=Falsetype&&res.CT==shangjia.CT&&res.Len==shangjia.Len&&res.Nmin>shangjia.Nmin{
		handcard=Reduce(playcard ,handcard)
		return res,handcard
	}
	fmt.Println("出牌不符合规则，请重新出牌")
	return Cards{CT: Falsetype},handcard
}
//手牌减去要出牌
func Reduce(playcard ,handcard []string)[]string{
	for i:=0;i<len(playcard);i++{
		for j:=0;j<len(handcard);j++{
			if playcard[i]==handcard[j]{
				if j+1==len(handcard){
					handcard=handcard[:j]
				}else{
					handcard=append(handcard[0:j],handcard[j+1:]...)
				}
				break
			}
		}
	}
	return handcard
}