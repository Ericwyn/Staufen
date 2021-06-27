package repo

import "time"

type Bucket struct {
	Id         int64
	BucketId   string    `xorm:"varchar(32) notnull unique 'bucket_id'"`
	Name       string    `xorm:"varchar(200) notnull 'name' comment('bucket名称')"`
	ApiToken   string    `xorm:"varchar(64) notnull unique 'api_token'"`
	ReqToken   string    `xorm:"varchar(64) notnull 'req_token'"`
	Salt       string    `xorm:"varchar(64) notnull 'salt' comment('token计算salt')"`
	Public     bool      // 公共
	Compress   bool      // 存储到 bucket 里面的时候是否压缩图片
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

type Picture struct {
	Id         int64     // 自增方便顺序查找
	Uuid       string    `xorm:"varchar(32) notnull 'uuid' comment('uuid')"`        // 32 位 id
	FileName   string    `xorm:"varchar(200) notnull 'file_name' comment('原始文件名')"` // 原始文件名
	ExtName    string    `xorm:"varchar(10) notnull 'ext_name' comment('拓展名')"`     // 原始文件名
	FilePath   string    `xorm:"varchar(200) notnull 'file_path' comment('存储路径')"`  // 存储的数据路径
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
	BucketId   int64     `xorm:"BIGINT comment('bucket 外键')"`
	DeletedAt  time.Time `xorm:"deleted"` // 软删除
}
