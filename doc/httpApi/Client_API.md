# Client API 文档设计


## 图片数据库

## 鉴权
- 开放 API TOKEN
    - 用户鉴定
    
- 使用 API Token 换取短期 token
    - 短期 token 由前端使用，前端可用该可直接上传
    
## 图片上传
### 图片上传
- POST /api/pic/upload
- 参数
    - File
    - Upload ?
    - token （必须）
    
- 返回
    - 图床 id


### 图片更新
- POST /api/pic/upload
- 参数
    - File
    - Upload ?
    - token （必须）


## 图片获取
- GET /api/pic/get/xxxxxxx.jpg
- 参数
    - xxxxx 
      - 指代 id 名称
    - type
      - 缩略图尺寸
      - mini： 高 300 的缩略图
      - middle： 高 1000 的缩略图
      - 只有在启用了压缩并且图片尺寸原本就大于这个的情况下才有效
      - 否则无效，会直接返回原图
    - reqToken
      - 如果是私有的 bucket，需要带上这个参数才可以获取，否则返回无效
    

## 图片删除
### 删除图片