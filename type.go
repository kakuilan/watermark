package watermark

import "image/color"

type Font struct {
}

type Option struct {
	UseText    bool       //是否文字水印
	WorkDir    string     //工作目录
	MaxHeight  int        //水印图-最大高度
	MaxWidth   int        //水印图-最大宽度
	Top        int        //起点水印距背景图顶部(Y轴)距离
	Left       int        //起点水印距背景图左边(X轴)距离
	DIP        float64    //扫描精度,默认56
	Opacity    float32    //透明度
	Replicate  bool       //水印是否重复
	Rotate     int        //旋转角度(0~360)
	Font       *Font      //字体
	FontColor  color.RGBA //字体颜色
	FontSize   int        //字体大小
	MarkText   string     //水印文本
	MarkImage  string     //水印图片
	OutputPath string     //输出路径
}

type Watermark struct {
}
