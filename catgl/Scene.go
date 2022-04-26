package catgl

import (
	"github.com/Chao2020x/funny/catgl/port"
	"github.com/go-gl/gl/v4.1-core/gl"
)

///
// 场景包
// ? 日志
//! 2019年11月1日 重写
///

// Scene 场景结构体
type Scene struct {
	port.Scene
	// 初始化用
	syncSize   int                // 通道数量
	synclock   chan bool          // 锁
	syncObject chan port.CallInit // 更新通道
}

// Init 初始化
func (S *Scene) Init() error {
	//? 初始化灯光与摄像机
	if err := S.Scene.Init(); err != nil {
		return err
	}
	//? 设置场景参数
	gl.Enable(gl.DEPTH_TEST)                           //* 深度测试
	gl.DepthFunc(gl.LESS)                              //* 小于模式
	gl.Enable(gl.BLEND)                                //* 颜色混合
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA) //* 指定混合因子
	gl.BlendEquation(gl.FUNC_ADD)                      //* 指定混合程式
	gl.Enable(gl.STENCIL_TEST)                         //* 模板测试
	gl.StencilOp(gl.KEEP, gl.KEEP, gl.REPLACE)         //* 参数测试
	gl.Enable(gl.MULTISAMPLE)                          //* 开启多采样
	//? 场景对象初始化通道
	S.synclock = make(chan bool, 1)         // 锁
	S.syncObject = make(chan port.CallInit) // 更新通道
	return nil
}

// Update 主更新
func (S *Scene) Update() {
	// 初始化场景对象
	if S.syncSize > 0 {
		S.synclock <- true
		for i := 0; i < S.syncSize; i++ {
			in := (<-S.syncObject)
			// 初始化并回调
			in.Call <- in.Init.Init()
		}
		S.syncSize = 0
		<-S.synclock
	}
	// 场景默认对象更新
	S.Camera.Update() // 更新摄像机
	S.Light.Update()  // 更新灯光
	// 游戏对象预更新
	for _, M := range S.GameObjectList {
		M.UpdateBeforehand()
	}
	// 游戏对象渲染更新
	for _, M := range S.GameObjectList {
		M.UpdateRendering()
	}
}

// AddObject 添加更新
func (S *Scene) AddObject(init port.Init) error {
	// 更新数量
	S.synclock <- true
	S.syncSize++
	<-S.synclock
	// 发送到更新
	in := port.CallInit{
		Init: init,
		Call: make(chan error),
	}
	S.syncObject <- in
	return <-in.Call
}
