package port

import "errors"

// Scene 场景对象
type Scene struct {
	Light  LightObject  //? 底层灯光对象
	Camera CameraObject //? 底层摄像机对象
	// 游戏数据
	GameObjectList   map[ID]GameObject   //? 游戏对象列表
	ShaderObjectList map[ID]ShaderObject //? 游戏着色器列表
}

// Init 初始化
func (S *Scene) Init() error {
	if S.Light == nil {
		return errors.New("未定义灯光对象")
	}
	if S.Camera == nil {
		return errors.New("未定义摄像机对象")
	}
	// 初始化
	if err := S.Light.Init(); err != nil {
		return err
	}
	if err := S.Camera.Init(); err != nil {
		return err
	}
	S.GameObjectList = make(map[ID]GameObject)
	S.ShaderObjectList = make(map[ID]ShaderObject)
	return nil
}

// Delete 销毁
func (S *Scene) Delete() {
	S.Light.Delete()
	S.Camera.Delete()
	for _, S := range S.ShaderObjectList {
		S.Delete()
	}
	for _, G := range S.GameObjectList {
		G.Delete()
	}
}
