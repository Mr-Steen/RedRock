package main
import "fmt"
type dear struct{
	nick string
	age int
	master string
}
type dove interface{
	gugugu()
}
type repeater interface{
	repeat(str string)
}
type lemon interface{
	su()
}
type xiangxiang interface{
	xiangxiang()
}
func (darling *dear) gugugu(){
	fmt.Printf("gugugu")
}
func (darling *dear) repeat(str string){
	fmt.Printf("%s",str)
}
func (darling *dear)su(){
	fmt.Printf("soury")
}
func (darling *dear)xiangxiang(){
	fmt.Printf("smell good!")
}
func main(){
	darling:=dear{"xiao ai",12,"me"}
	fmt.Println(darling.nick,darling.age,darling.master)
	darling.gugugu()
	darling.repeat("i love U\nkiss,kiss,daisiki")
	darling.xiangxiang()
	darling.su()
}
