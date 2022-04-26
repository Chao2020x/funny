package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// GameObject 测试对象
type GameObject struct {
	ShaderProgram uint32 // 绑定的着色器对象指针,是场景着色器的指向.
	// 顶点数据
	Vertices []float32
	// GPU 缓冲
	VBO  uint32 // 顶点数据缓冲指针
	VAO  uint32 // 顶点结构指针
	Size int32  // 顶点数量
}

// Init 初始化
func (V *GameObject) Init() error {
	V.Size = int32(len(V.Vertices) / 3) // 计算顶点数量
	// 创建(顶点数组对象)并绑定
	gl.GenVertexArrays(1, &(V.VAO))
	gl.BindVertexArray(V.VAO)
	// 创建(顶点缓冲对象)并绑定
	gl.GenBuffers(1, &(V.VBO))                                                            //? 创建 VBO
	gl.BindBuffer(gl.ARRAY_BUFFER, V.VBO)                                                 //? 激活缓冲区
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(V.Vertices), gl.Ptr(V.Vertices), gl.STATIC_DRAW) //? 添加缓冲区数据
	// 设置顶点结构
	{
		// 参数
		var posltion int32 = 24 // 一个顶点结构占用大小
		// 绑定顶点坐标
		//            绑定索引 元素数量 元素类型 不标准化 结构大小 开始偏移
		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, posltion, gl.PtrOffset(0)) //- 设置顶点结构
		gl.EnableVertexAttribArray(0)                                            //- 设置顶点layout索引
		// 绑定顶点颜色
		gl.VertexAttribPointer(2, 3, gl.FLOAT, false, posltion, gl.PtrOffset(12)) //- 设置顶点结构
		gl.EnableVertexAttribArray(2)                                             //- 设置顶点layout索引
	}
	gl.BindBuffer(gl.ARRAY_BUFFER, 0) //? 取消激活缓冲区
	// 取消激活
	gl.BindVertexArray(0)
	return nil
}

// UpdateBeforehand 预更新
func (V *GameObject) UpdateBeforehand() {

}

// UpdateRendering 着色器更新
func (V *GameObject) UpdateRendering() {
	gl.UseProgram(V.ShaderProgram)
	// 绘画顶点
	gl.BindVertexArray(V.VAO)
	gl.DrawArrays(gl.TRIANGLES, 0, V.Size)
	gl.BindVertexArray(0)
}

// Delete 更新接口
func (V *GameObject) Delete() {

}
