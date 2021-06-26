# Staufen JavaScript SDK

## SDK 上传流程
- 自定义 UploadToken 获取逻辑
    - 设置为空的情况下可以直接上传？
    - UploadToken 有效期设定，一个 UploadToken 可能会有2个小时的有效期
    - 单 UploadToken 失效了可以自动重新获取
- 自定义上传回调
    - 返回请求 ID
    - 返回图片格式
        - 小尺寸 small
        - 中间尺寸 middle
    - 原图 large
- 自定义失败回调