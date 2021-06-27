package storage

import (
	"bytes"
	"github.com/Ericwyn/GoTools/date"
	"github.com/Ericwyn/Staufen/util/compress"
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

type FileQuality int

var mini FileQuality = 1
var middle FileQuality = 1

var MiniQuality *FileQuality = &mini
var MiddleQuality *FileQuality = &middle

// 保存图片，返回一个文件地址
func SavePic(picFile *multipart.FileHeader, compressFlag bool, saveType SaveType) (string, error) {
	randomId := gen.GeneralRandomStr(10)
	fileSaveName := randomId + "." + strutil.GetExtName(picFile.Filename)
	savePath := getSaveDirPath() + "/" + fileSaveName

	out, err := os.Create(savePath)
	if err != nil {
		return "", err
	}

	open, err := picFile.Open()
	if err != nil {
		return "", err
	}
	_, err = io.Copy(out, open)

	_ = out.Close()
	_ = open.Close()

	// 如果压缩的话需要再保存两遍，一个 mini 尺寸(高300), 一个 middle 尺寸(高 1000 )
	if compressFlag {
		// 获取原始图片的 bytes

		compressOriginFile, err := picFile.Open()

		buf := bytes.NewBuffer(nil)
		_, _ = io.Copy(buf, compressOriginFile)

		originalBytes := buf.Bytes()

		if err != nil {
			log.E(err)
		}
		miniCompressBytes, _ := compress.CompressImg(originalBytes, picFile.Filename, compress.QualityMini)
		if miniCompressBytes != nil {
			//fileSaveName := randomId + ".mini." + strutil.GetExtName(picFile.Filename)
			//savePath := getSaveDirPath() + "/" + fileSaveName

			_ = ioutil.WriteFile(savePath+".mini", miniCompressBytes, 0777)
		}

		middleCompressBytes, _ := compress.CompressImg(originalBytes, picFile.Filename, compress.QualityMiddle)
		if middleCompressBytes != nil {
			//fileSaveName := randomId + ".middle." + strutil.GetExtName(picFile.Filename)
			//savePath := getSaveDirPath() + "/" + fileSaveName

			_ = ioutil.WriteFile(savePath+".middle", middleCompressBytes, 0777)
		}

		_ = compressOriginFile.Close()
	}

	if err != nil {
		return "", err
	} else {
		return savePath, nil
	}
}

// 直接返回 IO 流程
func GetPicBytes(picPath string, quality *FileQuality, saveType SaveType) []byte {
	if quality == MiniQuality {
		picPath = picPath + ".mini"
	}
	if quality == MiddleQuality {
		picPath = picPath + ".middle"
	}
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

func DeletePics(picPaths []string, saveType SaveType) {
	for _, picPath := range picPaths {
		DeletePic(picPath, saveType)
	}
}

func DeletePic(picPath string, saveType SaveType) error {
	//open, err := os.Open(picPath)
	return os.Remove(picPath)
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
