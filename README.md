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


### Set 集合

### Hash 哈希

### List 列表

### Zset 有序集合

## Redis 新数据类型

## Redis 配置文件

## Redis 订阅与发布

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
