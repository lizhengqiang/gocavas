package drawing

import (
	"image/color"
	"image"
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
