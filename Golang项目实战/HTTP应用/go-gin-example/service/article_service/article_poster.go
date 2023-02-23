package article_service

import (
	"github.com/EDDYCJY/go-gin-example/pkg/file"
	"github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

type ArticlePoster struct {
	PosterName string
	*Article
	Qr *qrcode.QrCode
}

func NewArticlePoster(posterName string, article *Article, qr *qrcode.QrCode) *ArticlePoster {
	return &ArticlePoster{
		PosterName: posterName,
		Article:    article,
		Qr:         qr,
	}
}

func GetPosterFlag() string {
	return "poster"
}

// 如果图片文件不存在返回false，存在则返回true
func (a *ArticlePoster) CheckMergedImageFile(path string) bool {
	if file.CheckNotExist(path+a.PosterName) == true {
		return false
	}
	return true
}

// 打开图片文件
func (a *ArticlePoster) OpenMergedImage(path string) (*os.File, error) {
	f, err := file.MustOpen(a.PosterName, path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

type ArticlePosterBg struct {
	Name string
	*ArticlePoster
	*Rect
	*Pt
}

type Rect struct {
	Name string
	X0   int
	Y0   int
	X1   int
	Y1   int
}

type Pt struct {
	X int
	Y int
}

func NewArticlePosterBg(name string, ap *ArticlePoster, rect *Rect, pt *Pt) *ArticlePosterBg {
	return &ArticlePosterBg{
		Name:          name,
		ArticlePoster: ap,
		Rect:          rect,
		Pt:            pt,
	}
}

func (a *ArticlePosterBg) Generate() (string, string, error) {
	fullPath := qrcode.GetQrCodeFullPath()
	fileName, path, err := a.Qr.Encode(fullPath) // 在指定路径生成二维码图片。如果存在则直接返回文件名和路径
	if err != nil {
		return "", "", err
	}

	if !a.CheckMergedImageFile(path) { // 被合并后的图片文件不存在
		mergedF, err := a.OpenMergedImage(path) // 创建合并图片文件
		if err != nil {
			return "", "", err
		}
		defer mergedF.Close()

		bgF, err := file.MustOpen(a.Name, path) // 打开背景图片文件
		if err != nil {
			return "", "", err
		}
		defer bgF.Close()

		qrF, err := file.MustOpen(fileName, path) // 打开二维码图片文件
		if err != nil {
			return "", "", err
		}
		defer qrF.Close()

		bgImage, err := jpeg.Decode(bgF) // 读取背景图jpeg文件生成image结构体
		if err != nil {
			return "", "", err
		}
		qrImage, err := jpeg.Decode(qrF) // 读取二维码jpeg文件生成image结构体
		if err != nil {
			return "", "", err
		}

		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1)) // 返回具有给定边界的新RGBA图像

		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)                               // 添加背景图
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over) // 添加二维码

		jpeg.Encode(mergedF, jpg, nil) // 重新生成JPEG图片(合并后的)
	}

	return fileName, path, nil
}
