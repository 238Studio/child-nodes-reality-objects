package object

type Photo struct {
	//二维数组指针
	photo *[][]byte
	//相片宽度
	width int
	//相片高度
	height int
	//像素值
	pixelValues int
}
