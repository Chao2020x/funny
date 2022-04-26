package port

import (
	"errors"
	"unsafe"

	"github.com/Chao2020x/funny/lib/mgl32"

	"github.com/go-gl/gl/v4.1-core/gl"
)

///
// 灯光底层绑定
//? 日志
//!  2019年11月19日
//     绑定底层
///

// LightObjectData 静态灯光
type LightObjectData struct {
	Position mgl32.Vec3 //* 位置
	//? 颜色
	Ambient  mgl32.Vec3 //* 环境色
	Diffuse  mgl32.Vec3 //* 漫反射
	Specular mgl32.Vec3 //* 高光反射
	//? 物理参数
	Linear    float32 //* 线性
	Constant  float32 //* 常量
	Quadratic float32 //* 二次方程式
	//? 是否生效
	IsLight bool
}

// LightData 灯光
type LightData struct {
	global uint32
	//? 灯光数据
	Light    []*LightObjectData
	Quantity int
}

// Init 初始化
func (L *LightData) Init() error {
	if L.Quantity < 1 {
		return errors.New("未设置灯光初始化数量")
	}
	size := int(L.Quantity * 64)
	// 创建缓冲区
	gl.GenBuffers(1, &L.global)
	gl.BindBuffer(gl.UNIFORM_BUFFER, L.global)
	// 分配数据
	gl.BufferData(gl.UNIFORM_BUFFER, size, nil, gl.DYNAMIC_DRAW) // 分配内存空间
	gl.BindBufferRange(gl.UNIFORM_BUFFER, 1, L.global, 0, size)  // 绑定着色器变量定义的绑定点,灯光定义绑定在1号绑定点
	// 映射绑定
	pointer := (uintptr)(gl.MapBufferRange(gl.UNIFORM_BUFFER, 0, size, gl.MAP_WRITE_BIT|gl.MAP_INVALIDATE_BUFFER_BIT))
	for i := 0; i < L.Quantity; i++ {
		L.Light = append(L.Light, (*LightObjectData)(unsafe.Pointer(uintptr(i*64)+pointer)))
	}
	// 解除绑定
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
	return nil
}

// Delete 销毁
func (L *LightData) Delete() {
	gl.DeleteBuffers(1, &L.global)
}

// Update 主更新
func (L *LightData) Update() {

}
