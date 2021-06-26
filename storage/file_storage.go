package storage

import (
	"github.com/Ericwyn/GoTools/date"
	"github.com/Ericwyn/Staufen/util/gen"
	"github.com/Ericwyn/Staufen/util/log"
	"github.com/Ericwyn/Staufen/util/strutil"
	"github.com/spf13/viper"
	"io"
	"time"

	//"github.com/Ericwyn/GoTools/file"
	"mime/multipart"
	"os"
)

// 保存图片，返回一个文件地址
func SavePic(picFile *multipart.FileHeader) (string, error) {

	fileSaveName := gen.GeneralRandomStr(10) + strutil.GetExtName(picFile.Filename)
	savePath := getSaveDirPath() + "/" + fileSaveName

	out, err := os.Create(savePath)
	if err != nil {
		return "", err
	}

	defer out.Close()

	open, err := picFile.Open()
	if err != nil {
		return "", err
	}

	defer open.Close()

	_, err = io.Copy(out, open)

	if err != nil {
		return "", err
	} else {
		return savePath, nil
	}
}

func GetPic(picPath string) *os.File {
	fileOpen, err := os.Open(picPath)
	if err != nil {
		log.E("open file error :"+picPath, " --> ", err)
		return nil
	} else {
		return fileOpen
	}
	//return nil
}

// 基于上传时间节点，获取年/月/日
func getSaveDirPath() string {
	picDataPath := viper.GetString("path.picData")

	if picDataPath == "" {
		picDataPath = "./picData"
	}
	dirPath := picDataPath + date.Format(time.Now(), "/yy/MM/dd")
	if !isExist(dirPath) {
		_ = os.MkdirAll(dirPath, os.ModePerm)
	}
	return dirPath
}

func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
