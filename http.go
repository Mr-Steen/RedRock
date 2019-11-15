package main
import "fmt"
import "net/http"
import "io/ioutil"
func main(){
        response,err:=http.Get("http://cn.bing.com")
        if (err!=nil){
                fmt.Println(err)
        }
        defer response.Body.Close()
        body,err:=ioutil.ReadAll(response.Body)
        fmt.Println(string(body))

}

