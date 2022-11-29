package watermark

import "image/color"

type FontInfo struct {
	fontPath  string //字体路径
	fontBytes []byte //字体文件内容
}

// 字体接口
type Font interface {
	SetFontFromPath(path string) error
	SetFontFromByte(bs []byte) error
	ReadFont() ([]byte, error)
}

type Option struct {
	MarkImage string     //水印图片(路径或URL)
	MarkText  string     //水印文本
	MaxHeight int        //水印图-最大高度
	MaxWidth  int        //水印图-最大宽度
	Top       int        //起点水印距背景图顶部(Y轴)距离
	Left      int        //起点水印距背景图左边(X轴)距离
	Opacity   float32    //水印透明度
	Replicate bool       //水印是否重复
	Rotate    int        //旋转角度(0~360)
	Font      Font       //字体
	FontColor color.RGBA //字体颜色
	FontSize  int        //字体大小
	DIP       float64    //扫描精度,默认56
}

type Watermark struct {
	tempDir   string //临时目录
	outputDir string //输出目录
	opt       Option //选项
}
