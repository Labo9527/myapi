# myapi

## 实验目标

规范：REST API 设计 [Github API v3 overview](https://developer.github.com/v3/) ；[微软](https://docs.microsoft.com/zh-cn/azure/architecture/best-practices/api-design)
作业：模仿 Github，设计一个博客网站的 API

## 实验内容

实现功能有：

1. 实现了用户的创建
2. 实现了用户的查询
3. 实现了单个用户的指定查询
4. 实现了用户的删除
5. 实现了文章的创建
6. 实现了文章阅读量机制
7. 实现了文章的查询
8. 实现了指定作者的文章查询
9. 实现了删除文章功能

### 总体描述

我的程序利用了`gorilla/mux`包，参考了[博客](https://www.cnblogs.com/oxspirt/p/10863154.html)，并且实现了自己的特定类似博客网站的功能（目前只有用户和文章，后续考虑加入评论，为了先交作业。。。），整体设计思路符合无状态这一思路，在服务器只存储数据库信息，不存储状态信息

### 用户的创建

用户在我的程序中使用的结构体如下：

```go
type User struct{
	Id int
	Username string
	Password string
	Mail string
}
```

使用方法是用`POST`访问 `localhost:9090/users`  并提交相应`json`信息

![image-20191122233612945](/Users/yang/Library/Application Support/typora-user-images/image-20191122233612945.png)

服务端会返回创建用户的信息，每个用户用唯一表示符号Id表示，因此允许重名

```go
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var myuser User
	_ = json.NewDecoder(r.Body).Decode(&myuser)
	if myuser.Username==""||myuser.Password==""||myuser.Mail=="" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Please Enter username, password and mail")
		return
	}
	myuser.Id=len(users)+1
	users = append(users, myuser)
	json.NewEncoder(w).Encode(myuser)
}
```



### 用户的查询

![image-20191122233806193](/Users/yang/Library/Application Support/typora-user-images/image-20191122233806193.png)

使用`POST`方法访问`localhost:9090/users`会返回所有用户`json`信息，但没有密码

```go
func QueryUser(w http.ResponseWriter, r *http.Request) {
	for _, user := range users{
		var mypre Presentuser
		mypre.Username=user.Username
		mypre.Mail=user.Mail
		mypre.Id=user.Id
		json.NewEncoder(w).Encode(mypre)
	}
}
```

### 单个用户查询

使用`GET`方法来访问`localhost:9090/user/{id}`，会返回特定用户`json`信息

![image-20191122234128739](/Users/yang/Library/Application Support/typora-user-images/image-20191122234128739.png)

```go
func QueryAuser(w http.ResponseWriter, r *http.Request){
	inputs := mux.Vars(r)
	fmt.Println(inputs["id"],inputs["xx"])
	for _, user := range users{
		if strconv.Itoa(user.Id)!=inputs["id"] {
			continue
		}
		var mypre Presentuser
		mypre.Username=user.Username
		mypre.Mail=user.Mail
		mypre.Id=user.Id
		json.NewEncoder(w).Encode(mypre)
		break
	}
}
```



### 用户删除

