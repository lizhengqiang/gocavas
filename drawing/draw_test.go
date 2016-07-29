package drawing

import (
	"testing"
	"os"
	"fmt"
)

func TestGraphics_ResizeGraphics(t *testing.T) {
	file,err:=os.Open("/MacData/GOPATH/src/code.aliyun.com/timeloveboy/gocavas/upload/ocr_57fdc0a6ed5a33ccf9311e629ac94b8e8ae6cb2e.gif.jpg")
	if(err!=nil){
		panic(err)
	}
	fmt.Println(file.Stat())

}
