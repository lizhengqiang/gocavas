package drawing

import (
	"image/color"
	"image"
	"golang.org/x/image/font"
	"github.com/golang/freetype/truetype"

	"golang.org/x/image/math/fixed"
	"io/ioutil"

	"image/png"
	"os"
)

type Graphics struct  {
	image.RGBA
}
type Pen  struct{
	color	color.Color
	width int
}

func (this *Graphics)ResizeGraphics(newrect [2]int) *Graphics{
	rgbimg:=image.NewRGBA(image.Rectangle{image.Point{0,0},image.Point{newrect[0],newrect[1]}})
	for y:=0;y<newrect[1];y++ {
		if(y>rgbimg.Rect.Dy()){
			break
		}
		for x:=0;x<newrect[0];x++{
			if(x>rgbimg.Rect.Dx()){
				break
			}
			rgbimg.Set(x,y,this.At(x,y))
		}
	}
	return &Graphics{*rgbimg}
}
func (this *Graphics)SavePng(name string){
	f,err:=os.Create(name)
	if(err!=nil){
		panic(err)
	}
	err=png.Encode(f,&this.RGBA)
	if(err!=nil){
		panic(err)
	}
}
func (this *Graphics)Drawbrokenline(pen Pen,points...([2]int)) *Graphics{

	return this
}
func (this *Graphics)DrawPicture(img image.Image ,location,rect ([2]int)) *Graphics{
	grect :=[2]int{this.Rect.Dx(),this.Rect.Dx()}

	if(PointInRect(grect,location)){
		panic("your draw location out of range")
	}

	return this
}
func (this *Graphics)DrawString(text string,FontName string,hinting font.Hinting,size float64,dpi float64,dot fixed.Point26_6)(*Graphics)  {
	fontBytes, err := ioutil.ReadFile(FontName)
	if err != nil {
		panic(err)

	}
	fo, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)

	}
	fg:=image.White
	d := &font.Drawer{
		Dst: &this.RGBA,
		Src: fg,
		Face: truetype.NewFace(fo, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: hinting,
		}),
	}
	d.DrawString(text)
	return this
}