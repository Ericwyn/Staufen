# Staufen

一个图床程序


## 基本设计
- 提供基础图床功能
- 一个图床，多个服务可用
- 分布式？（TODOOOOOOO）

## 模块
- cmd 命令行
    - 以 client 启动
    - 以 中心 server 启动
    

- io
    - 数据存储模块
        - 输入一个文件
        - 返回一个地址
    - 支持多种底层文件存储
        - LocalFile
        - COS(TODO)
        - HDFS(TODO)

- server 模块
    - picServer
        - 图片增删改查
        - 前端/上传页面
        - 前端，超管，图片查看和管理页面
        
    - staufenServer
        - staufen client 管理？发现？
    
## SDK
- Java SDK
    - 引入一个依赖
    - 注解一个 api 方法作为图片上传方法
    - 注解一个 api 方法作为图片下载，自动返回 301 地址

- JS SDK
