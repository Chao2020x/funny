package catgl

import (
	"errors"

	"github.com/Chao2020x/funny/catgl/port"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

///
// 渲染器
//  创建渲染窗口与管理窗口事件,绑定场景等操作
// ? 日志
//! 2019年11月2日 重写
///

// Renderer 渲染器
type Renderer struct {
	window *glfw.Window
	err    chan error
	//? 场景对象
	Scene port.SceneObject
	//? 窗口参数
	Width  int
	Height int
	Title  string
	Color  [4]float32
	//? 结构体参数
	AspectRatio float32 //* 屏幕高宽比
}

// GetWindow 得到窗口
func (R *Renderer) GetWindow() *glfw.Window {
	return R.window
}

// New 创建窗口
func (R *Renderer) New() chan error {
	// 检测错误
	{
		if WindowManager == nil {
			R.err <- errors.New("初始化失败: 窗口管理器未初始化")
			return R.err
		}
		if R.Width < 0 || R.Height < 0 {
			R.err <- errors.New("初始化参数失败: 窗口大小为空")
			return R.err
		}
		if R.Scene == nil {
			R.err <- errors.New("初始化参数失败: 未绑定场景")
			return R.err
		}
	}
	// 调用创建目标
	R.err = make(chan error, 1)
	R.AspectRatio = float32(R.Width / R.Height)
	WindowManager.add(R)
	return R.err
}

// Update 更新函数
//? 在返回前住线程会被阻断
func (R *Renderer) Update() {
	gl.ClearColor(R.Color[0], R.Color[1], R.Color[2], R.Color[3])
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
	// 更新场景
	R.Scene.Update()
}
