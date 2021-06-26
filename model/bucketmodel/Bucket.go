package bucketmodel

import "time"

type Bucket struct {
	Id         int64
	Name       string
	ApiToken   string
	reqToken   string
	public     bool
	createTime time.Time
	compress   bool
}
