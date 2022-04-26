package port

import (
	"github.com/Chao2020x/funny/lib/mgl32"

	"github.com/go-gl/gl/v4.1-core/gl"
)

///
// 摄像机数据对象
// ? 日志
//! 2019年11月1日 重写
//! 2019年11月18日
//		完成着色器底层绑定
//! 2019年11月19日
//		修改底层绑定
///

// CameraObjectData 摄像机底层对象数据
type CameraObjectData struct {
	// 着色器变量指针
	GlobalMatrix uint32 // 着色器矩阵全局变量地址
	// 着色器变量
	ProjectionMatrix mgl32.Mat4 // 投影矩阵
	ViewMatrix       mgl32.Mat4 // 视口矩阵(动态)
}

// Init 初始化
func (C *CameraObjectData) Init() error {
	gl.GenBuffers(1, &C.GlobalMatrix)                                // 创建缓冲器
	gl.BindBuffer(gl.UNIFORM_BUFFER, C.GlobalMatrix)                 // 绑定缓冲区
	gl.BufferData(gl.UNIFORM_BUFFER, 128, nil, gl.DYNAMIC_DRAW)      // 分配内存空间
	gl.BindBufferRange(gl.UNIFORM_BUFFER, 0, C.GlobalMatrix, 0, 128) // 绑定着色器变量定义的绑定点,矩阵定义绑定在0号绑定点
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)                              // 解除引用
	return nil
}

// Delete 销毁
func (C *CameraObjectData) Delete() {
	gl.DeleteBuffers(1, &C.GlobalMatrix)
}

// UpdateProjectionMatrix 更新(投影矩阵)
func (C *CameraObjectData) UpdateProjectionMatrix() {
	gl.BindBuffer(gl.UNIFORM_BUFFER, C.GlobalMatrix)
	gl.BufferSubData(gl.UNIFORM_BUFFER, 0, 64, gl.Ptr(&C.ProjectionMatrix[0]))
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
}

// UpdateViewMatrix 动态更新(视图矩阵)
func (C *CameraObjectData) UpdateViewMatrix() {
	// 更新着色器值
	gl.BindBuffer(gl.UNIFORM_BUFFER, C.GlobalMatrix)
	gl.BufferSubData(gl.UNIFORM_BUFFER, 64, 64, gl.Ptr(&C.ViewMatrix[0]))
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
}

// CameraViewMatrixData 摄像机视口矩阵结构体
type CameraViewMatrixData struct {
	Up     mgl32.Vec3 // 朝向
	Eye    mgl32.Vec3 // 位置
	Center mgl32.Vec3 // 视点
}

// CameraProjectionMatrixData 摄像机投影矩阵结构体
type CameraProjectionMatrixData struct {
	Fovy   float32 // 视角
	Aspect float32 // 屏幕宽高比
	ZNear  float32 // 近平面距离
	ZFar   float32 // 远平面距离
}

// CameraData 摄像机数据默认结构体
type CameraData struct {
	object CameraObjectData // 绑定着色器底层
	// 默认摄像机数据
	CameraViewMatrixData
	CameraProjectionMatrixData
}

// Init 初始化
func (C *CameraData) Init() error {
	C.object = CameraObjectData{}
	if err := C.object.Init(); err != nil {
		return err
	}
	C.UpdateProjectionMatrix()
	return nil
}

// UpdateProjectionMatrix 投影矩阵更新
func (C *CameraData) UpdateProjectionMatrix() {
	// 投影矩阵
	C.object.ProjectionMatrix = mgl32.Perspective(mgl32.DegToRad(C.Fovy), C.Aspect, C.ZNear, C.ZFar)
	C.object.UpdateProjectionMatrix()
}

// Update 主更新(视图矩阵)
func (C *CameraData) Update() {
	// 视图矩阵
	C.object.ViewMatrix = mgl32.LookAtV(C.Eye, C.Center, C.Up)
	C.object.UpdateViewMatrix()
}

// Delete 销毁
func (C *CameraData) Delete() {
	C.object.Delete()
}
