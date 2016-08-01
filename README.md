# gocavas
>go画板

## func
+ 绘制文字的图片
```
textimg:=drawing.DrawString(300,200,"李鹏","static/张海山锐线体简1.0.ttf",font.HintingNone,50,color.RGBA{0,0,255,255},100,0,0)
```
+ 以一张图片作为背景生成画板
```
g:=drawing.NewFromImage(scrimg)
```
+ 把图片绘制在一张画板上
```
g.DrawPicture(textimg,[2]int{0,0})
```
+ 保存
```
g.SavePng("img/hh.png")
```