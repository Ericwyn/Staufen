package repo

import (
	"fmt"
	"github.com/Ericwyn/Staufen/util/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"xorm.io/xorm"
)

var SqlEngine *xorm.Engine

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
