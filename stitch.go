package main

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

const finalImgNameSuffix = "_final"

type subImage interface {
	SubImage(r image.Rectangle) image.Image
}

func stitch(images []string, height int) (string, error) {
	files, err := openFiles(images)
	if err != nil {
		return "", err
	}
	defer closeFiles(files)

	// 读取第一张图片
	baseFile := files[0]
	baseImg, baseImgWidth, baseImgHeight, baseImgFormat, err := decodeImage(baseFile)
	if err != nil {
		return "", err
	}

	// 生成最终图片
	finalImgHeight := baseImg.Bounds().Max.Y + height*(len(images)-1)
	finalImg := image.NewRGBA(image.Rect(0, 0, baseImgWidth, finalImgHeight))

	// 拼接第一张图片
	draw.Draw(finalImg, baseImg.Bounds(), baseImg, baseImg.Bounds().Min, draw.Over)

	// 读取并拼接剩余图片
	for index, file := range files[1:] {
		img, imgWidth, imgHeight, imgFormat, err := decodeImage(file)
		if err != nil {
			return "", err
		}

		// 校验图片格式
		if imgFormat != baseImgFormat {
			return "", errors.New(fmt.Sprintf("image formats between %s(%s) and %s(%s) are not the same",
				imgFormat, file.Name(), baseImgFormat, baseFile.Name()))
		}

		// 校验图片宽度
		if imgWidth != baseImgWidth {
			return "", errors.New(fmt.Sprintf("image width between %d(%s) and %d(%s) are not the same",
				imgWidth, file.Name(), baseImgWidth, baseFile.Name()))
		}

		if imgImpl, ok := img.(subImage); ok {
			// 裁剪图片字幕
			nextImg := imgImpl.SubImage(image.Rect(0, imgHeight-height, imgWidth, imgHeight))
			// 拼接图片字幕
			nextImgRect := image.Rect(0, baseImgHeight+height*index,
				baseImgWidth, baseImgHeight+height*(index+1))
			draw.Draw(finalImg, nextImgRect, nextImg, nextImg.Bounds().Min, draw.Over)
		} else {
			return "", errors.New("image: unsupported method SubImage")
		}
	}

	// 生成最终文件
	targetName := getFileName(baseFile) + finalImgNameSuffix + getFileExtension(baseFile)
	targetPath := filepath.Join(".", targetName)
	targetFile, err := os.Create(targetPath)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	// 保存最终文件
	err = encodeByFormat(targetFile, finalImg, baseImgFormat)
	if err != nil {
		return "", err
	}
	return targetPath, nil
}

func decodeImage(file *os.File) (image.Image, int, int, string, error) {
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, 0, 0, "", err
	}
	return img, img.Bounds().Max.X, img.Bounds().Max.Y, format, nil
}

func encodeByFormat(file *os.File, img image.Image, format string) error {
	var err error = nil
	switch format {
	case "png":
		err = png.Encode(file, img)
	case "jpeg":
		err = jpeg.Encode(file, img, &jpeg.Options{Quality: 90})
	default:
		err = image.ErrFormat
	}
	return err
}
