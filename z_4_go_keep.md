#### go 版本更新 

brew update
brew install go


#### flag provided but not defined: -goversion
一个是版本原因, 一个是vscode也要修改配置, 坑爹

Thank you, I was able to solve this by running brew uninstall --force go and then downloading the latest installer. Anyone who reads this and wants to use brew you could probably just do brew install go after the forced uninstall. I had to restart my terminal and Gogland after doing this.



#### sql

1. redis
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


2. xorm
https://github.com/go-xorm/xorm



