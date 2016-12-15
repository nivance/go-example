package main

import (
	//	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	//	"io/ioutil"
	//	"log"
	//	"math/rand"
	"os"
	//	"runtime"
	//	"strconv"
	//	"strings"
	//	"time"

	//	"github.com/nfnt/resize"
)

func init2() {
	var original_img string = "C:\\Users\\lenovo\\Desktop\\huiben\\IMG_7394.jpg"
	var wm_img string = "C:\\Users\\lenovo\\Desktop\\huiben\\IMG_WM.jpg"
	var wm_logo string = "C:\\Users\\lenovo\\Desktop\\huiben\\logo.png"
	cmd_watermark(original_img, wm_img, wm_logo)
}

func cmd_watermark(file string, to string, logo string) {
	// 打开图片并解码
	file_origin, _ := os.Open(file)
	origin, _ := jpeg.Decode(file_origin)
	defer file_origin.Close()

	// 打开水印图并解码
	file_watermark, _ := os.Open(logo)
	watermark, _ := png.Decode(file_watermark)
	defer file_watermark.Close()

	offset := image.Pt(origin.Bounds().Dx()-watermark.Bounds().Dx()-10, origin.Bounds().Dy()-watermark.Bounds().Dy()-10)

	//原始图界限
	origin_size := origin.Bounds()

	//创建新图层
	canvas := image.NewNRGBA(origin_size)
	// 贴原始图
	draw.Draw(canvas, origin_size, origin, image.ZP, draw.Src)
	// 贴水印图
	draw.Draw(canvas, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	//生成新图片
	create_image, _ := os.Create(to)
	jpeg.Encode(create_image, canvas, &jpeg.Options{85})
	defer create_image.Close()
}
