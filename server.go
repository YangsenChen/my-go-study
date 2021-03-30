
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strings"
)
type user struct{
	Id int64
	Name string
	Age int64
}
var  (
	Router = mux.NewRouter()
)
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // 解析参数，默认是不会解析的
	fmt.Println(r.Form)  // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello a!") // 这个写入到 w 的是输出到客户端的
}
func sayhelloName1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // 解析参数，默认是不会解析的
	fmt.Println(r.Form)  // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello b!") // 这个写入到 w 的是输出到客户端的
}
func Select() []user{
	sql := "select * from user"
	rowz,err:=Dql(sql)
	if err!=nil{
		fmt.Println(err)
		return nil
	}
	users:=make([]user,0)
	for rowz.Next(){
		var u user
		rowz.Scan(&u.Id,&u.Name,&u.Age)
		//fmt.Println(u.name)
		users=append(users,u)
	}
	CloseConn()
	return users
}

func sayhelloName2(w http.ResponseWriter, r *http.Request) {
	us := Select()
	//fmt.Println(us)
	jus,_:=json.Marshal(us)
	fmt.Println(string(jus))
	w.Header().Set("Content-Type","application/json;charset=utf-8")
	//w.Write(jus)
	fmt.Fprintf(w, string(jus))
}
func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("hhh.html")
	t.Execute(w, nil)
}
func main() {
	http.HandleFunc("/",welcome)

	http.HandleFunc("/a", sayhelloName) // 设置访问的路由
	http.HandleFunc("/b", sayhelloName1) // 设置访问的路由
	http.HandleFunc("/q", sayhelloName2) // 设置访问的路由

	err := http.ListenAndServe(":9090", nil) // 设置监听的端口,handler is nil which means using default servemux
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}