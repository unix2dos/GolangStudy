# package

* 为什么golang编译快
    + 所有包都要在头顶声明, 用的都写上
    + 禁止包的环状依赖, 有向无环图
    + 编译后文件记录导出信息,还记录依赖关系
    
* 默认的包名是最后一个, 有三个例外
    + main包
    + _test.go结尾的文件, 包名也是以 _test 后缀
    + 后面带版本号的
    
* 下载包
    go get -u命令只是简单地保证每个包是最新版本，如果是第一次下载包则是比较方便的；但是对于发布程序则可能是不合适的，因为本地程序可能需要对依赖的包做精确的版本依赖管理。

* 内部包

 1. Go语言的构建工具对包含internal名字的路径段的包导入路径做了特殊处理。这种包叫internal包，一个internal包只能被和internal目录有同一个父目录的包所导入。
 
 2. net/http/internal/chunked内部包只能被net/http/httputil或net/http包导入，但是不能被net/url包导入。不过net/url包却可以导入net/http/httputil包。

 ```
 net/http
net/http/internal/chunked
net/http/httputil
net/url
 ```

* 查询包

```
go list reading-friends-api/...

reading-friends-api
reading-friends-api/controller
reading-friends-api/controller/payment
reading-friends-api/model
reading-friends-api/model/dic
reading-friends-api/router
reading-friends-api/service
reading-friends-api/utils
```

```
go list -json reading-friends-api //以json格式打开
```




