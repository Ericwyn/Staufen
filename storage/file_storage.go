package storage

import (
	"github.com/Ericwyn/GoTools/date"
	"github.com/Ericwyn/Staufen/util/gen"
	"github.com/Ericwyn/Staufen/util/log"
	"github.com/Ericwyn/Staufen/util/strutil"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"time"

	//"github.com/Ericwyn/GoTools/file"
	"mime/multipart"
	"os"
)

type SaveType string

const LocalFile SaveType = "LocalFile" // 本地存储
// TODO 支持 HDFS，COS

// 保存图片，返回一个文件地址
func SavePic(picFile *multipart.FileHeader, saveType SaveType) (string, error) {

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

//func GetPic(picPath string) *os.File {
//	fileOpen, err := os.Open(picPath)
//	if err != nil {
//		log.E("open file error :"+picPath, " --> ", err)
//		return nil
//	} else {
//		return fileOpen
//	}
//	//return nil
//}

//// 直接返回 IO 流程
//func GetPicIO(picPath string, saveType SaveType) *io.Reader {
//	fileOpen, err := os.Open(picPath)
//	if err != nil {
//		log.E("open file error :"+picPath, " --> ", err)
//		return nil
//	} else {
//		//reader := bufio.NewReader(fileOpen)
//		//return reader.
//		//defer fileOpen.Close()
//		var read io.Reader = fileOpen
//		return &read
//	}
//}

// 直接返回 IO 流程
func GetPicBytes(picPath string, saveType SaveType) []byte {
	fileOpen, err := os.Open(picPath)
	if err != nil {
		log.E("open file error :"+picPath, " --> ", err)
		return nil
	} else {
		//reader := bufio.NewReader(fileOpen)
		//return reader.
		defer fileOpen.Close()
		//var read io.Reader = fileOpen
		//return &read
		bytes, err := ioutil.ReadAll(fileOpen)
		if err != nil {
			log.E("read pic bytes fail:", err)
			return nil
		}
		return bytes
	}
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
