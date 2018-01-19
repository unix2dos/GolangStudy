
#### go 版本更新 

brew update
brew install go

#### cannot find package "fmt" in any of:
unset GOROOT

#### flag provided but not defined: -goversion

一个是版本原因, 一个是vscode也要修改配置gopath, 坑爹

>Thank you, I was able to solve this by running brew uninstall --force go and then downloading the latest installer. Anyone who reads this and wants to use brew you could probably just do brew install go after the forced uninstall. I had to restart my terminal and Gogland after doing this.

#### vscode not jump define

```
        "go.useLanguageServer": true,
        "go.docsTool": "gogetdoc",
```

#### vscode could not launch process: exec: "lldb-server": executable file not found in $PATH

```
        xcode-select --install
```
#### Goto Definition, Hover, Signature Help do not work for Dot imported functions 

go 1.92 貌似不行, 已提交vscode issue

#### vscode jump slow
安装https://github.com/sourcegraph/go-langserver
源码安装 需要 go install
```
  "go.useLanguageServer": true,
```


#### sql

+ redis
github.com/garyburd/redigo/redis

key *  //查看redis所有key

del key //删除key

type key // 获取key type

操作:
string | get key, set key
list | lrange key 0 10, lpush key value
set | smembers key, sadd key value
zset | zrange key 0 10 withscores, zadd key value
hash | hgetall key | hset key value


+ sql xorm
https://github.com/go-xorm/xorm



### data
set 可以用 gopkg.in/fatih/set.v0


