package main
import "fmt"
var myRes=make(map[int]int,20)
func factorial(n int){
        var res=1
        ch:=make(chan int,1)
        for i:=1;i<=n;i++{
                res*=i
        }
        ch<-1
        myRes[n]=res
        <-ch
}
func main(){
        for i:=1;i<=20;i++{
                go factorial(i)
        }
        ch:=make(chan int,1)
        ch<-1
        for i,v:=range myRes{
                fmt.Printf("myres[%d}=%d\n",i,v)
        }
        <-ch
}
