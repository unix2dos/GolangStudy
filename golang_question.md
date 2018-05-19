### go 版本更新 

brew update
brew install go

### cannot find package "fmt" in any of:
unset GOROOT


### sql

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


