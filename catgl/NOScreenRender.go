package catgl

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// NOScreenRender 默认渲染
func NOScreenRender() {
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
		//? 初始化 glfw
		if err := glfw.Init(); err != nil {
			panic(err)
		}
		defer glfw.Terminate()
		// 设置glfw参数
		glfw.WindowHint(glfw.Samples, 8)                            // MSAA
		glfw.WindowHint(glfw.Resizable, glfw.False)                 // 固定窗口大小
		glfw.WindowHint(glfw.ContextVersionMajor, 4)                // 主版本号
		glfw.WindowHint(glfw.ContextVersionMinor, 6)                // 次版本号
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // 核心模式
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)    // OpenGL 兼容
		// 主循环
		lock := make(chan bool)
		for {
			select {
			case R := <-NewWindow:
				// 初始化
				go newNOScreenRender(R, lock)
				<-lock
			case <-WindowManager.Fps.C:
				if WindowManager.Exit {
					return
				}
				// 处理用户事件
				glfw.PostEmptyEvent()
			}
		}
	}
	// 初始化
	go WindowManager.new()
}

// newNOScreenRender 默认创建
func newNOScreenRender(R *Renderer, lock chan bool) {
	// 解绑线程
	runtime.UnlockOSThread() // 隔离主线程
	// 重写绑定线程
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	// 创建窗口
	var err error
	R.window, err = glfw.CreateWindow(R.Width, R.Height, R.Title, nil, nil)
	if err != nil {
		R.err <- err
		lock <- false
		return
	}
	//? 设置当前窗口上下文
	R.window.MakeContextCurrent()
	//? 初始化Gl库
	if err := gl.Init(); err != nil {
		R.err <- err
		lock <- false
		return
	}
	//? 初始化场景
	err = R.Scene.Init()
	if err != nil {
		R.err <- err
		lock <- false
		return
	}
	//? 完成创建
	R.err <- nil
	lock <- false
	// 主窗口循环
	for !R.window.ShouldClose() && !WindowManager.Exit {
		// 循环事件
		glfw.PollEvents()
		// 用户绘画处理
		R.Update()
		// 交换缓冲区
		R.window.SwapBuffers()
	}
	// 关闭并退出
	glfw.DetachCurrentContext() //? 分离上下文
	R.window.Destroy()
	R.Scene.Delete()
	R.err <- nil
}
