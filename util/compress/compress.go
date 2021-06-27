package compress

import (
	"bytes"
	"fmt"
	"github.com/Ericwyn/Staufen/util/log"
	"github.com/Ericwyn/Staufen/util/strutil"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

type Quality string

const QualityMini Quality = "mini"
const QualityMiddle Quality = "middle"

const qualityMiniHeight uint = 300
const qualityMiddleHeight uint = 1000

// 压缩一张图片，返回压缩完的 bytes
func CompressImg(bts []byte, fileName string, quality Quality) ([]byte, error) {
	var err error

	extName := strutil.GetExtName(fileName)
	extName = strings.ToLower(extName)

	var height uint = qualityMiniHeight
	if quality == QualityMiddle {
		height = qualityMiddleHeight
	}

	//name := strings.ToLower(file.Name())
	var img image.Image
	var config image.Config

	if extName == "png" {
		// 过滤尺寸，如果原图尺寸已经过小，就没有必要压缩了
		if config, err = png.DecodeConfig(bytes.NewReader(bts)); err != nil {
			log.E(err)
			return nil, err
		}
		if config.Height <= int(height) {
			return nil, err
		}

		if img, err = png.Decode(bytes.NewReader(bts)); err != nil {
			log.E(err)
			return nil, err
		}

		resizeImg := resize.Resize(0, height, img, resize.Lanczos3)

		// create buffer
		buff := new(bytes.Buffer)

		// encode image to buffer
		err = png.Encode(buff, resizeImg)
		if err != nil {
			fmt.Println("failed to create buffer", err)
		}

		// convert buffer to reader
		return buff.Bytes(), nil

	} else if extName == "jpg" || extName == "jpeg" {
		// 过滤尺寸，如果原图尺寸已经过小，就没有必要压缩了
		if config, err = jpeg.DecodeConfig(bytes.NewReader(bts)); err != nil {
			log.E(err)
			return nil, err
		}
		if config.Height <= int(height) {
			return nil, err
		}

		if img, err = jpeg.Decode(bytes.NewReader(bts)); err != nil {
			log.E(err)
			return nil, err
		}

		resizeImg := resize.Resize(height, 0, img, resize.Lanczos3)

		// create buffer
		buff := new(bytes.Buffer)

		// encode image to buffer
		err = jpeg.Encode(buff, resizeImg, nil)
		if err != nil {
			fmt.Println("failed to create buffer", err)
		}

		// convert buffer to reader
		return buff.Bytes(), nil
	} else {
		err = fmt.Errorf("can not compress file: %s", fileName)
		log.E(err)
		panic(err)
		return nil, err

	}
}
