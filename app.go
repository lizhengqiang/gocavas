package main

import (
	"github.com/mengxiaozhu/gocavas/drawing"
	"golang.org/x/image/font"

	"os"

	"fmt"
	"image/color"
	"io/ioutil"
)

func main() {

	img_line1 := drawing.DrawString(300, 200, "我是李鹏", "static/张海山锐线体简1.0.ttf", font.HintingNone, 30, color.RGBA{0, 0, 255, 255}, 100, 0, 0)
	img_line2 := drawing.DrawString(300, 200, "我是韩寒", "static/张海山锐线体简1.0.ttf", font.HintingNone, 25, color.RGBA{0, 255, 0, 255}, 100, 0, 0)

	f, err := os.Open("img/hh.jpeg")

	defer f.Close()
	bs, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}
	scrimg := drawing.NewImageFromBytes(bs)

	if err != nil {
		panic(err)
	}
	g := drawing.NewFromImage(scrimg)
	g.DrawPicture(img_line1, [2]int{0, 0})
	g.DrawPicture(img_line2, [2]int{0, 40})
	bw := g.SaveJpg(100)
	ioutil.WriteFile("img/result.jpg", bw, os.ModePerm)
	fmt.Println(len(bw))

}
