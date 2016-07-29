package drawing

import (
	"image/color"
	"image"
	"os"
	"image/png"

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
func (this *Graphics)SavePng(name string){
	f2,err:=os.Create(name)
	if(err!=nil){
		panic(err)
	}
	png.Encode(f2,&this.RGBA)
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
func (this *Graphics)Drawbrokenline(pen Pen,points...([2]int)) *Graphics{

	return this
}
func (this *Graphics)DrawPicture(img image.Image ,location,rect ([2]int)) *Graphics{
	grect :=[2]int{this.Rect.Dx(),this.Rect.Dx()}

	if(!pointInRect(grect,location)){
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
			a:=this.RGBAAt(x,y)
			colorr,colorg,colorb,colora:=img.At(x,y).RGBA()
			b:=color.RGBA{uint8(colorr),uint8(colorg),uint8(colorb),uint8(colora)}

			ab:=float64(a.A)+float64(b.A)
			cr:=uint8((float64(a.R)*float64(a.A)+float64(b.R)*float64(b.A))/ab)
			cg:=uint8((float64(a.G)*float64(a.A)+float64(b.G)*float64(b.A))/ab)
			cb:=uint8((float64(a.B)*float64(a.A)+float64(b.B)*float64(b.A))/ab)
			var c =color.RGBA{
				cr,cg,cb,
				255,
			}
			if(a.A>b.A){
				c.A=a.A
			}else {
				c.A=b.A
			}
			g2.Set(x+location[0],y+location[1],c)

		}
	}
	return g2
}
