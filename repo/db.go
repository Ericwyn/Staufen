package repo

import (
	"fmt"
	"github.com/Ericwyn/Staufen/util/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"time"
	"xorm.io/xorm"
)

var SqlEngine *xorm.Engine

type Bucket struct {
	Id         int64
	BucketId   string    `xorm:"varchar(32) notnull unique 'bucket_id'"`
	Name       string    `xorm:"varchar(200) notnull 'name' comment('bucket名称')"`
	ApiToken   string    `xorm:"varchar(64) notnull unique 'api_token'"`
	ReqToken   string    `xorm:"varchar(64) notnull unique 'req_token'"`
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
}

//var sqlBuilder = builder.SQLite()

func InitDb() {
	var err error
	dataSourceName := viper.GetString("db.account") + ":" + viper.GetString("db.password") + "@" + viper.GetString("db.address")
	SqlEngine, err = xorm.NewEngine(
		"mysql",
		dataSourceName,
	)

	if err != nil {
		log.E(err)
		log.E("\n\n SQL ENGINE INIT FAIL!!")
		os.Exit(-1)
	}

	// 开启 SQL 打印

	SqlEngine.ShowSQL(viper.GetBool("db.showSql"))

	// 同步表结构
	err = SqlEngine.Sync2(
		new(Picture),
		new(Bucket),
	)
	if err != nil {
		fmt.Println(err)
		log.E("SYNC TABLE ERROR!!")
		os.Exit(-1)
	}
}
