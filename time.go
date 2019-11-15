package main
import "fmt"
import "time"
//import "strconv"
func main(){
        //var sli []int64
        var map1 map[int64]string=make(map[int64]string)
        //fmt.Println(time.Now().Format("2006-01-02 15:04:04"))
        //fmt.Println(time.Now().Unix())
        var a int64
        for{
                fmt.Printf("please input a time time:")
                fmt.Scanf("%d",&a)
                if a>0{
                        //sli=append(sli,a)
                        //var z string
                        //z=time.Unix(a,1).Format("")
                        //slic=append(slic,strconv.Atoi(z))
                        //i:=0
                        map1[a]=time.Unix(a,1).Format(" ")
                        fmt.Printf("input ok!\n")
                        fmt.Println(time.Unix(a,1))
                }else{
                        fmt.Printf("error")
                }
                //fmt.Println(sli)
                //fmt.Println(slic)
                var s string
                fmt.Printf("do you wanna to continue?(Y/N):")
                fmt.Scanf("%s",&s)
                if s=="N" || s=="n"{
                        //fmt.Println(sli)
                        for i,j:=range map1{
                                fmt.Printf("map[%d]:%s",i,j)
                        }
                        break
                }
        }

}

