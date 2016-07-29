package main

import (
	"os"
	"code.aliyun.com/timeloveboy/gocavas/drawing"
	"image/png"

)

func main() {
	file,err:=os.Open("/home/timeloveboy/图片/2016-07-08 21-00-28屏幕截图.png")
	if(err!=nil){
		panic(err)
	}
	defer file.Close()
	img,err:=png.Decode(file)
	g:=drawing.NewFromImage(img)

	//g=g.ResizeNewGraphics(100,100)
	//g.SavePng("i.png")


	f2,err:=os.Open("1.png")
	if(err!=nil){
		panic(err)
	}
	defer f2.Close()
	img2,err:=png.Decode(f2)

	g=g.DrawPicture(img2,[2]int{50,50},*new([2]int))
	g.SavePng("2.png")
}
