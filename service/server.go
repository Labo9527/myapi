package service

import (
	"fmt"
	"github.com/go-martini/martini" //使用martini框架
)

func second1() string {	//分别设定几个处理函数 second1, second2, second3
	return "second1"
}

func second2() string{
	return "second2"
}

func second3() string{
	return "second3"
}

func NewServer(port string) {
	m := martini.Classic()	//创建实例

	m.Group("/first", func(r martini.Router) {	//	采用了路由分组处理的方法
		r.Get("/second1", second1)
		r.Get("/second2", second2)
		r.Get("/second3", second3)
	})

	m.Use(func(c martini.Context) {
		fmt.Println("i am middle man")
	})

	m.Get("/", func(params martini.Params) string {	//普通路由处理
		return "hello world"
	})

	m.Get("/tom", func(params martini.Params) string {	//普通路由处理
		return "hello tom"
	})

	m.RunOnAddr(":"+port)	//监听指定端口
}


/*

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)


type User struct{
	Username string
	Password string
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	filename := "users/user1.json"
	fp, _ := os.OpenFile(filename, os.O_RDONLY, 0755)
	data := make([]byte, 100)
	n, _ := fp.Read(data)
	myuser := new(User)
	json.Unmarshal(data[:n],&myuser)
	myjson, _:=json.Marshal(myuser)


	fmt.Fprintf(w, string(myjson)) //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


 */