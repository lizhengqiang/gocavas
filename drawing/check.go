package drawing

func PointInRect(rect [2]int,p [2]int)  {
	if(p[0]>0&&p[0]<rect[0]&&p[1]>0&&p[1]<rect[1]){
		return true
	}else {
		return false
	}
}
