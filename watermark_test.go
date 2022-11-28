package watermark

import (
	"github.com/kakuilan/kgo"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_NewFontFromPath(t *testing.T) {
	var res *FontInfo
	var err error
	var bs []byte

	Convey("读取路径字体", t, func() {
		//字体存在
		res, err = NewFontFromPath(fontWqy)
		So(err, ShouldBeNil)
		So(res, ShouldNotBeNil)
		Convey("字体内容存在", func() {
			bs, err = res.ReadFont()
			So(err, ShouldBeNil)
			So(bs, ShouldNotBeNil)
		})

		//字体不存在
		res, err = NewFontFromPath(fontNotExist)
		So(err, ShouldNotBeNil)
		So(res, ShouldBeNil)
		Convey("字体内容不存在", func() {
			res = new(FontInfo)
			bs, err = res.ReadFont()
			So(bs, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err, ShouldEqual, ErrFontEmpty)
		})
	})
}

func Benchmark_NewFontFromPath(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewFontFromPath(fontWqy)
	}
}

func Test_NewFontFromByte(t *testing.T) {
	var res *FontInfo
	var err error
	var bs []byte

	Convey("读取字节字体", t, func() {
		//字体存在
		bs, _ = kgo.KFile.ReadFile(fontWqy)
		res, err = NewFontFromByte(bs)
		So(err, ShouldBeNil)
		So(res, ShouldNotBeNil)
		Convey("字体内容存在", func() {
			bs, err = res.ReadFont()
			So(err, ShouldBeNil)
			So(bs, ShouldNotBeNil)
		})

		//字体不存在
		bs = []byte{}
		res, err = NewFontFromByte(bs)
		So(err, ShouldNotBeNil)
		So(res, ShouldBeNil)
		Convey("字体内容不存在", func() {
			res = new(FontInfo)
			bs, err = res.ReadFont()
			So(bs, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err, ShouldEqual, ErrFontEmpty)
		})
	})
}

func Benchmark_NewFontFromByte(b *testing.B) {
	b.ResetTimer()
	bs, _ := kgo.KFile.ReadFile(fontWqy)
	for i := 0; i < b.N; i++ {
		_, _ = NewFontFromByte(bs)
	}
}
