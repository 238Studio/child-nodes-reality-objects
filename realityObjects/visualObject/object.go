package visualObject

import "github.com/UniversalRobotDriveTeam/child-nodes-reality-objects/realityObjects/space"

type Object struct {
	//唯一的物品ID，标记一个现实中的物体
	objectID string
	//物体图像识别标签
	typeTags []string
	//相片，物体的图像信息
	photo Photo
	//物体边缘轮廓核心位置的空间位置
	position space.D3Position
	//物体的图像识别状态描述
	stateTags map[string]string
	//元信息
	metaStr map[string]string
}

func distance(obj Object) {
	return
}

func excursion(obj Object) {
	return
}
