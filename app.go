package main

import (
	"code.aliyun.com/timeloveboy/gocavas/drawing"
	"golang.org/x/image/font"

	"os"
	"image/jpeg"
	"image/png"
)

func main() {

	textimg:=drawing.DrawString(300,200,"李鹏","static/张海山锐线体简1.0.ttf",font.HintingNone,100,100,0,0)

	f,err:=os.Open("img/hh.jpeg")
	scrimg,err:=jpeg.Decode(f)

	f2,err:=os.Create("img/h.png")
	png.Encode(f2,textimg)
	if(err!=nil){
		panic(err)
	}
	g:=drawing.NewFromImage(scrimg)
	g2:=g.DrawPicture(textimg,[2]int{0,0})
	g2.SavePng("img/hh.png")
}

