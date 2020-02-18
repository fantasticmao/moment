package app

import (
	"errors"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

const targetFileNameSuffix = "_final"

var (
	ErrFormatConsistency = errors.New("the format of each image should be consistency")
	ErrWidthConsistency  = errors.New("the width of each image should be consistency")
	ErrSubImage          = errors.New("image: unsupported method SubImage")
)

// 拼接图片
func Stitch(opts Options) string {
	// 校验命令行参数
	opts.checkArgs()

	// 读取文件
	files := opts.openFiles()
	defer opts.closeFiles(files)

	// 读取第一张图片
	firstFile := files[0]
	firstImg, firstImgFormat, err := image.Decode(firstFile)
	if err != nil {
		panic(err)
	}

	// 生成最终图片
	finalImageWidth := firstImg.Bounds().Max.X
	finalImageHeight := firstImg.Bounds().Max.Y + opts.Height*(len(opts.Paths)-1)
	finalImg := image.NewRGBA(image.Rect(0, 0, finalImageWidth, finalImageHeight))

	// 拼接第一张图片
	draw.Draw(finalImg, firstImg.Bounds(), firstImg, firstImg.Bounds().Min, draw.Over)

	// 读取并拼接剩余图片
	for index, file := range files[1:] {
		img, format, err := image.Decode(file)
		if err != nil {
			panic(err)
		}
		// 校验图片格式
		if format != firstImgFormat {
			panic(ErrFormatConsistency)
		}
		// 校验图片宽度
		if img.Bounds().Max.X != finalImageWidth {
			panic(ErrWidthConsistency)
		}
		imgImpl, ok := img.(interface {
			SubImage(r image.Rectangle) image.Image
		})
		if ok {
			// 裁剪图片字幕
			subImg := imgImpl.SubImage(image.Rect(0, img.Bounds().Max.Y-opts.Height, img.Bounds().Max.X, img.Bounds().Max.Y))
			// 拼接图片字幕
			draw.Draw(finalImg,
				image.Rect(0, firstImg.Bounds().Max.Y+index*opts.Height, finalImageWidth, firstImg.Bounds().Max.Y+(index+1)*opts.Height),
				subImg, subImg.Bounds().Min, draw.Over)
		} else {
			panic(ErrSubImage)
		}
	}

	// 生成最终文件
	targetName := getBaseNameWithoutExtension(firstFile) + targetFileNameSuffix + getExtension(firstFile)
	targetPath := filepath.Join(opts.Out, targetName)
	targetFile, err := os.Create(targetPath)
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()

	// 保存最终文件
	encodeByFormat(targetFile, finalImg, firstImgFormat)
	return targetPath
}

func encodeByFormat(file *os.File, img image.Image, format string) {
	var err error = nil
	switch format {
	case "png":
		err = png.Encode(file, img)
	case "jpeg":
		err = jpeg.Encode(file, img, &jpeg.Options{Quality: 90})
	default:
		err = image.ErrFormat
	}

	if err != nil {
		panic(err)
	}
}
