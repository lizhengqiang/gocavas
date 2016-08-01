package drawing

import (
	"image/color"
	"image"
	"golang.org/x/image/font"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/math/fixed"
	"io/ioutil"
	"os"
	"image/png"
	"image/draw"
	"fmt"
)

type Graphics struct  {
	image.RGBA
}
type Pen  struct{
	color	color.Color
	width int
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
func (this *Graphics)DrawPicture(img image.Image ,location ([2]int)) *Graphics{
	grect :=[2]int{this.Rect.Dx(),this.Rect.Dx()}

	if(pointInRect(grect,location)){
		fmt.Print(grect,location)
		panic("your draw location out of range")
	}
	neww:=0
	newh:=0
	if(img.Bounds().Dx()>this.Rect.Dx()){
		neww=img.Bounds().Dx()
	}else {
		neww=this.Rect.Dx()
	}
	if(img.Bounds().Dy()>this.Rect.Dy()){
		newh=img.Bounds().Dy()
	}else {
		newh=this.Rect.Dy()
	}

	g2:=this.ResizeNewGraphics(neww,newh)

	w,h:=img.Bounds().Dx(),img.Bounds().Dy()

	for y:=0;y<h;y++ {
		for x:=0;x<w;x++{
			colorr,colorg,colorb,colora:=img.At(x,y).RGBA()
			b:=color.RGBA{uint8(colorr),uint8(colorg),uint8(colorb),uint8(colora)}
			if(b.A==0){
				continue
			}else {
				g2.Set(x+location[0],y+location[1],img.At(x,y))
			}
		}
	}
	return g2
}
func DrawString(w,h int,text string,FontName string,hinting font.Hinting,size float64,dpi float64,locationX,locationY int)(image.Image)  {
	fontBytes, err := ioutil.ReadFile(FontName)
	if err != nil {
		panic(err)

	}
	fo, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
	rgba := image.NewRGBA(image.Rect(0, 0,w,h))
	fg, bg := image.Black, image.Transparent
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