package main

import (
	"os"
	"image/jpeg"
	"code.aliyun.com/timeloveboy/gocavas/drawing"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func main() {
	file,err:=os.Open("img/hh.jpeg")
	if(err!=nil){
		panic(err)
	}
	img,err:=jpeg.Decode(file)
	if(err!=nil){
		panic(err)
	}
	g:=&drawing.Graphics{img}
	g.DrawString("李鹏","static/张海山锐线体简1.0.ttf",font.HintingFull,12,128,fixed.Point26_6{10,10})
	g.SavePng("img/hh.png")
}
