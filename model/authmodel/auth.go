package authmodel

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/Ericwyn/Staufen/repo"
	"github.com/Ericwyn/Staufen/util/strutil"
	"strings"
	"time"
)

func GetUploadToken(bucket repo.Bucket) string {
	//sha1Encode()
	unix := time.Now().Unix()
	return calUploadToken(bucket.ApiToken, bucket.Salt, unix)
}

/*
bucket 创建时候会生成 API Token + salt

uploadToken 通过 md5(API Token + salt + yyyy/MM/dd/hh/mm/ss) 组成

uploadToken = secret.时间点(unix时间戳)
	secret = sha1(api token + salt + yyyy/MM/dd).取前10 - base64(时间戳)

*/
func calUploadToken(bucketApiToken string, salt string, unixTime int64) string {
	secret := sha1Encode(bucketApiToken + salt + fmt.Sprint(unixTime))
	timestamp := base64Encode(fmt.Sprint(unixTime))
	return secret + "-" + timestamp
}

func CheckUploadToken(bucket repo.Bucket, uploadToken string) bool {
	// 验证格式
	split := strings.Split(uploadToken, "-")
	if len(split) != 2 {
		return false
	}

	// 验证日期
	unixStr := base64Decode(split[1])
	if unixStr == "" {
		return false
	}

	unixTime := strutil.ToInt64(unixStr)

	s := time.Now().Unix() - unixTime

	// 过期时间是 3600s 也就是一个小时，SDK 的话过期时间设置成 3000 左右 (50分钟就过期)
	if unixTime != 0 && 0 < s && s < 3600 {
		return sha1Encode(bucket.ApiToken+bucket.Salt+fmt.Sprint(unixTime)) == split[0]
	}

	// TODO UploadToken 校验
	return false
}

func sha1Encode(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// b64 加密
func base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// b64 解密
func base64Decode(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(decoded)
}
