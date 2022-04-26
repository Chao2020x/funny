package catgl

import (
	"time"
)

// manager 窗口管理程序
type manager struct {
	Fps  *time.Ticker
	Exit bool
	//? 内部参数
	new func()
	add func(R *Renderer)
}

// WindowManager 窗口管理程序
var WindowManager *manager
