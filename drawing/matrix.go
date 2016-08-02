package drawing

import (
	"image/color"
	"image"
	"golang.org/x/image/font"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/math/fixed"
	"io/ioutil"

	"image/png"
	"image/draw"

	"bufio"
	"bytes"

	"image/jpeg"
	"image/gif"
)

type Graphics struct  {
	image.RGBA
}
type Pen  struct{
	color	color.Color
	width int
}

func NewImageFromBytes(bs []byte)(image.Image){
	r:=bytes.NewReader(bs)

	img,err:=jpeg.Decode(r)
	if(err==nil){
		return img
	}

	r=bytes.NewReader(bs)
	img,err=png.Decode(r)
	if(err==nil){
		return img
	}

	r=bytes.NewReader(bs)
	img,err=gif.Decode(r)
	if(err==nil){
		return img
	}
	return nil
}

func NewFromImage(img image.Image)*Graphics{

	rect:=img.Bounds()

	this :=&Graphics{ *image.NewRGBA(rect)}

	for y := 0; y <rect.Dy(); y++ {
		for x := 0; x < rect.Dx(); x++ {
			this.Set(x, y, img.At(x, y))
		}
	}
	return this
}

func (this *Graphics)ResizeNewGraphics(w,h int) *Graphics{
	rgbimg:=image.NewRGBA(image.Rectangle{image.Point{0,0},image.Point{w,h}})
	for y:=0;y<h;y++ {
		if(y>rgbimg.Rect.Dy()){
			break
		}
		for x:=0;x<w;x++{
			if(x>rgbimg.Rect.Dx()){
				break
			}
			rgbimg.Set(x,y,this.At(x,y))
		}
	}
	return &Graphics{*rgbimg}
}
func (this *Graphics)SavePng()(bs []byte){
	b:=bytes.NewBuffer(make([]byte,0))
	w:=bufio.NewWriterSize(b,this.Stride)
	err:=png.Encode(w,&this.RGBA)
	if(err!=nil){
		panic(err)
	}
	return b.Bytes()
}
func (this *Graphics)SaveJpg(Quality int)(bs []byte){
	b:=bytes.NewBuffer(make([]byte,0))
	w:=bufio.NewWriterSize(b,this.Stride)
	err:=jpeg.Encode(w,&this.RGBA,&jpeg.Options{Quality})
	if(err!=nil){
		panic(err)
	}
	return b.Bytes()
}
func (this *Graphics)Drawbrokenline(pen Pen,points...([2]int)) *Graphics{

	return this
}
func (this *Graphics)DrawPicture(img image.Image ,location ([2]int)) *Graphics{
	w,h:=img.Bounds().Dx(),img.Bounds().Dy()
	for y:=0;y<h;y++ {
		for x:=0;x<w;x++{
			colorr,colorg,colorb,colora:=img.At(x,y).RGBA()
			b:=color.RGBA{uint8(colorr),uint8(colorg),uint8(colorb),uint8(colora)}
			if(b.A<128){
				continue
			}else {
				this.Set(x+location[0],y+location[1],img.At(x,y))
			}
		}
	}
	return this
}

//
func DrawString(w,h int,text string,FontName string,hinting font.Hinting,size float64,color color.Color,dpi float64,locationX,locationY int)(image.Image)  {
	fontBytes, err := ioutil.ReadFile(FontName)
	if err != nil {
		panic(err)

	}
	fo, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
	rgba := image.NewRGBA(image.Rect(0, 0,w,h))

	fg, bg := image.NewUniform(color), image.Transparent
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	d := &font.Drawer{
		Dst: rgba,
		Src: fg,
		Face: truetype.NewFace(fo, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: hinting,
		}),
	}
	d.Dot=fixed.Point26_6{fixed.I(locationX),fixed.I(locationY+int(size))}
	d.DrawString(text)
	return rgba
}
