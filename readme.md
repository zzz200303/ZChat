<h1>ZChat</h1>
基于Go-Zero的ZChat即时聊天系统<br>
本项目用于学习Go-Zero<br>
采用微服务架构，支持单/多聊，支持离线消息<br>
使用APISIX作为网关<br>
使用Kafka消息队列，作用：<br>
1. 新用户注册的时候通知聊天系统初始化新用户<br>
2. 削峰，聊天消息先送到Kafka再进行处理<br>
3. 未来可以引入go-stash和elesticsearch作为日志存储和聊天记录搜索<br>
使用Redis作为缓存，使用MySQL作为数据库<br>
作者：王品泽
