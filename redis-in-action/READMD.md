## 《Redis实战》笔记
这个文件记录的是《redis实战》的学习笔记 
- 图书中的github地址：https://github.com/josiahcarlson/redis-in-action/
- Go代码地址：https://github.com/josiahcarlson/redis-in-action/tree/master/golang

## 做一个对文章投票的网站demo
- 要构建一个文章投票网站，我们首先要做的就是为了这个网站设置一些数值和限制条件：
1. 如果一篇文章获得了至少200张支持票（up vote），那么网站就认为这篇文章是一篇有趣的文章；
2. 假如这个网站每天发布1000篇文章，而其中的50篇符合网站对有趣文章的要求，那么网站要做的就是把这50篇文章放到文章列表前100位至少一天；
3. 另外，这个网站暂时不提供投反对票（down vote）的功能。