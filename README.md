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

### 编译安装

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
