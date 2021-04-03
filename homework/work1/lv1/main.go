package main

import "fmt"

func main() {
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
