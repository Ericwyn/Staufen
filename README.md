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
    
- storage
    - 数据存储模块
        - 输入一个文件
        - 返回一个地址
    - 支持多种底层文件存储
        - LocalFile
        - COS(TODOOOOOOO)
        - HDFS(TODOOOOOOO)

- server 模块
    - picServer
        - 图片增删改查
        - 前端/上传页面
        - 前端，超管，图片查看和管理页面
        
    - staufenServer
        - staufen client 管理？发现？
    

- compress
    - 图片压缩

## 配置文件
```yaml
db:  # mysql 数据库连接信息
  address: tcp(localhost:3306)/staufen?charset=utf8
  account: staufen
  password: staufen
  showSql: true
server:  # pic server http 服务器配置
  port: 9000
path:   # 本地图片文件存储位置
  picData: ./pic_data
```

## 启动参数
 - `-cli`
    
    命令行模式启动，可管理当前 bucket，直接获取 uploadToken
    
    ```
    欢迎来到 Staufen Cli

    ======================================================
    0 . list all bucket
    1 . create bucket
    2 . update bucket
    3 . delete bucket
    4 . get upload token
    ======================================================
    >>:
   ```

 - `-ps`

    picture server, 启动图片存储服务器

## SDK
- Java SDK
    - 引入一个依赖
    - 注解一个 api 方法作为图片上传方法
    - 注解一个 api 方法作为图片下载，自动返回 301 地址

- JS SDK
