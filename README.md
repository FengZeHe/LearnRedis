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


## Redis五种数据类型
### String 字符串


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
