# 数据库设计

## Bucket 表
每一个 Bucket 类似于一个存储桶

- ID
  
- name
  - 名称
- apiToken
  - 调用 api 接口相关
- reqToken
  - 请求的 token ？ 
  - 如果这个 bucket并非 public 的话，需要带上 reqToken 参数才可以访问
- public
  - 公开可访问?
- createTime
  - 创建日期
- compression
  - 是否启用压缩？
    如果启用的话，存储这个 bucket 里面的图片都会存储一份压缩版本
    比如自动将图片压缩到 300 边长？


## Picture 表
存储图片信息

- id
  - 图片的 id
- filename
  - 原始文件图片名称
- path
  - 本地存储的路径
- createTime
  - 创建的时间戳
- bucket
  - bucketId