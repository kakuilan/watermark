package watermark

import (
	"errors"
	"github.com/kakuilan/kgo"
)

var ErrFontEmpty = errors.New("font content is empty")
var ErrDirectoryEmpty = errors.New("directory is empty")
var ErrDirectoryNotExist = errors.New("directory is not exist")
var ErrDirectoryNotWritable = errors.New("directory is not writable")

func NewFontFromPath(p string) (*FontInfo, error) {
	f := new(FontInfo)
	err := f.SetFontFromPath(p)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func NewFontFromByte(bs []byte) (*FontInfo, error) {
	f := new(FontInfo)
	err := f.SetFontFromByte(bs)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// SetFontFromPath 根据路径设置字体.
func (f *FontInfo) SetFontFromPath(p string) error {
	bs, err := kgo.KFile.ReadFile(p)
	if err != nil {
		return err
	}

	f.fontPath = p
	f.fontBytes = bs

	return nil
}

// SetFontFromByte 根据字节设置字体.
func (f *FontInfo) SetFontFromByte(bs []byte) error {
	if len(bs) == 0 {
		return ErrFontEmpty
	}

	f.fontPath = ""
	f.fontBytes = bs

	return nil
}

// ReadFont 读取字体.
func (f *FontInfo) ReadFont() ([]byte, error) {
	if len(f.fontBytes) == 0 {
		return nil, ErrFontEmpty
	}

	return f.fontBytes, nil
}

// NewWatermark 新建水印对象.
func NewWatermark(opt Option) *Watermark {
	res := new(Watermark)

	return res
}

// SetOption 设置选项.
func (w *Watermark) SetOption(opt Option) *Watermark {
	w.opt = opt
	return w
}

// SetTempDir 设置临时目录.
func (w *Watermark) SetTempDir(dir string) error {
	if dir == "" {
		return ErrDirectoryEmpty
	} else if !kgo.KFile.IsDir(dir) {
		return ErrDirectoryNotExist
	} else if !kgo.KFile.IsWritable(dir) {
		return ErrDirectoryNotWritable
	}

	w.tempDir = dir

	return nil
}

// SetOutputDir 设置输出目录.
func (w *Watermark) SetOutputDir(dir string) error {
	if dir == "" {
		return ErrDirectoryEmpty
	} else if !kgo.KFile.IsDir(dir) {
		return ErrDirectoryNotExist
	} else if !kgo.KFile.IsWritable(dir) {
		return ErrDirectoryNotWritable
	}

	w.outputDir = dir

	return nil
}
