package object

type objectSet struct {
	objects            map[string]Object //根据物体唯一ID对object进行索引
	coreObject         Object            //用于建系的参照物体，为像素最大的物体
	isAbsolutePosition bool              //是否采用绝对位置关系
}

// 初始化object set
// 传入：物体集合，是否采用绝对位置关系
// 返回：object set
func newObjectSet(objects map[string]Object, isAbsolutePosition bool) objectSet {
	return objectSet{
		objects:            objects,
		coreObject:         selectCoreObject(objects), //选取像素最大的物体作为核心物体
		isAbsolutePosition: isAbsolutePosition,
	}
}

// 在objects中选取最大的object。即选择object的photo中像素最大的那个。
// 传入：物体集合
// 返回：像素最大的物体
func selectCoreObject(objects map[string]Object) Object {
	//初始化最大object
	var maxObject Object
	maxObject.photo.pixelValues = 0

	//遍历objects，选取像素最大的object
	for _, object := range objects {
		if object.photo.pixelValues > maxObject.photo.pixelValues {
			maxObject = object
		}

		//TODO：阈值中断
		//if maxObject.photo.pixelValues > 114514 {
		//	break
		//}
	}
	return maxObject
}
