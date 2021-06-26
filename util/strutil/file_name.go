package strutil

import (
	"github.com/Ericwyn/Staufen/util/log"
	"strconv"
	"strings"
)

// 获取文件后缀名
func GetExtName(fileName string) string {
	if strings.Index(fileName, ".") >= 0 {
		split := strings.Split(fileName, ".")
		return split[len(split)-1]
	}
	return ""
}

func ToInt64(str string) int64 {
	parseInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.E("parse str to int64 error:", err)
		return 0
	} else {
		return parseInt
	}
}

// parse int64 to string
func PInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}
