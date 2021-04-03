package main

import "fmt"

func main() {
	a:=5
	b:=1.0
	c:="wfes"
	d:=false
	e:=-12.6
	receive(a)
	receive(b)
	receive(c)
	receive(d)
	receive(e)
}
func receive(v interface{}){
	switch v.(type){
	case int:
		fmt.Println("这是int类型")
	case float32:
		fmt.Println("这是float32类型")
	case float64:
		fmt.Println("这是float64类型")
	case byte:
		fmt.Println("这是byte类型")
	case string:
		fmt.Println("这是string类型")
	case bool:
		fmt.Println("这是bool类型")
	}
}
