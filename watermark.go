package watermark

import (
	"errors"
	"github.com/kakuilan/kgo"
)

var ErrFontEmpty = errors.New("font content is empty")

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
