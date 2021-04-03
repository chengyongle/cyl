package main
import "fmt"

func main() {
	//打印心
	var h heart
	Paint(h)
	//打印圆
	var c circle
	Paint(c)
}
//接口定义
type Painter interface {
	paint()
}
type heart struct{
	//心形结构体
}
type circle struct {
	//圆形结构体
}
//可调用任意实现了painter接口的结构中的paint函数
func Paint(p Painter){
	p.paint()
}
//心形函数实现打印方法
func (h heart)paint(){
	var a,b float64
	for i:=2.0;i>=-2.0;i-=0.1{
		for j:=-2.0;j<=2.0;j+=0.05{
			a=(j*j+i*i-1)*(j*j+i*i-1)*(j*j+i*i-1)
			b=j*j*i*i*i
			if a<=b{
				fmt.Printf("*")
			}else{
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
//圆形函数实现打印方法
func (c circle)paint(){
	var a,b float64
	for i:=1.0;i>=-1.0;i-=0.1{
		for j:=-2.0;j<=2.0;j+=0.05{
			a=1-i*i
			b=j*j
			if a>=b{
				fmt.Printf("*")
			}else{
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}