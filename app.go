package main

import (
	"code.aliyun.com/timeloveboy/gocavas/drawing"
	"golang.org/x/image/font"

	"os"
	"image/jpeg"

	"image/color"
)

func main() {

	textimg:=drawing.DrawString(300,200,"李鹏","static/张海山锐线体简1.0.ttf",font.HintingNone,50,color.RGBA{0,0,255,255},100,0,0)

	f,err:=os.Open("img/hh.jpeg")
	defer f.Close()
	scrimg,err:=jpeg.Decode(f)

	if(err!=nil){
		panic(err)
	}
	g:=drawing.NewFromImage(scrimg)
	g.DrawPicture(textimg,[2]int{0,0})
	g.SavePng("img/hh.png")
}

