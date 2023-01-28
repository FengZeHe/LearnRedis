# Redis学习笔记

## NoSQL概述
NoSQL的意思是No only Sql,Nosql因为没有IO操作，数据存储在内存中，因此读写速度飞快
。NoSQL属于非关系型数据库，适用于 **对数据高并发读写、海量数据的读写、对数据高可扩展性**的场景。

## Redis概述＆安装
### 什么是Redis
Redis是一个开源的key-value数据库。Redis和Memcacahed类似，支持多种数据类型，如string,list,set,zet,hash；
Redis数据支持push/pop、add/remove操作并且都是原子性的；Redis支持Master-Slave主从同步。



### Redis的安装

### 在ubuntu下安装
    sudo apt update
    sudo apt upgrade
    sudo apt install redis-server

检查是否安装成功＆查看运行状态

    redis-cli --version
    
    systemctl status redis

### 使用Docker运行
    // 查看redis版本
    docker search redis
    
    // 获取最新的redis版本
    docker pull redis:latest
    
    // 查看本地镜像
    docker images
    
    // 运行redis的docker容器
    docker run -itd --name my-redis -p 6379:6379 redis
    
    // 查看容器运行状态
    docker ps
    
    // 用交互模式进入容器
    docker exec -it my-redis /bin/bash

### 编译安装
    //下载源码文件
    wget https://download.redis.io/redis-stable.tar.gz
    
    // 编译redis
    // 切换到根目录,然后运行make
    tar -xzvf redis-stable.tar.gz
    cd redis-stable
    make
    
    // 如果编译成功，则在src目录下多出两个redis二进制文件
    redis-server : Redis 服务器本身
    redis-cli : Redis 对话的命令行界面实用程序
    
    // 在 /usrl/local/bin 下编译二进制文件
    make run 
    
    // 启动redis
    redis-server


​    
## Redis的基本操作
### Redis的key操作
1. 查看库当前的key `keys *`
2. 判断某个key是否存在 `exists key`
   1. key存在 返回 （integer）1
   2. key不存在 返回 （integer） 0
3. 查看key的类型 `type key`
   1. key存在 返回 具体类型（string）
   2. key不存在 返回 none 
4. 删除指定key `del key`
   1. 删除成功 key 返回 (integer) 1
   2. 删除失败 key 返回 (integer) 0
5. 根据value选择非阻塞 `unlink key`
   1. 删除成功 key 返回 (integer) 1
   2. 删除失败 key 返回 (integer) 0
6. 设置key的过期时间 `expire key 10(s)`
   1. 设置成功 key 返回 (integer) 1
   2. 设置失败 key 返回 (integer) 0
7. 获取redis key过期时间
   1. `ttl key` 返回过期时间的秒数，如果不存在过期时间则返回-1，如果key不存在返回-2。
   2. `pttl key` 返回过期时间的毫秒数；其余跟`ttl key`相同。
8. 移除redis key的过期时间
   1. `persist key` 


## Redis五种数据类型
### String 字符串

#### String概述
String类型是二进制安全性的，意味着redis的string可以包含任何数据，如jpg或序列化对象；
在redis字符串中value最大可以是512M。

#### 什么是二进制安全
二进制安全的意思就是，只关心二进制化的字符串，不关心具体格式，只会严格的按照二进制的数据存取，不会妄图以某种特殊格式解析数据

#### 关于redis string的数据结构
内部为当前字符串分配空间capacity一般高于实际字符串长度len；当字符串长度小于1M时，扩容会加倍当前空间。
如果字符串长度超过1M,那么每次扩容只会增加1M空间。

#### Redis String 操作
##### SET
* `SET key value` 将字符串值关联到key,set相同的key会覆盖之前的value
* 可选参数
  * EX seconds  `SET key value EX seconds` 设置value的同时设置过期时间。
  * NX  `SET key value NX` 仅当key不存在时才能设置;执行成功返回1，否则返回0。
    * 设置成功返回OK,失败则返回nil

##### SETNX
* `SETNX key value`
* 仅当key不存在时才能设置成功

##### SETEX
* `SETNX key seconds value`
* 设置关联value，同时设置key的过期时间。

##### PSETEX
* `PSETEX key milliseconds value`
* 设置关联value，同时设置key的过期时间（微秒）

##### GET
* `GET key`
* 获取key关联的value值

##### GETSET
* `GETSET key value`
* 返回key被设置前的值，将key的值设为value

##### STRLEN
* `STRLEN key`
* 返回字符串value的长度

##### APPEND
*  `APPEND key value`
* 若key存在则在关联的值后追加value并返回字符串长度
* 若key不存在则跟 `SET`关键字作用一样

##### SETRANGE
* `SETRANGE key offset value`
* 从偏移量`offset`开始，用value参数覆写key存储的字符串value值
* 简单来说，偏移到哪位，就将这位的值设置为value


##### GETRANGE
* `GETRANGE key start end`
* 从key存储的字符串中截取从start到end的字符串

##### MSET
* `MSET key value key value`
* 同时设置多个key value
* 如果某个key已经存在，那么`MSET`将使用新值覆盖旧值

##### MGET
* `MGET key key key`
* 同时获取多个key的value

##### INCR 
* `INCR key`
* 将key中存储值数字加1（只能用于数字）

##### DECR
* `DECR key`
* 将key中存储值数字减1（只能用于数字）

##### INCRBY
* `INCRBY key value`
* 自定义增量（可以为负值）

### Set 集合
* Set 是一个集合
* Set 对外提供的功能与list类似是一个列表的功能，特殊之处在于set是可以自动重排的。
* Set有去重的功能

#### Redis SET操作
##### SADD
* `SADD key value value value ...`
* 将一个或多个member元素加入到集合key当中，如member已经存在将被忽略。
* 若key不存在，则创建一个只包含member元素成员的集合。

##### SMEMBERS 
* `SMEMBERS key`
* 查看集合key中所有成员，不存在的key被视为空集合
* 如查看不存在key k2

` 
    127.0.0.1:6379> smembers k2
    (error) WRONGTYPE Operation against a key holding the wrong kind of value
`

##### SISMEMBER 
* `SISMEMBER key member`
* 判断member是否在集合key的成员，若是返回1，否则返回0

##### SPOP
* `SPOP key`
* 移除并`随机`返回集合中的一个元素

##### SRANDMEMBER
* `SRANDMEMBER key [count]`
* 若不提供参数 `count` ,则返回一个随机元素
* 若提供参数 `count` 小于集合基数，那么会返回`count`个`随机`集合。
* 若提供参数 `count` 大于或集合基数，那么返回整个集合。
* 若提供参数 `count` 为负数，返回数组长度为`count`的绝对值，数组元素可能重复

##### SCARD

* ` SCARD key`
* 返回集合key中元素的数量

##### SREM

* `SREM key member member member ...`
* 移除集合key中一个或多个member元素，不存在的member元素将会被忽略

##### SMOVE

* `SMOVE source destination member`
* 将`membe`r元素从`source`集合移动到`destination`集合
* 操作完成后`source`集合`member`就被搬走了（没有了）

##### SUNION

* `SUNION key key ...` 
* 返回一个集合的全部成员，该集合是所有给定集合的并集
* 并集就是将两个集合并在一起的集合。

##### SUNIONSTORE

* `SUNIONSTORE destination key key ... `
* 这个指令类似于`SUNION`，但它将结果保存到`destination`集合

#### SET的数据结构

* SET是通过哈希表实现的



### Hash 哈希

* Hash是`键值对集合`
* redis Hash是一个`string`类型的`field`和`value`的映射表，Hash特别适用于存储对象。

#### Redis Hash 操作

##### HSET 

* `HSET hash field value field value ...`
* 将哈希表`hash`中域 `field`的值设置为`value`

##### HSETNX

* `HSETNX hash field value `
* 仅设置不存在的`field`，否则设置失败
* 设置成功返回1，失败返回0

##### HGET

* `HGET key field`
* 获取指定`key`中的`field`关联的值

##### HEXISTS

* `HEXISTS key field`
* 检查给定域field是否存在

##### HDEL 

* `HDEL key field  field ...`
* 可以删除单个域、多个域、不存在的键
* 删除成功返回1，否则返回0 

##### HLEN

* `HLEN key field`
* 返回哈希表key中域的数量

##### HSTRLEN

* `HSTRLEN key field`
* 返回哈希表key中，field关联值的字符串长度

##### HINCRBY

* `HINCRBY key field increment`
* 为哈希表`key`中域给定域field的值加上增量`increment`
* 只能是`integer`整形，不然报错`(error) ERR hash value is not an integer`

##### HINCRBYFLOAT 

* `HINCRBYFLOAT key field increment`
* 为哈希表`key`中域`field`加上浮点数增量`increment`

##### HMSET

* `HMSET key field value field value ...`
* 在哈希表`key`中同时设置多个值

##### HMGET

* `HMGET key field field field ...`
* 在哈希表中`key`中同时获取多个值
* 若`field`不存在则返回`nil`

##### HEKYS

* `HKEYS key` 
* 返回哈希表`key`中所有的域

##### HVALS 

* `HVALS key`
* 返回哈希表`key`中所有域的值

##### HGETALL

* `HGETALL key`
* 返回哈希表`key`中所有的域和值

##### HSCAN	

* `HSCAN key cursor [MATCH pattern] [COUNT count]`
* 字段
  * cursor 游标
  * Pattern 匹配的模式
  * count  指定从数据库力返回多少元素，默认值是0 
* 返回值：返回的每个元素都是一个元组，每一个元组元素由一个字段(field) 和值（value）组成。

#### 哈希表的数据结构

* 哈希表对应的数据结构有两种，一种是`ziplist`（压缩列表），另一种是`hashtable`（哈希表）
* 当`field-value`长度较短且个数比较少时，使用`ziplist`，否则使用`hashtable`
* ziplist：当哈希类型元素个数小于`hash-max-ziplist-entries`配置（默认512个），同时所有值都小于`hash-max-ziplist-value`配置（默认64个字节）时，Redis会使用`ziplist`作为哈希的内部实现`ziplist`使用更加紧凑的结构实现多个元素的连续存储，所以在节省内存方面比`hashtable`更加优秀。
* hashtable：当哈希类型无法满足`ziplist`的条件时，Redis会使用`hashtable`作为哈希的内部实现。因为此时`ziplist`的读写效率会下降，而`hashtable`的读写时间复杂度为O(1)。



### List 列表

- Redis列表是简单的字符串列表，按照插入顺序排序。可以添加一个元素到列表的头部（左边）或尾部（右边）。

- list的底层实现是双向链表，对两端的操作性很强，但随机访问性能较弱。





#### Redis的List操作

##### LPUSH /RPUSH

- `LPUSH / RPUSH key value value ...`

* ```
  127.0.0.1:6379> LPUSH k1 one two three
  (integer) 3
  127.0.0.1:6379> LINDEX k1 0 
  "three"
  127.0.0.1:6379> LINDEX k1 1
  "two"
  127.0.0.1:6379> LINDEX k1 2
  "one"
  ```

* 将一个或多个值插入到队列key的表头/尾

##### LPUSHX / RPUSHX

- `LPUSHX / RPUSHX key value value ...`
- 仅当`key`存在时才能插入`value`

##### LPOP / RPOP

- `LPOP / RPOP key `
- 从 左边/右边 移除并返回列表key的头元素 / 尾元素
- 到最后如果元素都没有了，那么key也不存在了

##### RPOPLPUSH 

- `RPOPLPUSH source destination`
- 命令`RPOPLPUSH`在一个原子时间内执行两个动作
  1. 将队列source中的尾元素弹出并返回客户端
  2. 将`source`弹出的元素插入到列表`destination`，作为`destination`列表的头元素

##### LRANGE 

- `LRANGE key start stop`
- 按照索引获得元素（从左往右）
- 获取全部元素 `LRANGE key 0 -1`



##### LREM

- `LREM key count element`
- 根据参数`count`的值，移除列表中与参数`element`相等的元素
- `count`的类型
  - `count > 0` 从表头开始向表尾搜索，移除与value相等的元素，数量为`count`
  - `count < 0` 从表尾开始向表头搜索，移除与`value`相等的元素，数量为`count`的绝对值
  - `count = 0` 移除表中所有与`value`相等的值


##### LLEN

* `LLEN key`
* 返回列表key的长度

##### LINDEX

* `LINDEX key index`
* 返回列表key中，下标为index的元素

##### LINSERT

* `LINSERT key BEFORE|AFTER pivot value`
* 将值value插入到列表key当中，位于值`pivot`之前或之后
* 当`pivot`不存在时， 不进行任何操作
* 当`key`不存在时,不进行任何操作
* 当`key` 不是列表类型时，返回错误

##### LSET

* `LSET key index value`
* 将列表`key`下标为`index`元素的值设置为`value`

##### LRANGE

* `LRANGE key start stop`
* 返回列表`key`中指定区间的元素，区间以偏移量`start`和`stop`指定
* 下标参数`start`和`stop`都以0为底，既0表示列表的第一个元素，-1表示列表最后的元素,-2表示列表倒数第二位的元素。

##### LTRIM

* `LTRIM key start stop`
* 只保留列表中从`start`到`stop`下标的元素

##### BLPOP / BRPOP / BRPOPLPUSH

* `BLPOP / BRPOP key timeout`

* `BLPOP`指令会移出并获取列表的第一个元素，若没有元素则会阻塞到`timeout`时间（秒），或期间发现有可弹出元素为止。

* `BRPOP`指令会移出并获取列表的最后一个元素，若没有元素则会阻塞到`timeout`时间（秒），或期间发现有可弹出元素为止。

* `BRPOPLPUSH source destination timeout`

* `BRPOPLPUSH`会弹出最后一个元素，并将元素插入到另一个队列的头部；若没有元素则会阻塞到`timeout`时间（秒），或期间发现有可弹出元素为止。

* ```
  127.0.0.1:6379> blpop k1 5
  (nil)
  (5.01s)
  ```

  

### Zset 有序集合

* Zset与普通集合set非常相似，是一个没有重复元素的字符串集合。
* 不同之处是有序集合Zset的每个成员都关联了一个评分(score)，这个评分(score)被用来按照最低分到最高分的方式排序集合中的成员。集合中的成员是唯一的，但评分是可以重复的。
* 因为元素是有序的，所以能根据`评分`或`次序`来获取一个范围的元素。

#### Redis的Zset操作

##### ZADD

* `ZADD key score member`
* 将一个或多个`member`元素和`score`值加入到有序集`key`当中

##### ZSCORE

* `ZSCORE key member`
* 返回有序集`key`中，成员`member`的`score`值

##### ZINCRBY

* `zincrby key increment member`
* 给`zset`中`key`的成员`member`的`socre`加上`increment`

##### ZCARD

* `ZCARD key`
* 返回有序集合key的基数

##### ZCOUNT

* `ZCOUNT key min max`
* 返回集合中`score`的值介于min和max之间（含等于）的成员数

##### ZRANGE

* `ZRANGE key start stop`
* 返回指定区间的成员`member`
* 成员位置按照`socre`值从小到大排序

##### ZREVRANGE

* `ZREVRANGE key start stop`
* 返回指定区间的成员`member`
* 成员位置按照`socre`值从大到小排序

##### ZRANGEBYSOCRE

* `ZRANGEBYSCORE key min max`
* 返回有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。有序集成员按 score 值递增(从小到大)次序排列。
* 可选参数
  * WITHSCORES  返回带分数
  * LIMIT offset count 

##### ZREVRANGEBYSCORE

*  `ZREVRANGEBYSCORE key min max`
* 返回有序集 key 中， score 值介于 max 和 min 之间(默认包括等于 max 或 min )的所有的成员。有序集成员按 score 值递减(从大到小)的次序排列。

##### ZRANK

* `ZRANK key member`
* 返回有序集 key 中成员 member 的排名。其中有序集成员按 score 值递增(从小到大)顺序排列。

##### ZREVRANK

* `ZREVRANK key member`
* 返回有序集 key 中成员 member 的排名。其中有序集成员按 score 值递增(从大到小)顺序排列。

##### ZREM

* `ZREM key member` 
* 移除有序集合key中的一个或多个成员，不存的成员将直接被忽略

##### ZREMRANGEBYRANK

* `ZREMRANGEBYRANK key start stop`
* 移除有序集 key 中，指定排名(rank)区间内的所有成员

##### ZREMRANGEBYSCORE

* `ZREMRANGEBYSCORE key min max`
* 移除有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。

##### ZRANGEBYLEX

* `ZRANGEBYLEX key min max`
* 当有序集合的所有成员都具有相同的分值时， 有序集合的元素会根据成员的字典序（lexicographical ordering）来进行排序， 而这个命令则可以返回给定的有序集合键 key 中， 值介于 min 和 max 之间的成员。

##### ZLEXCOUNT

* `ZLEXCOUNT key min max`
* 对于一个所有成员的分值都相同的有序集合键 key 来说， 这个命令会返回该集合中， 成员介于 min 和 max 范围内的元素数量。

## Redis 配置文件
### Redis配置文件目录
#### 在ubuntu下


### Redis常见配置
- bind 
  - 默认情况下，如果没有指定bind，redis会监听主机的所有可用网络接口。通过配置bind可以监听一个或多个网络端口
  -  bind 192.168.1.100 10.0.0.1   `监听两个IPV4端口`
  -  bind 127.0.0.1 ::1  `监听本地回环地址IPV4和IPV6`
- port 
  - 配置监听端口
- protected-mode 
  - 默认开启
  - 只有当你确认其他主机的客户端连接到Redis(即使没有验证身份)才应该禁用
  - 配置Redis集群时需要将protected-mode设置为禁用
- tcp-backlpg
  - 用于设置linux系统中控制TCP三次握手`已完成连接队列`的长度
  - 在高并发系统中，需要设置一个较高的`tcp-backlog`来避免客户端访问速度慢的问题
- timeout
  - 设置客户端空闲多少秒后关闭连接
  - 设置0为禁用
- tcp-keepalive
  - 设置`tcp-keepalive`的主要用途是心跳检测
  - 这个选项设置合理值为300
- supervised 
  - 设置redis的进程守护
  - 选项
    - no  没有监督
    - upstart  通过将Redis置于SIGSTOP模式来启动信号
    - systemd signal systemd将READY = 1写入$ NOTIFY_SOCKET
    - auto  检测upstart或systemd方法基于 UPSTART_JOB或NOTIFY_SOCKET环境变量
- pidfile 
  - `/var/run/redis/redis-server.pid`
  - 如果指定了pid文件，Redis会在启动时将其写入指定的位置
并在退出时删除它
  - 如果在配置中没有指定pid文件，那么当服务器运行时，不会创建pid文件。当服务器被守护时，即使没有指定pid文件，也会使用pid文件，默认为"/var/run/redis.pid"
  - 创建一个pid文件是最好的努力:如果Redis不能创建它，没有什么不好的事情发生，服务器将正常启动和运行。
  - 注意，在现代Linux系统中，“/run/redis. conf”Pid”更符合要求，应该用它来代替。

- loglevel 
  - 指定redis服务级别
  - 选项
    - debug 大量信息，对开发/测试很有用
    - verbose 许多很少有用的信息，但不像调试级别那样混乱
    - notice 稍微有点啰嗦，可能是您在生产中想要的
    - warning 只记录非常重要/关键的消息
- databases 
  - 设置redis默认数据库数量
- maxclients
  - 设置最多客户端连接数
  - 默认是10000  


## Redis 订阅与发布



## Redis 新数据类型



## Go-Redis

## 事务与锁机制
### 乐观锁和悲观锁

## 持久化操作
### RDB

### AOF

## 主从复制

## 搭建Redis集群

## 异常处理
### 缓存击穿

### 缓存穿透

### 缓存预热

### 缓存降级

## 分布式缓存

## Redis ACL
