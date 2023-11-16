package space

import "math"

type D3Vector struct {
	originPosition D3Position //起始位置
	endPosition    D3Position //终止位置
	length         float64    //空间距离
}

// 初始化D3Vector
// 传入：起始位置，终止位置
// 返回：D3Vector
func newD3Vector(originPosition D3Position, endPosition D3Position) D3Vector {
	return D3Vector{
		originPosition: originPosition,
		endPosition:    endPosition,
		length:         d3VectorDistance(originPosition, endPosition),
	}
}

// 计算两个点之间的距离
// 传入：两个点的位置
// 返回：两个点之间的距离
func d3VectorDistance(originPosition D3Position, endPosition D3Position) float64 {
	//获取两个点的坐标
	x1, y1, z1 := originPosition.x, originPosition.y, originPosition.z
	x2, y2, z2 := endPosition.x, endPosition.y, endPosition.z

	//计算两个点之间的距离
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2) + math.Pow(z1-z2, 2))
}

// 矢量相加，返回一个新的位置
// 传入：一个位置
// 返回：一个新的位置
func (d3Vector D3Vector) addVector(position D3Position) D3Position {
	return D3Position{
		x: d3Vector.endPosition.x + position.x,
		y: d3Vector.endPosition.y + position.y,
		z: d3Vector.endPosition.z + position.z,
	}
}
