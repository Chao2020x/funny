package catgl

import (
	"runtime"
	"time"
)

// OffScreenRender 离屏渲染
func OffScreenRender() {
	// 判断是否初始化
	if WindowManager != nil {
		return
	}
	// 初始化变量
	NewWindow := make(chan *Renderer)
	WindowManager = &manager{
		//? 默认窗口刷新频率
		Fps: time.NewTicker(time.Millisecond * time.Duration(6)),
	}
	// 添加窗口
	WindowManager.add = func(R *Renderer) {
		// 将初始化参数发送到创建函数
		NewWindow <- R
	}
	// 初始化
	WindowManager.new = func() {
		// 绑定到主进程
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		// 无窗口初始化
	}
	// 初始化
	go WindowManager.new()
}
