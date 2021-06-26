package picmodel

import "time"

type Picture struct {
	Id         int64  // 自增方便顺序查找
	Uuid       string // 32 位 id
	FileName   string // 原始文件名
	Path       string // 存储的数据路径
	CreateTime time.Time
	BucketId   string
}
